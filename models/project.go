package models

import (
	"time"

	"github.com/4nmt/team-manage-api/db"
	"github.com/4nmt/team-manage-api/forms"
)

//Project ...
type Project struct {
	ID          int64            `db:"id, primarykey, autoincrement" json:"id"`
	Name        string           `db:"name" json:"name"`
	Description string           `db:"description" json:"description"`
	UpdatedAt   int64            `db:"updated_at" json:"updated_at"`
	CreatedAt   int64            `db:"created_at" json:"created_at"`
	Users       []User `json:"users"`
}

//ProjectModel ...
type ProjectModel struct{}

//Create ...
func (m ProjectModel) Create(form forms.ProjectForm) (projectID int64, err error) {
	getDb := db.GetDB()

	err = getDb.QueryRow("INSERT INTO public.projects( name,description, updated_at, created_at) VALUES($1, $2, $3, $4) RETURNING id", form.Name, form.Description, time.Now().Unix(), time.Now().Unix()).Scan(&projectID)
	if err != nil {
		return -1, err
	}

	return projectID, err
}

//Delete ...
func (m ProjectModel) Delete(projectID int) error {
	_, err := m.One(projectID)
	if err != nil {
		return err
	}

	if _, err = db.GetDB().Exec("DELETE FROM public.user_project WHERE project_id=$1", projectID); err != nil {
		return err
	}

	_, err = db.GetDB().Exec("DELETE FROM public.users WHERE id=$1", projectID)

	return err
}

//One ...
func (m ProjectModel) One(projectID int) (project Project, err error) {
	err = db.GetDB().SelectOne(&project, "SELECT id, name, description FROM projects WHERE id=$1 LIMIT 1", projectID)
	if err != nil {
		return project, err
	}

	users := []User{}
	_, err = db.GetDB().Select(&users, "SELECT distinct u.id, u.name,u.email FROM (users u JOIN user_project up ON u.id = up.user_id) JOIN projects p ON up.project_id = p.id WHERE p.id=$1", projectID)
	if err != nil {
		return project, err
	}
	project.Users = users

	return project, err
}

//All ...
func (m ProjectModel) All() (projects []Project, err error) {
	_, err = db.GetDB().Select(&projects, "SELECT id, name, description, updated_at, created_at FROM public.projects")
	return projects, err
}
