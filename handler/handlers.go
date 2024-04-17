package handler

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var userkey = "user"

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)

	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unautherized"})
		return
	}

	c.Next()
}

func Me(c *gin.Context) {
	session := sessions.Default(c)

	user := session.Get(userkey)

	c.JSON(http.StatusOK, gin.H{"message": user})
}

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "your are logged in"})
}

func Login(c *gin.Context) {

	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no valid data"})
		return
	}

	if username != "robert" && password != "12345678" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong credentials"})
		return
	}

	session.Set(userkey, username)

	if err := session.Save(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "succesfully saved"})

}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)

	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}

	session.Delete(userkey)

	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
