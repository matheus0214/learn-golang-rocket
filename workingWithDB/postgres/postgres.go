package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func PostgresSql() {
	url := "postgres://user:password@localhost:5432/tests"
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	if err := conn.Ping(context.Background()); err != nil {
		panic(err)
	}

	query := "create table foo (id bigserial primary key, bar varchar(255))"

	if _, err := conn.Exec(context.Background(), query); err != nil {
		panic(err)
	}

	query = "insert into foo (bar) values($1)"
	if _, err := conn.Exec(context.Background(), query, "this is a test"); err != nil {
		panic(err)
	}

	var id int64
	var bar string
	err = conn.QueryRow(context.Background(), "select id, bar from foo where bar=$1", "this is a test").Scan(&id, &bar)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(id, bar)
}
