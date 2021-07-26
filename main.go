package main

import (
	handler "BaiDuPan/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.FileQueryHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Print("Failed to start server,err:%s", err.Error())
	}
}
