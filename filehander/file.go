package filehander

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/biggjoker/file-manager/g"
	"github.com/biggjoker/file-manager/protocol"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
)

//读写权限
var DEF_FILE_MODE = os.FileMode(0644)

type File struct {
	Path string
}

func NewFile(path string) (*File, error) {
	fullPath := g.GetFullPath(path)
	stat, err := os.Stat(fullPath)
	if err != nil {
		return nil, fmt.Errorf("path error")
	}

	if stat.IsDir() {
		return nil, fmt.Errorf("path error")
	}

	return &File{Path: fullPath}, nil
}

func CreateFile(path, content string) (*File, error) {
	fullPath := g.GetFullPath(path)
	_, err := os.Stat(fullPath)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	if err := ioutil.WriteFile(fullPath, []byte(content), DEF_FILE_MODE); err != nil {
		return nil, fmt.Errorf("create file")
	}
	return &File{Path: fullPath}, nil
}

func DeleteFile(path string) error {
	fullPath := g.GetFullPath(path)
	return os.Remove(fullPath)
}

func (f *File) ReadFile() ([]byte, error) {
	return ioutil.ReadFile(f.Path)
}

func (f *File) SaveFile(content []byte) error {
	return ioutil.WriteFile(f.Path, []byte(content), DEF_FILE_MODE)
}

func (f *File) RenameFile(newPath string) error {
	err := os.Rename(f.Path, g.GetFullPath(newPath))
	if err == nil {
		f.Path = g.GetFullPath(newPath)
	}
	return err
}

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
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

func GetFileInfoByPath(path string) (protocol.FileInfoRep, error) {

	res := protocol.FileInfoRep{}
	h := md5.New()

	f, err := os.Open(path)
	if err != nil {
		return res, err
	}
	defer f.Close()

	_, err = io.Copy(h, f)
	if err != nil {
		return res, err
	}
	res.Md5 = hex.EncodeToString(h.Sum(nil))
	fi, err := f.Stat()
	if err != nil {
		return res, err
	}
	res.Timestamp = fi.ModTime().Unix()

	return res, nil
}
