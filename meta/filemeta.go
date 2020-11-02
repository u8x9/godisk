package meta

import (
	"time"
)

// FileMeta 文件元信息
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

func NewFileMeta(filename, location string) *FileMeta {
	return &FileMeta{
		FileName: filename,
		Location: location,
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}

type FileMetaData map[string]*FileMeta // key: FileSha1

var fileMetas FileMetaData

func init() {
	fileMetas = make(FileMetaData)
}

// UpdateFileMeta 新增/更新文件元信息
func UpdateFileMeta(fmeta *FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

// GetFileMeta 获取文件元信息
func GetFileMeta(fileSha1 string) *FileMeta {
	return fileMetas[fileSha1]
}
