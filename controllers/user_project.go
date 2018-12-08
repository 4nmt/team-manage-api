package controllers

import (
	"strconv"

	"github.com/4nmt/team-manage-api/forms"
	"github.com/4nmt/team-manage-api/models"

	"github.com/gin-gonic/gin"
)

//ProjectController ...
type UserProjectController struct{}

var userProjectModel = new(models.UserProjectModel)

//Assign ...
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

//Remove ...
func (ctrl UserProjectController) Remove(c *gin.Context) {
	userIDTmp, projectIDTmp := c.Param("user_id"), c.Param("project_id")

	var (
		userID, projectID int
		err               error
	)

	if userID, err = strconv.Atoi(userIDTmp); err != nil {
		c.JSON(406, gin.H{"message": "Count not parse userID string to int", "error": err.Error()})
		c.Abort()
	}

	if projectID, err = strconv.Atoi(projectIDTmp); err != nil {
		c.JSON(406, gin.H{"message": "Count not parse projectID string to int", "error": err.Error()})
		c.Abort()
	}

	err = userProjectModel.Remove(userID, projectID)
	if err != nil {
		c.JSON(406, gin.H{"message": "could not be removed this user_project", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "This has been removed"})
}
