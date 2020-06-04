package dic

import (
	"github.com/biggjoker/file-manager/app/helper"
	"github.com/biggjoker/file-manager/filehander"
	"github.com/biggjoker/file-manager/g"
	"github.com/biggjoker/file-manager/protocol"
	"github.com/biggjoker/file-manager/zlog"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	uploadFileKey  = "file"
	uploadFilePath = "path"
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

	f, err := filehander.NewFile(path)
	if err != nil {
		helper.JSONR(c, http.StatusInternalServerError, err)
		zlog.Errorw("read file failed", "dir", path, "err", err)
		return
	}

	content,err := f.ReadFile()
	if err != nil {
		helper.JSONR(c, http.StatusInternalServerError, err)
		zlog.Errorw("read file failed", "dir", path, "err", err)
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

	f, err := filehander.NewFile(req.Name)
	if err != nil {
		helper.JSONR(c, http.StatusInternalServerError, err)
		zlog.Errorw("save file failed", "dir", req.Name, "err", err)
		return
	}

	err = f.SaveFile([]byte(req.Content))
	if err != nil {
		helper.JSONR(c, http.StatusInternalServerError, err)
		zlog.Errorw("save file failed", "dir", req.Name, "err", err)
		return
	}

	helper.JSONR(c, "save success")
}

func CreateFile(c *gin.Context) {
	var req protocol.SaveFileReq
	if err := c.Bind(&req); err != nil {
		helper.JSONR(c, http.StatusBadRequest, helper.INPUT_FORMATE_ERROR)
		return
	}

	_, err := filehander.CreateFile(req.Name,req.Content)
	if err != nil {
		helper.JSONR(c, http.StatusInternalServerError, err)
		zlog.Errorw("save file failed", "dir", req.Name, "err", err)
		return
	}

	helper.JSONR(c, "save success")
}

func DeleteFile(c *gin.Context) {
	req := struct {
		Name string `json:"name"`
	}{}
	if err := c.Bind(&req); err != nil {
		helper.JSONR(c, http.StatusBadRequest, helper.INPUT_FORMATE_ERROR)
		return
	}

	err := filehander.DeleteFile(req.Name)
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

	f, err := filehander.NewFile(req.Old)
	if err != nil {
		helper.JSONR(c, http.StatusInternalServerError, err)
		zlog.Errorw("rename file failed", "dir", req.Old, "err", err)
		return
	}
	err = f.RenameFile(req.New)
	if err != nil {
		zlog.Errorw("rename faile", "old", req.Old, "new", req.New, "err", err)
		helper.JSONR(c, http.StatusInternalServerError, err)
		return
	}
	helper.JSONR(c, "rename success!")
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
	dst := g.GetFullPath(path + header.Filename)

	if err := filehander.SaveUploadedFile(header, dst); err != nil {
		helper.JSONR(c, http.StatusInternalServerError, err)
		zlog.Errorw("upload file", "dir", path, "err", err)
	}
	res, err := filehander.GetFileInfoByPath(dst)
	if err != nil {
		helper.JSONR(c, http.StatusInternalServerError, err)
		zlog.Errorw("upload file", "dir", path, "err", err)
	}
	helper.JSONR(c, res)
}