package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectSqlite() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}

	createTableSql := `
        CREATE TABLE foo (
            id integer not null primary key,
            name text
        )
    `

	res, err := db.Exec(createTableSql)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.RowsAffected())

	insertSql := `
        INSERT INTO foo (id, name) values (1, "Matheus")
    `

	res, err = db.Exec(insertSql)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.RowsAffected())

	type User struct {
		ID   int64
		Name string
	}

	var u User
	querySql := "SELECT * FROM foo WHERE id = ?;"

	err = db.QueryRow(querySql, 1).Scan(&u.ID, &u.Name)
	if err != nil {
		panic(err)
	}

	fmt.Println(u)
}
