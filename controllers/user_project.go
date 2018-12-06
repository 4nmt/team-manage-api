package controllers

import (
	"strconv"

	"github.com/4nmt/TeamManageApp-api/forms"
	"github.com/4nmt/TeamManageApp-api/models"

	"github.com/gin-gonic/gin"
)

//ProjectController ...
type UserProjectController struct{}

var userProjectModel = new(models.UserProjectModel)

//Create ...
func (ctrl UserProjectController) Assign(c *gin.Context) {
	var userProjectForm forms.UserProjectForm

	if c.BindJSON(&userProjectForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": userProjectForm})
		c.Abort()
		return
	}

	userProjectID, err := userProjectModel.Assign(userProjectForm)

	if userProjectID <= 0 && err != nil {
		c.JSON(406, gin.H{"message": "Project could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Project created", "id": userProjectID})
}

//Create ...
func (ctrl UserProjectController) Remove(c *gin.Context) {
	userProjectID := c.Param("id")

	var (
		id  int
		err error
	)

	if id, err = strconv.Atoi(userProjectID); err != nil {
		c.JSON(406, gin.H{"message": "Count not parse string to int", "error": err.Error()})
		c.Abort()
	}

	err = userProjectModel.Remove(id)
	if err != nil {
		c.JSON(406, gin.H{"message": "could not be removed this user_project", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "This has been removed", "id": userProjectID})
}
