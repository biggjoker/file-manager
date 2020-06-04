package dic

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/biggjoker/file-manager/app/helper"
	"github.com/biggjoker/file-manager/protocol"
	"github.com/biggjoker/file-manager/g"
	"github.com/biggjoker/file-manager/zlog"
	"github.com/gin-gonic/gin"
)

func GetDic(c *gin.Context) {
	path := g.Config().BaseDir
	if pathAdd, ok := c.GetQuery("path"); ok && pathAdd != "" {
		path +=  pathAdd
	}
	fileSlice := getCurrentChildFiles(path)

	for _, v := range fileSlice {
		v.Name = strings.TrimPrefix(v.Name, g.Config().BaseDir)
	}
	helper.JSONR(c, fileSlice)
}

func getCurrentChildFiles(dirpath string) (fileSlices []*protocol.FileInfo) {
	fileSlices = make([]*protocol.FileInfo, 0, 10)
	files, _ := ioutil.ReadDir(dirpath)
	for _, f := range files {
		fileInfo := &protocol.FileInfo{
			Name:  dirpath + "/" + f.Name(),
			IsDir: f.IsDir(),
			Size: f.Size(),
		}
		fileSlices = append(fileSlices, fileInfo)
	}
	return fileSlices
}

func getFullPath(file string) string {
	file = strings.TrimPrefix(file, "./")
	return g.Config().BaseDir + file
}

func CreateDir(c *gin.Context) {
	var req protocol.CreateDirReq
	if err := c.Bind(&req); err != nil {
		helper.JSONR(c, http.StatusBadRequest, helper.INPUT_FORMATE_ERROR)
		return
	}
	path := getFullPath(req.Path)
	println(path)
	if err := os.MkdirAll(getFullPath(req.Path), 0755); err != nil {
		helper.JSONR(c, http.StatusInternalServerError, err)
		return
	}
	helper.JSONR(c, "create dir success!")
}

func DeleteDir(c *gin.Context) {
	req := struct {
		DeletePath string `json:"dir_path"`
		force      string `json:"force"`
	}{}
	if err := c.Bind(&req); err != nil {
		helper.JSONR(c, http.StatusBadRequest, helper.INPUT_FORMATE_ERROR)
		return
	}
	force := strings.ToLower(req.force)

	var err error
	if force == "TRUE" {
		err = os.RemoveAll(getFullPath(req.DeletePath))
	} else {
		err = os.Remove(getFullPath(req.DeletePath))
	}
	if err != nil {
		zlog.Errorw("delete dir faile", "dir", req.DeletePath, "err", err)
		helper.JSONR(c, http.StatusInternalServerError, err)
		return
	}
	helper.JSONR(c, "delete dir success!")
}

func RenameDir(c *gin.Context) {
	var req protocol.RenameReq
	if err := c.Bind(&req); err != nil {
		helper.JSONR(c, http.StatusBadRequest, helper.INPUT_FORMATE_ERROR)
		return
	}

	err := os.Rename(getFullPath(req.Old), getFullPath(req.New))
	if err != nil {
		zlog.Errorw("rename faile", "old", req.Old, "new", req.New, "err", err)
		helper.JSONR(c, http.StatusInternalServerError, err)
		return
	}
	helper.JSONR(c, "rename success!")
}

func ReadFile(c *gin.Context) {
	path, ok := c.GetQuery("name")
	if !ok {
		helper.JSONR(c, http.StatusBadRequest, helper.INPUT_FORMATE_ERROR)
		return
	}
	if !g.IsPermissionFile(path) {
		helper.JSONR(c, http.StatusBadRequest, helper.INPUT_FILE_NO_PERMISSION)
		return
	}
	content, err := ioutil.ReadFile(getFullPath(path))
	if err != nil {
		zlog.Errorw("readfile faile", "file", path, "err", err)
		helper.JSONR(c, http.StatusInternalServerError, err)
		return
	}
	if len(content) > 30*1024*1024 {
		helper.JSONR(c, http.StatusInternalServerError, errors.New("file is not large"))
		return
	}
	helper.JSONR(c, string(content))
}

func SaveFile(c *gin.Context) {
	var req protocol.SaveFileReq
	if err := c.Bind(&req); err != nil {
		helper.JSONR(c, http.StatusBadRequest, helper.INPUT_FORMATE_ERROR)
		return
	}
	if !g.IsPermissionFile(req.Name) {
		helper.JSONR(c, http.StatusBadRequest, helper.INPUT_FILE_NO_PERMISSION)
		return
	}
	pem := 0644

	err := ioutil.WriteFile(getFullPath(req.Name), []byte(req.Content), os.FileMode(pem))
	if err != nil {
		zlog.Errorw("save faile", "file", req.Name, "err", err)
		helper.JSONR(c, http.StatusInternalServerError, err)
		return
	}
	helper.JSONR(c, "save file successed!")
}

func DeleteFile(c *gin.Context) {
	req := struct {
		Name string `json:"name"`
	}{}
	if err := c.Bind(&req); err != nil {
		helper.JSONR(c, http.StatusBadRequest, helper.INPUT_FORMATE_ERROR)
		return
	}
	if !g.IsPermissionFile(req.Name) {
		helper.JSONR(c, http.StatusBadRequest, helper.INPUT_FILE_NO_PERMISSION)
		return
	}
	err := os.Remove(getFullPath(req.Name))
	if err != nil {
		zlog.Errorw("delete faile", "file", req.Name, "err", err)
		helper.JSONR(c, http.StatusInternalServerError, err)
		return
	}
	helper.JSONR(c, "delete success!")
}

func RenameFile(c *gin.Context) {
	var req protocol.RenameReq
	if err := c.Bind(&req); err != nil {
		helper.JSONR(c, http.StatusBadRequest, helper.INPUT_FORMATE_ERROR)
		return
	}

	if !g.IsPermissionFile(req.Old) || !g.IsPermissionFile(req.New) {
		helper.JSONR(c, http.StatusBadRequest, helper.INPUT_FILE_NO_PERMISSION)
		return
	}
	err := os.Rename(getFullPath(req.Old), getFullPath(req.New))
	if err != nil {
		zlog.Errorw("rename faile", "old", req.Old, "new", req.New, "err", err)
		helper.JSONR(c, http.StatusInternalServerError, err)
		return
	}
	helper.JSONR(c, "rename success!")
}
