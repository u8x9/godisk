package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash"
	"io"
	"os"
	"path/filepath"
)

type Sha1Stream struct {
	Sha1 hash.Hash
}

func (obj *Sha1Stream) Update(data []byte) {
	if obj.Sha1 == nil {
		obj.Sha1 = sha1.New()
	}
	obj.Sha1.Write(data)
}

func Sha1(data []byte) string {
	s := sha1.New()
	s.Write(data)
	return hex.EncodeToString(s.Sum([]byte("")))
}

func FileSha1(file *os.File) string {
	s := sha1.New()
	io.Copy(s, file)
	return hex.EncodeToString(s.Sum(nil))
}

func MD5(data []byte) string {
	m := md5.New()
	m.Write(data)
	return hex.EncodeToString(m.Sum([]byte("")))
}

func FileMD5(file *os.File) string {
	m := md5.New()
	io.Copy(m, file)
	return hex.EncodeToString(m.Sum(nil))
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetFileSize(filename string) (result int64) {
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}
