package controller

import (
	"log"
	"strings"
	CON "web-form-go/src/config"
	"web-form-go/src/config/err"
	models "web-form-go/src/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserLogin(c *gin.Context) {
	var user models.User

	identifier := strings.TrimSpace(c.PostForm("identifier"))
	password := strings.TrimSpace(c.PostForm("password"))

	resp := err.ErrorResponse{
		Error: make(map[string]string),
	}

	isEmail := strings.Contains(identifier, "@")
	var queryField string
	if isEmail {
		queryField = "email"
	} else {
		queryField = "username"
	}

	db := CON.DB()

	err := db.QueryRow("SELECT id, email, password FROM user1 WHERE "+queryField+"=?", identifier).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		log.Println("Error executing SQL statement:", err)
		resp.Error["credentials"] = "Invalid credentials"
	}

	encErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if encErr != nil {
		resp.Error["password"] = "Invalid password"
	}

	if len(resp.Error) > 0 {
		c.JSON(400, resp)
		return
	}

	c.JSON(200, gin.H{"message": "User logged in successfully"})
}
