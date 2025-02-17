package main

import (
	"workingWithDB/mysql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mysql.MySql()
}
