package controllers

import (
	"strconv"

	"github.com/4nmt/team-manage-api/forms"
	"github.com/4nmt/team-manage-api/models"

	"github.com/gin-gonic/gin"
)

//UserController ...
type UserController struct{}

var userModel = new(models.UserModel)

//Create ...
func (ctrl UserController) Create(c *gin.Context) {
	var UserForm forms.UserForm

	if c.BindJSON(&UserForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": UserForm})
		c.Abort()
		return
	}

	user, err := userModel.Create(UserForm)

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

//Delete ...
func (ctrl UserController) Delete(c *gin.Context) {
	userID := c.Param("id")

	var (
		id  int
		err error
	)

	if id, err = strconv.Atoi(userID); err != nil {
		c.JSON(406, gin.H{"message": "Count not parse string to int", "error": err.Error()})
		c.Abort()
	}

	err = userModel.Delete(id)
	if err != nil {
		c.JSON(406, gin.H{"message": "could not be removed this user", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "This has been removed", "id": userID})
}

//One ...
func (ctrl UserController) One(c *gin.Context) {
	userID := c.Param("id")

	var (
		id   int
		err  error
		user models.User
	)
	if id, err = strconv.Atoi(userID); err != nil {
		c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
		c.Abort()
		return
	}

	user, err = userModel.One(id)
	if err != nil {
		c.JSON(404, gin.H{"Message": "User not found", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": user})
}

//All ...
func (ctrl UserController) All(c *gin.Context) {
	users, err := userModel.All()
	if err != nil {
		c.JSON(404, gin.H{"Message": "Users not found", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": users})
}
