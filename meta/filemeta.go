package meta

import (
	"sort"
	"time"

	"github.com/u8x9/godisk/db"
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
type ByUploadTime []*FileMeta

func (t ByUploadTime) Len() int {
	return len(t)
}
func (t ByUploadTime) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t ByUploadTime) Less(i, j int) bool {
	itime, _ := time.Parse("2006-01-02 15:04:05", t[i].UploadAt)
	jtime, _ := time.Parse("2006-01-02 15:04:05", t[j].UploadAt)
	return itime.Sub(jtime) < 0
}

var fileMetas FileMetaData

func init() {
	fileMetas = make(FileMetaData)
}

// UpdateFileMeta 新增/更新文件元信息
func UpdateFileMeta(fmeta *FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

func UpdateFileMetaDB(fmeta *FileMeta) (id int64, err error) {
	return db.OnFileUploadFinished(fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
}

// GetFileMeta 获取文件元信息
func GetFileMeta(fileSha1 string) *FileMeta {
	return fileMetas[fileSha1]
}
func GetFileMetaDB(filehash string) (*FileMeta, error) {
	tf, err := db.GetFileMeta(filehash)
	if err != nil {
		return nil, err
	}
	return &FileMeta{
		FileSha1: tf.FileHash,
		FileName: tf.Filename.String,
		Location: tf.FileAddr.String,
		FileSize: tf.FileSize.Int64,
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}
func GetLastFileMetas(limit int) []*FileMeta {
	s := make([]*FileMeta, 0, len(fileMetas))
	for _, v := range fileMetas {
		s = append(s, v)
	}
	sort.Sort(ByUploadTime(s))
	return s[0:limit]
}
func RemoveFileMeta(filehash string) *FileMeta {
	fileMeta, ok := fileMetas[filehash]
	if !ok {
		return nil
	}
	delete(fileMetas, filehash)
	return fileMeta
}
