package main

import (
	"Go-CloudDrive/handler"
	"fmt"
	"net/http"
)

/**
* @Description
* @Author YuCheng
* @Date 2023/1/13 1:23
**/

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server,err:%s", err.Error())
		return
	}

}
