package main

import (
	"KuKaHome/ORM"
	"KuKaHome/server"
)

func main() {
	db := ORM.OpenSQL()
	defer ORM.CloseSQL(db)
	server.Run()
}
