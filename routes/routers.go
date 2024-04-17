package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/robert/authmiddleware/handler"
)

var secret = []byte("secret")

//var userkey = "user"

func Startserver() {

	r := gin.Default()

	r.Use(gin.Logger())

	r.Use(sessions.Sessions("mysession", cookie.NewStore(secret)))

	// login and logout routes
	r.POST("/login", handler.Login)
	r.GET("/logout", handler.Logout)

	// auth group
	private := r.Group("/private")

	private.Use(handler.AuthRequired)
	{
		private.GET("/me", handler.Me)
		private.GET("/status", handler.Status)
	}

	r.Run()
}
