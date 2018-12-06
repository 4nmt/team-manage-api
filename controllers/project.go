package controllers

import (
	"strconv"

	"github.com/4nmt/TeamManageApp-api/forms"
	"github.com/4nmt/TeamManageApp-api/models"

	"github.com/gin-gonic/gin"
)

//ProjectController ...
type ProjectController struct{}

var projectModel = new(models.ProjectModel)

//Create ...
func (ctrl ProjectController) Create(c *gin.Context) {
	var projectForm forms.ProjectForm

	if c.BindJSON(&projectForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": projectForm})
		c.Abort()
		return
	}
	projectID, err := projectModel.Create(projectForm)

	if projectID <= 0 && err != nil {
		c.JSON(406, gin.H{"message": "Project could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Project created", "id": projectID})
}

//One ...
func (ctrl ProjectController) One(c *gin.Context) {
	id := c.Param("id")

	if id, err := strconv.Atoi(id); err == nil {
		data, err := projectModel.One(id)
		if err != nil {
			c.JSON(404, gin.H{"Message": "Project not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"data": data})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}
