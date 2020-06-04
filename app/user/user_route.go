package user

import "github.com/gin-gonic/gin"

func Routes(r gin.IRouter) {
	saltstackapi := r.Group("/user")

	saltstackapi.GET("/login", Login)
}
