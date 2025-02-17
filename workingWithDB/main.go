package main

import (
	"workingWithDB/postgres"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	postgres.PostgresSql()
}
