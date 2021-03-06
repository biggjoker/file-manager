package dic

import "github.com/gin-gonic/gin"

func Routes(r gin.IRouter) {
	dicapi := r.Group("/dir")

	dicapi.GET("/dir", GetDic)
	dicapi.POST("/dir", CreateDir)
	dicapi.DELETE("/dir", DeleteDir)
	dicapi.PUT("/dir", RenameDir)

	dicapi.GET("/file", ReadFile)
	dicapi.POST("/file", CreateFile)
	dicapi.DELETE("/file", DeleteFile)
	dicapi.PUT("/file", SaveFile)
	dicapi.POST("/renamefile", RenameFile)

	dicapi.POST("/upfile", Upload)
	dicapi.GET("/upfile", Download)
}
