package forms

//UserForm ...
type UserForm struct {
	Email string `form:"email" json:"email" binding:"required,email"`
	Name  string `form:"name" json:"name" binding:"required"`
}
