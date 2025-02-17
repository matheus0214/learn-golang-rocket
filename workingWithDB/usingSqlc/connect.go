package usingsqlc

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() {
	ctx := context.Background()

	url := "postgres://user:password@/tests"
	db, err := pgxpool.New(ctx, url)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	queries := New(db)
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(authors)

	author, err := queries.CreateAuthor(ctx, CreateAuthorParams{Name: "Matheus", Bio: pgtype.Text{String: "This is a bio", Valid: true}})
	if err != nil {
		panic(err)
	}

	fmt.Println(author)
}
