package filehander

import (
	"fmt"
	"github.com/biggjoker/file-manager/g"
	"github.com/biggjoker/file-manager/protocol"
	"io/ioutil"
	"os"
)

var DEF_DIC_MODE = os.FileMode(0755)

type Founder struct {
	Path string
}

func NewFounder(path string) (*Founder, error) {
	fullPath := g.GetFullPath(path)
	_, err := os.Stat(fullPath)
	if err != nil {
		return nil, fmt.Errorf("path error")
	}

	return &Founder{Path: fullPath}, nil
}

func CreateFounder(path string) (*Founder, error) {
	fullPath := g.GetFullPath(path)

	if err := os.MkdirAll(fullPath, DEF_DIC_MODE); err != nil {
		return nil, fmt.Errorf("create file")
	}
	return &Founder{Path: fullPath}, nil
}

func (f *Founder) Rename(newPath string) error {
	newFullPath := g.GetFullPath(newPath)
	err := os.Rename(f.Path, newFullPath)
	if err == nil {
		f.Path = newPath
	}
	return err
}

func (f *Founder) GetChilds() ([]*protocol.FileInfo, error) {
	files, err := ioutil.ReadDir(f.Path)
	if err != nil {
		return nil, err
	}
	fileSlices := make([]*protocol.FileInfo, 0, 10)
	for _, file := range files {
		fileInfo := &protocol.FileInfo{
			Name:  g.GetRelativePath(f.Path) + "/" + file.Name(),
			IsDir: file.IsDir(),
			Size:  file.Size(),
		}
		fileSlices = append(fileSlices, fileInfo)
	}
	return fileSlices, nil
}

func (f *Founder) Delete(force bool) (err error) {
	if force {
		err = os.RemoveAll(f.Path)
	} else {
		err = os.Remove(f.Path)
	}
	if err != nil {
		f.Path = ""
	}
	return
}
