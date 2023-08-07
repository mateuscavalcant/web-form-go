package routes

import (
	"web-form-go/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.UserLogin)
	r.POST("/validate-email", controller.ExistEmail)
	r.POST("/validate-username", controller.ExistUsername)
}
