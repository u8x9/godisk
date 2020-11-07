package main

import (
	"log"
	"net/http"

	"github.com/u8x9/godisk/handler"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/success", handler.UploadSuccessHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/query", handler.FileQueryHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
