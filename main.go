package main

import (
	"log"
	"net/http"

	"github.com/u8x9/godisk/handler"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/success", handler.UploadSuccessHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
