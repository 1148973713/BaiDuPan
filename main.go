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
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)
	http.HandleFunc("/user/signup", handler.SignupHandler)
	http.HandleFunc("/user/signin", handler.SignInHandler)
	http.HandleFunc("/user/info", handler.UserInfoHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Print("Failed to start server,err:%s", err.Error())
	}
}
