package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func OpenSQL() *sql.DB {
	database, err := sql.Open("mysql", "root:admin123@/gj_db")
	if err != nil {
		log.Fatalln("启动数据库失败")
	}
	err = database.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	return database
}

func CloseSQL(database *sql.DB) {
	err := database.Close()
	if err != nil {
		log.Fatalln("关闭数据库出错")
	}
}

func TestDB(database *sql.DB) {
	rows, err := database.Query("select * from product")
	if err != nil {
		log.Fatalln("查询出错", err)
	}
	var name string
	var id int
	for rows.Next() {
		err := rows.Scan(&name, &id)
		if err != nil {
			log.Fatalln(err)
		}
		log.Print(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
