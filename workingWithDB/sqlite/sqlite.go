package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"

	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
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

func InMemory() {
	conn, err := sqlite.OpenConn(":memory:", sqlite.OpenReadWrite)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	err = sqlitex.ExecuteTransient(conn, "SELECT 'hello, world';", &sqlitex.ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			fmt.Println(stmt.ColumnText(0))
			return nil
		},
	})

	if err != nil {
		panic(err)
	}
}
