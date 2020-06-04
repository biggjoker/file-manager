package app

import (
	"github.com/gin-gonic/gin"
	"github.com/biggjoker/file-manager/app/dic"
	"github.com/biggjoker/file-manager/app/user"
)

func Routes(r *gin.Engine) {
	gr := r.Group("/file-manager")
	dic.Routes(gr)
	user.Routes(gr)
}
