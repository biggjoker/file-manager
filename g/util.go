package g

import "strings"

// todo  某些关键文件不能访问
func IsPermissionFile(fileName string) bool {
	return true
}

func GetFullPath(path string) string {
	if path == "/" {
		return Config().BaseDir
	}
	return Config().BaseDir + path
}

func GetPathFileName(path string) (name string) {
	index := strings.LastIndex(path, "/")
	if index != -1 {
		name = path[index:]
	} else {
		name = path
	}
	return
}

func GetRelativePath(path string) string {
	if len(path) <= len(Config().BaseDir) {
		return ""
	}
	return path[len(Config().BaseDir):]
}
