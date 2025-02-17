package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func MySql() {
	db, err := sql.Open("mysql", "root:password@/tests")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	query := "create table foo (id bigint auto_increment primary key, bar varchar(255))"
	if _, err := db.Query(query); err != nil {
		panic(err)
	}

	query = "insert into foo (bar) values(?)"
	if _, err := db.Query(query, "This is a teste"); err != nil {
		panic(err)
	}

	query = "select * from foo where bar = ?"

	type Foo struct {
		ID  int64
		Bar string
	}

	var f Foo

	if err := db.QueryRow(query, "This is a teste").Scan(&f.ID, &f.Bar); err != nil {
		panic(err)
	}

	fmt.Printf("%#+v\n", f)
}
