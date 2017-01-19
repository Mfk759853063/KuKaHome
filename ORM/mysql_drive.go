package ORM

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Database *sql.DB

func OpenSQL() *sql.DB {
	database, err := sql.Open("mysql", "root:admin123@/gj.db")
	if err != nil {
		log.Fatalln("启动数据库失败")
	}
	err = database.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	Database = database
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

func GetRowData(rows *sql.Rows) ([]map[string]interface{}, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	datas := make([]map[string]interface{}, 0)
	scanArges := make([]interface{}, len(columns))
	values := make([]sql.RawBytes, len(columns))
	for i := range values {
		scanArges[i] = &values[i]
	}
	for rows.Next() {
		rows.Scan(scanArges...)
		data := make(map[string]interface{})
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			data[columns[i]] = value
		}
		datas = append(datas, data)
	}
	return datas, nil
}
