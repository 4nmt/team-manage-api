package forms

//ProjectForm ...
type ProjectForm struct {
	Name        string `form:"name" json:"name" binding:"required,max=100"`
	Description string `form:"description" json:"description" binding:"required,max=1000"`
}
