package dic

import (
	//"strings"

	"io"
	"net/http"
	"os"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/biggjoker/file-manager/app/helper"
	"github.com/biggjoker/file-manager/g"
	"github.com/biggjoker/file-manager/file"
)

var (
	uploadFileKey  = "file"
	uploadFilePath = "path"
)

func saveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	//创建 dst 文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	// 拷贝文件
	_, err = io.Copy(out, src)

	return err
}

func Upload(c *gin.Context) {

	header, err := c.FormFile(uploadFileKey)
	path, flag := c.GetPostForm(uploadFilePath)
	if err != nil || !flag {
		helper.JSONR(c, http.StatusBadRequest, err)
	}
	if path != "/" {
		path += "/"
	}
	dst := g.Config().BaseDir + path + header.Filename

	if err := saveUploadedFile(header, dst); err != nil {
		helper.JSONR(c, http.StatusInternalServerError, err)
	}
	res, err := file.GetFileInfoByPath(dst)
	if err != nil {
		helper.JSONR(c, http.StatusInternalServerError, err)
	}
	helper.JSONR(c, res)
}
