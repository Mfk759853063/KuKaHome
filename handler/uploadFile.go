package handler

import (
	"KuKaHome/helper"
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("uploadfile")
	var buff bytes.Buffer
	fileSize, err := buff.ReadFrom(file)
	fmt.Println("file size is ", fileSize/1024)

	if err != nil {
		helper.WriteResponse(w, http.StatusBadRequest, -1, map[string]string{
			"error": err.Error(),
		})
		return
	}
	defer file.Close()

	directory, err := os.Getwd()
	directory = directory + "/files/"
	_, err = os.Stat(directory)
	if err != nil {
		os.Mkdir(directory, 0777)
	}
	f, err := os.OpenFile(directory+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	f.Write(buff.Bytes())
	helper.WriteResponse(w, http.StatusOK, 200, map[string]string{
		"message": "上传成功",
	})

}
