package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/u8x9/godisk/meta"
	"github.com/u8x9/godisk/util"
)

// showErr 在页面上显示错误信息
func showErr(w http.ResponseWriter, err error) {
	w.Write([]byte("内部服务错误:" + err.Error()))
}

//UploadHandler 文件上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// 返回上传html页面
		buf, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			showErr(w, err)
			return
		}
		w.Write(buf)
	} else if r.Method == http.MethodPost {
		// 接收文件流及存储到本地目录
		file, head, err := r.FormFile("file")
		if err != nil {
			showErr(w, err)
			return
		}
		defer file.Close()
		location := "/tmp/" + head.Filename
		fileMeta := meta.NewFileMeta(head.Filename, location)
		newFile, err := os.Create(location)
		if err != nil {
			showErr(w, err)
			return
		}
		defer newFile.Close()
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			showErr(w, err)
			return
		}
		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)
		fmt.Printf("%#v\n", fileMeta)
		// 重定向在上传成功页面
		http.Redirect(w, r, "/file/upload/success", http.StatusFound)
	}

}

// UploadSuccessHandler 文件上传成功
func UploadSuccessHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("文件上传成功"))
}

// GetFileMetaHandler 通过文件hash获取文件meta信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filehash := r.Form.Get("filehash")
	fileMeta := meta.GetFileMeta(filehash)
	buf, err := json.Marshal(fileMeta)
	if err != nil {
		showErr(w, err)
		return
	}
	w.Header().Add("content-type", "application/json;charset=utf-8")
	w.Write(buf)
}
