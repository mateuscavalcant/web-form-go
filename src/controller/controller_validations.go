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
	existEmail, err := validators.ExistEmail(email)
	if err != nil {
		log.Println("Error checking email existence:", err)
		c.JSON(500, gin.H{"error": "Failed to validate email"})
		return
	}

	if email == "" {
		resp.Error["missing"] = "Some values are missing!"
	}

	if validators.ValidateFormatEmail(email) != nil {
		resp.Error["email"] = "Invalid email format!"
	}

	if existEmail {
		resp.Error["email"] = "Email already exists!"
	}

	if len(resp.Error) > 0 {
		c.JSON(400, resp)
		return
	}

	c.JSON(200, gin.H{"message": "Email Validad"})

}

func ExistUsername(c *gin.Context) {
	resp := err.ErrorResponse{
		Error: make(map[string]string),
	}

	username := strings.TrimSpace(c.PostForm("username"))
	existUsername, errUsername := validators.ExistUsername(username)
	if errUsername != nil {
		log.Println("Error checking email existence:", errUsername)
		c.JSON(500, gin.H{"error": "Failed to validate email"})
		return
	}

	if username == "" {
		resp.Error["missing"] = "Some values are missing!"
	}

	if existUsername {
		resp.Error["username"] = "Username already exists!"
	}

	if len(resp.Error) > 0 {
		c.JSON(400, resp)
		return
	}

	c.JSON(200, gin.H{"message": "Username Validad"})

}
