package main

import (
	"fmt"

	_ "modernc.org/sqlite"

	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

func main() {
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
