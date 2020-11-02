package handler

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
		newFile, err := os.Create("/tmp/" + head.Filename)
		if err != nil {
			showErr(w, err)
			return
		}
		defer newFile.Close()
		if _, err := io.Copy(newFile, file); err != nil {
			showErr(w, err)
			return
		}
		// 重定向在上传成功页面
		http.Redirect(w, r, "/file/upload/success", http.StatusFound)
	}

}

// UploadSuccessHandler 文件上传成功
func UploadSuccessHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("文件上传成功"))
}
