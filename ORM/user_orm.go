package ORM

import (
	"KuKaHome/Models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func CheckUserValid(u *Models.User) bool {
	var database *sql.DB
	database = Database
	if database == nil {
		log.Fatalln("没有数据库实例")
	}

	rows, err := database.Query("select * from user where username = ? ", u.Name)
	if err != nil {
		fmt.Println("查询数据出错", err)
	}
	defer rows.Close()

	data, err := GetRowData(rows)
	if err != nil {
		fmt.Println("GetRowdata 报错")
	}

	if data[0]["username"] == u.Name && data[0]["password"] == u.Password {
		fmt.Println("验证成功")
		return true
	} else {
		fmt.Println("验证失败")
		return false
	}

}
