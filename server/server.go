package server

import (
	"log"
	"net/http"
)

func Run() {
	router := NewRouter()
	// err := http.ListenAndServeTLS(":8080", "./servercrts/2_qwertyuiop123456789.top.crt", "./servercrts/privite.key", router)
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatal("启动服务器失败 :", err)
	}
}
