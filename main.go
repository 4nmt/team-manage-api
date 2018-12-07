package main

import (
	"fmt"
	"os"

	"github.com/4nmt/team-manage-api/controllers"
	"github.com/4nmt/team-manage-api/db"

	"github.com/gin-gonic/gin"
)

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	db.Init()

	v1 := r.Group("/v1")
	{
		/*** START USER ***/
		users := new(controllers.UserController)

		v1.GET("/user", users.All)
		v1.GET("/user/:id", users.One)
		v1.POST("/user", users.Create)
		v1.DELETE("/user/:id", users.Delete)

		/*** START PROJECT ***/
		projects := new(controllers.ProjectController)

		v1.GET("/project", projects.All)
		v1.GET("/project/:id", projects.One)
		v1.POST("/project", projects.Create)
		v1.DELETE("/project/:id", projects.Delete)

		/*** START USER PROJECT ***/
		userProject := new(controllers.UserProjectController)

		v1.POST("/user_project", userProject.Assign)
		v1.DELETE("/user_project/:id", userProject.Remove)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	r.Run(":" + port)
}
