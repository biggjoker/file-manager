package protocol

type FileInfo struct {
	Name  string `json:"name"`
	IsDir bool   `json:"is_dir"`
	Size int64   `json:"size"`
}

type SaveFileReq struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type CreateDirReq struct {
	Path string `json:"dir_path"`
}

type RenameReq struct {
	Old string `json:"oldname"`
	New string `json:"newname"`
}

type FileInfoRep struct {
	Md5       string `json:"md5"`
	Timestamp int64  `json:"timestamp"`
}