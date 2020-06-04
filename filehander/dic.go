package filehander

import (
	"github.com/biggjoker/file-manager/g"
	"github.com/biggjoker/file-manager/protocol"
	"io/ioutil"
	"os"
)

type Founder struct {
	path string
}

func (f *Founder) SetPath(path string) bool {
	fullPath := g.GetFullPath(path)
	_, err := os.Stat(fullPath)
	if err == nil || os.IsExist(err) {
		return false
	}
	f.path = fullPath
	return true
}

func (f *Founder) Rename(newPath string) error {
	newFullPath := g.GetFullPath(newPath)
	err := os.Rename(f.path, newFullPath)
	if err != nil {
		f.path = newPath
	}
	return err
}

func (f *Founder) GetChilds()([]*protocol.FileInfo, error) {
	files, err := ioutil.ReadDir(f.path)
	if err != nil {
		return nil,err
	}
	fileSlices := make([]*protocol.FileInfo, 0, 10)
	for _, file := range files {
		fileInfo := &protocol.FileInfo{
			Name:  g.GetRelativePath(f.path) + "/" + file.Name(),
			IsDir: file.IsDir(),
			Size:  file.Size(),
		}
		fileSlices = append(fileSlices, fileInfo)
	}
	return fileSlices,nil
}

func (f *Founder) Delete(force bool)(err error) {
	if force {
		err = os.RemoveAll(f.path)
	} else {
		err = os.Remove(f.path)
	}
	if err != nil {
		f.path = ""
	}
	return
}

