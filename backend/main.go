package main

import (
	"aikido/controllers"
	"aikido/db"
	"aikido/db/relations"
	"aikido/forms"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	// db
	db.Init()
	if err := relations.Create(); err != nil {
		panic(err)
	}

	// validator
	binding.Validator = new(forms.DefaultValidator)

	// router
	router := gin.Default()

	// static
	router.Use(static.Serve("/", static.LocalFile("static", true)))
	router.NoRoute(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, "/api") {
			c.File("static/index.html")
		}
	})

	// cors
	cors := new(controllers.CorsController)
	router.Use(cors.Middleware)

	// api
	api := router.Group("/api")

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
	api.GET("/pool/all", auth.Middleware, pool.All)
	api.GET("/pool/:id", auth.Middleware, pool.Get)
	api.DELETE("/pool/:id", auth.Middleware, pool.Delete)

	// vote
	vote := new(controllers.VoteController)
	api.POST("/pool/:id/vote", auth.Middleware, vote.Create)
	api.DELETE("/pool/:id/vote", auth.Middleware, vote.Delete)
	api.GET("/pool/:id/vote/:option_id", auth.Middleware, vote.VotedUsers)

	// group
	group := new(controllers.GroupController)
	api.GET("/group", auth.Middleware, group.All)
	api.POST("/group", auth.Middleware, group.Create)
	api.PATCH("/group/:id", auth.Middleware, group.Update)
	api.DELETE("/group/:id", auth.Middleware, group.Delete)
	api.GET("/group/:id/user", auth.Middleware, group.Users)
	api.POST("/group/:id/user", auth.Middleware, group.Add)
	api.DELETE("/group/:id/user/:user_id", auth.Middleware, group.Remove)

	// chat
	chat := new(controllers.ChatController)
	api.POST("/chat/:id", auth.Middleware, chat.WriteMessage)
	api.GET("/chat/:id", auth.Middleware, chat.ReadMessages)

	panic(
		autotls.RunWithManager(router, &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("web-portfolio.tech"),
			Cache:      autocert.DirCache("/cache"),
		}),
	)
}
