package api

import (
	"log"
	"net/http"
	"web-form-go/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitAPI() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.LoadHTMLGlob("C:/web-form-go/templates/*")
	r.Static("/static", "./static")

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{})
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	routes.InitRoutes(r.Group("/"))

	errServer := r.Run(":8080")
	if errServer != nil {
		return errServer
	}

	return nil
}
