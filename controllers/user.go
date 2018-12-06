package controllers

import (
	"github.com/4nmt/team-manage-api/forms"
	"github.com/4nmt/team-manage-api/models"

	"github.com/gin-gonic/gin"
)

//UserController ...
type UserController struct{}

var userModel = new(models.UserModel)

//AddUser ...
func (ctrl UserController) AddUser(c *gin.Context) {
	var UserForm forms.UserForm

	if c.BindJSON(&UserForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": UserForm})
		c.Abort()
		return
	}

	user, err := userModel.AddUser(UserForm)

	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if user.ID > 0 {
		c.JSON(200, gin.H{"message": "add user successfully ", "user": user})
	} else {
		c.JSON(406, gin.H{"message": "Could not add this user", "error": err.Error()})
	}

}

// func (ctrl UserController) One(c *gin.Context) {
// 	var UserForm forms.UserForm

// 	if c.BindJSON(&UserForm) != nil {
// 		c.JSON(406, gin.H{"message": "Invalid form", "form": UserForm})
// 		c.Abort()
// 		return
// 	}

// 	user, err := userModel.AddUser(UserForm)

// 	if err != nil {
// 		c.JSON(406, gin.H{"message": err.Error()})
// 		c.Abort()
// 		return
// 	}

// 	if user.ID > 0 {
// 		c.JSON(200, gin.H{"message": "add user successfully ", "user": user})
// 	} else {
// 		c.JSON(406, gin.H{"message": "Could not add this user", "error": err.Error()})
// 	}

// }
