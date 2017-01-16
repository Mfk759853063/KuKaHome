package main

func main() {
	db := OpenSQL()
	defer CloseSQL(db)
	TestDB(db)
	Run()
}
