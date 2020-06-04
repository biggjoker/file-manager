package file

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/biggjoker/file-manager/zlog"
	"io"
	"os"

	"github.com/biggjoker/file-manager/app/protocol"
)

func GetFileInfoByPath(path string) (protocol.FileInfoRep, error) {

	res := protocol.FileInfoRep{}
	h := md5.New()

	f, err := os.Open(path)
	if err != nil {
		zlog.Errorw("open file error", "path", path, "err", err)
		return res, err
	}
	defer f.Close()

	_, err = io.Copy(h, f)
	if err != nil {
		zlog.Errorw("copy file error", "path", path, "err", err)
		return res, err
	}
	res.Md5 = hex.EncodeToString(h.Sum(nil))
	fi, err := f.Stat()
	if err != nil {
		zlog.Errorw("get file error", "path", path, "err", err)
		return res, err
	}
	res.Timestamp = fi.ModTime().Unix()

	return res, nil
}
