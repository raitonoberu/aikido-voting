package main

import (
	"aikido/controllers"
	"aikido/db"
	"aikido/db/relations"
	"aikido/forms"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	// db
	db.Init()
	relations.Create()

	// validator
	binding.Validator = new(forms.DefaultValidator)

	// api
	r := gin.Default()
	api := r.Group("/api")
	api.Use(CORSMiddleware)

	// auth
	auth := new(controllers.AuthController)

	// user
	user := new(controllers.UserController)
	api.POST("/user", user.Register)
	api.POST("/user/login", user.Login)
	api.GET("/user/:id", auth.Middleware, user.Get)
	api.PATCH("/user", auth.Middleware, user.Update)
	api.DELETE("/user", auth.Middleware, user.Delete)

	// pool
	pool := new(controllers.PoolController)
	api.POST("/pool", auth.Middleware, pool.Create)
	api.GET("/pool", auth.Middleware, pool.Available)
	api.GET("/pool/:id", auth.Middleware, pool.Get)
	api.DELETE("/pool/:id", auth.Middleware, pool.Delete)

	panic(r.Run(":8080"))
}

func CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PATCH, DELETE")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}
