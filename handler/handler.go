package handler

/**
* @Description
* @Author YuCheng
* @Date 2023/1/13 1:09
**/

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// UploadHandler 处理文件上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回上传html页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			_, err := io.WriteString(w, "internal server error")
			if err != nil {
				return
			}
			return
		}
		_, err = io.WriteString(w, string(data))
		if err != nil {
			return
		}
	} else if r.Method == "POST" {
		//接受文件流及存储到本地目录
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data,err:%s", err.Error())
			return
		}
		defer file.Close()
		newFile, err := os.Create("./tmp/" + head.Filename)
		if err != nil {
			fmt.Printf("Failed to create file,err:%s", err.Error())
			return
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save data into file,err:%s", err.Error())
			return
		}
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

// UploadSucHandler 上传已完成
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Upload finished")
	if err != nil {
		return
	}
}
