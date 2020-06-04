package user

import (
	"github.com/gin-gonic/gin"
	"github.com/biggjoker/file-manager/app/helper"
	"net/http"
)

const CookieUser string = "user"
const CookiePwd string = "passwd"

func Login(c *gin.Context) {

	user, _ := c.GetQuery(CookieUser)
	passwd, _ := c.GetQuery(CookiePwd)
	if user != "" && passwd != "" {
		helper.JSONR(c, "success")
		return
	}
	helper.JSONR(c, http.StatusUnauthorized, "用户名密码随便填点")
}
