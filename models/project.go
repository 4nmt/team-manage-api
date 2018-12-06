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
	Users       []forms.UserForm `json:"users"`
}

//ProjectModel ...
type ProjectModel struct{}

//Create ...
func (m ProjectModel) Create(form forms.ProjectForm) (articleID int64, err error) {
	getDb := db.GetDB()

	var projectID int64
	err = getDb.QueryRow("INSERT INTO public.projects( name,description, updated_at, created_at) VALUES($1, $2, $3, $4) RETURNING id", form.Name, form.Description, time.Now().Unix(), time.Now().Unix()).Scan(&projectID)

	if err != nil {
		return -1, err
	}

	return projectID, err
}

//One ...
func (m ProjectModel) One(projectID int) (project Project, err error) {
	err = db.GetDB().SelectOne(&project, "SELECT id, name, description FROM projects WHERE id=$1 LIMIT 1", projectID)
	if err != nil {
		return project, err
	}

	users := []forms.UserForm{}
	_, err = db.GetDB().Select(&users, "SELECT distinct u.name,u.email FROM (users u JOIN user_project up ON u.id = up.user_id) JOIN projects p ON up.project_id = p.id WHERE p.id=$1", projectID)
	if err != nil {
		return project, err
	}
	project.Users = users

	return project, err
}

// //All ...
// func (m ArticleModel) All(userID int64) (articles []Article, err error) {
// 	_, err = db.GetDB().Select(&articles, "SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM public.article a LEFT JOIN public.user u ON a.user_id = u.id WHERE a.user_id=$1 ORDER BY a.id DESC", userID)
// 	return articles, err
// }
