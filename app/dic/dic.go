package dic

import (
	"errors"
	"github.com/biggjoker/file-manager/filehander"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/biggjoker/file-manager/app/helper"
	"github.com/biggjoker/file-manager/g"
	"github.com/biggjoker/file-manager/protocol"
	"github.com/biggjoker/file-manager/zlog"
	"github.com/gin-gonic/gin"
)

func GetDic(c *gin.Context) {
	path, ok := c.GetQuery("path")
	if !ok {
		helper.JSONR(c, http.StatusBadRequest, "path no exit")
		zlog.Errorw("get dir faile", "dir", path)
		return
	}

	f, err := filehander.NewFounder(path)
	if err != nil {
		helper.JSONR(c, http.StatusInternalServerError, err)
		zlog.Errorw("get dir faile", "dir", path, "err", err)
		return
	}

	fileSlice, err := f.GetChilds()
	if err != nil {
		helper.JSONR(c, http.StatusInternalServerError, err)
		zlog.Errorw("get dir faile", "dir", path, "err", err)
		return
	}

	helper.JSONR(c, fileSlice)
}

func CreateDir(c *gin.Context) {
	var req protocol.CreateDirReq
	if err := c.Bind(&req); err != nil {
		helper.JSONR(c, http.StatusBadRequest, helper.INPUT_FORMATE_ERROR)
		return
	}
	_, err := filehander.CreateFounder(req.Path)
	if err != nil {
		helper.JSONR(c, http.StatusInternalServerError, err)
		zlog.Errorw("create dir faile", "dir", req.Path, "err", err)
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
		zlog.Errorw("delete dir faile", "dir", req.DeletePath, "err", err)
		return
	}

	f, err := filehander.NewFounder(req.DeletePath)
	if err != nil {
		helper.JSONR(c, http.StatusInternalServerError, err)
		zlog.Errorw("delete dir faile", "dir", req.DeletePath, "err", err)
		return
	}

	if err := f.Delete(req.force == "TRUE"); err != nil {
		helper.JSONR(c, http.StatusInternalServerError, err)
		zlog.Errorw("delete dir faile", "dir", req.DeletePath, "err", err)
		return
	}

	helper.JSONR(c, "delete dir success!")
}

func RenameDir(c *gin.Context) {
	var req protocol.RenameReq
	if err := c.Bind(&req); err != nil {
		helper.JSONR(c, http.StatusBadRequest, helper.INPUT_FORMATE_ERROR)
		zlog.Errorw("rename dir faile", "dir", req.New, "err", err)
		return
	}

	f, err := filehander.NewFounder(req.Old)
	if err != nil {
		helper.JSONR(c, http.StatusInternalServerError, err)
		zlog.Errorw("rename dir faile", "dir", req.New, "err", err)
		return
	}

	if err := f.Rename(req.New); err != nil {
		helper.JSONR(c, http.StatusInternalServerError, err)
		zlog.Errorw("rename dir faile", "dir", req.New, "err", err)
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
