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
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
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

	// store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	// r.Use(sessions.Sessions("gin-boilerplate-session", store))

	r.Use(CORSMiddleware())

	db.Init()

	v1 := r.Group("/")
	{
		/*** START USER ***/
		users := new(controllers.UserController)

		v1.POST("/user", users.AddUser)
		// v1.POST("/user/signup", user.Signup)
		// v1.GET("/user/signout", user.Signout)

		/*** START Article ***/
		project := new(controllers.ProjectController)

		v1.POST("/project", project.Create)
		// v1.GET("/articles", article.All)
		v1.GET("/project/:id", project.One)
		// v1.PUT("/article/:id", article.Update)
		// v1.DELETE("/article/:id", article.Delete)

		/*** START Article ***/
		userProject := new(controllers.UserProjectController)

		v1.POST("/user_project", userProject.Assign)
		// v1.GET("/articles", article.All)
		// v1.GET("/article/:id", article.One)
		// v1.PUT("/article/:id", article.Update)
		v1.DELETE("/user_project/:id", userProject.Remove)
	}

	// r.LoadHTMLGlob("./public/html/*")

	// r.Static("/public", "./public")

	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"ginBoilerplateVersion": "v0.03",
	// 		"goVersion":             runtime.Version(),
	// 	})
	// })

	// r.NoRoute(func(c *gin.Context) {
	// 	c.HTML(404, "404.html", gin.H{})
	// })

	port := os.Getenv("PORT") || "9000"

	r.Run(port)
}
