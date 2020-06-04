package filehander

import (
	"fmt"
	"github.com/biggjoker/file-manager/g"
	"os"
)

type File struct {
	Path string
}

func NewFile(path string) (*File, error) {
	fullPath := g.GetFullPath(path)
	stat, err := os.Stat(fullPath)
	if err == nil ||
		stat.IsDir() {
		return nil, fmt.Errorf("path error")
	}
	return &File{Path: path}, nil
}

func CreateFounder(path string) (*Founder,error) {
	fullPath := g.GetFullPath(path)
	_, err := os.Stat(fullPath)
	if err == nil {
		return nil ,fmt.Errorf("error")
	}
	if os.IsExist(err) {
		return nil ,fmt.Errorf("path exit")
	}
	if err := os.MkdirAll(fullPath, 0755); err != nil {
		return nil ,fmt.Errorf("create file")
	}
	return &Founder{Path: path},nil
}