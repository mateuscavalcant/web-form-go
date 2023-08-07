package controller

import (
	"log"
	"strings"
	"web-form-go/src/config/err"
	"web-form-go/src/config/validators"

	"github.com/gin-gonic/gin"
)

func ExistEmail(c *gin.Context) {
	resp := err.ErrorResponse{
		Error: make(map[string]string),
	}

	email := strings.TrimSpace(c.PostForm("email"))
	username := strings.TrimSpace(c.PostForm("username"))
	existEmail, errEmail := validators.ExistEmail(email)
	existUsername, errUsername := validators.ExistUsername(username)
	if errEmail != nil {
		log.Println("Error checking email existence:", errEmail)
		c.JSON(500, gin.H{"error": "Failed to validate email"})
		return
	}

	if email == "" {
		resp.Error["email"] = "Some values are missing!"
	}
	if validators.ValidateFormatEmail(email) != nil {
		resp.Error["email"] = "Invalid email format!"
	}
	if existEmail {
		resp.Error["email"] = "Email already exists!"
	}
	if errUsername != nil {
		log.Println("Error checking email existence:", errUsername)
		c.JSON(500, gin.H{"error": "Failed to validate email"})
		return
	}
	if username == "" {
		resp.Error["username"] = "Some values are missing!"
	}
	if existUsername {
		resp.Error["username"] = "Username already exists!"
	}
	if len(resp.Error) > 0 {
		c.JSON(400, resp)
		return
	}

	c.JSON(200, gin.H{"message": "Email Validad"})

}
