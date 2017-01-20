package handler

import (
	"KuKaHome/ORM"
	"KuKaHome/helper"
	"bytes"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
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

	fileNames := strings.Split(handler.Filename, ".")
	var extension string
	var fileName string
	if len(fileNames) <= 1 {
		extension = handler.Filename
		fileName = handler.Filename
	} else {
		extension = fileNames[len(fileNames)-1]
		fileName = fileNames[0]
	}

	m := md5.New()
	m.Write([]byte(fileName))
	newPath := hex.EncodeToString(m.Sum(nil))
	destinPath := directory + newPath + "." + extension
	f, err := os.OpenFile(destinPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	f.Write(buff.Bytes())
	helper.WriteResponse(w, http.StatusOK, 200, map[string]string{
		"message":  "上传成功",
		"filePath": newPath + "." + extension,
	})
	// err = savePath2DB(handler.Filename, newPath)
	// if err != nil {
	// 	fmt.Println(err)
	// }

}

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	filePath := vars["imgPath"]
	destinPath, _ := os.Getwd()
	destinPath = destinPath + "/files/" + filePath
	_, err := os.Stat(destinPath)
	if err != nil {
		helper.WriteResponse(w, http.StatusOK, -1, map[string]string{
			"message":  "文件不存在",
			"filePath": destinPath,
		})
	} else {
		http.ServeFile(w, r, destinPath)
	}
}

// 保存文件路径到mysql
func savePath2DB(directory string, newPath string) error {
	var database *sql.DB
	database = ORM.Database
	if database == nil {
		fmt.Println("没有数据库实例")
		return errors.New("没有数据库实例")
	}

	// see http://studygolang.com/articles/3022
	tx, err := database.Begin()
	if err != nil {
		return err
	}
	_, err1 := tx.Exec("insert into files (oriPath,newPath) values (?,?)", directory, newPath)
	if err1 != nil {
		return err1
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
