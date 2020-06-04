package helper

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const INPUT_FORMATE_ERROR string = "the input formate is wrong!!"
const INPUT_FILE_NO_PERMISSION string = "the input cannot be visit"

type RespJson struct {
	RetCode int         `json:"code"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// func JSONR(c *gin.Context, wcode int, msg interface{}) (werror error) {
func JSONR(c *gin.Context, arg ...interface{}) (werror error) {
	var (
		wcode int
		msg   interface{}
	)
	if len(arg) == 1 {
		wcode = http.StatusOK
		msg = arg[0]
	} else {
		wcode = arg[0].(int)
		msg = arg[1]
	}
	var body interface{}

	if wcode == 200 {
		switch msg.(type) {
		case string:
			body = RespJson{RetCode: wcode, Data: msg.(string)}
			c.JSON(http.StatusOK, body)
		default:
			body = RespJson{RetCode: wcode, Data: msg}
			c.JSON(http.StatusOK, body)
		}
	} else {
		switch msg.(type) {
		case string:
			body = RespJson{RetCode: wcode, Error: msg.(string)}
			c.JSON(http.StatusOK, body)
		case error:
			body = RespJson{RetCode: wcode, Error: msg.(error).Error()}
			c.JSON(http.StatusOK, body)
		default:
			body = RespJson{RetCode: wcode, Error: "system type error. please ask admin for help"}
			c.JSON(http.StatusOK, body)
		}
	}
	return
}

// GetQueryInt gin.Context获取int型参数
func GetQueryInt(c *gin.Context, key string) (int, error) {
	idStr, ok := c.GetQuery(key)
	if !ok {
		return 0, fmt.Errorf("gin context has no key of %s", key)
	}
	id, err := strconv.Atoi(idStr)
	if idStr == "" || err != nil {
		return 0, fmt.Errorf("gin context key of %s is ", idStr)
	}
	return id, nil
}
