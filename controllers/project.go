package controllers

import (
	"strconv"

	"github.com/4nmt/team-manage-api/forms"
	"github.com/4nmt/team-manage-api/models"

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

//Delete ...
func (ctrl ProjectController) Delete(c *gin.Context) {
	projectID := c.Param("id")

	var (
		id  int
		err error
	)

	if id, err = strconv.Atoi(projectID); err != nil {
		c.JSON(406, gin.H{"message": "Count not parse string to int", "error": err.Error()})
		c.Abort()
	}

	err = projectModel.Delete(id)
	if err != nil {
		c.JSON(406, gin.H{"message": "could not be removed this project", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "This has been removed", "id": projectID})
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

//All ...
func (ctrl ProjectController) All(c *gin.Context) {
	projects, err := projectModel.All()
	if err != nil {
		c.JSON(404, gin.H{"Message": "Projects not found", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": projects})
}
