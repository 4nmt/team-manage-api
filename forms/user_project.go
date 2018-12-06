package forms

//UserProjectForm ...
type UserProjectForm struct {
	UserID    int `form:"user_id" json:"user_id"`
	ProjectID int `form:"project_id" json:"project_id"`
}
