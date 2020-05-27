package main

import (
	"log"

	"github.com/icrowley/fake"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nleof/goyesql"
)

func main() {
	db, err := sqlx.Open("postgres", "dbname=goyesql_example sslmode=disable")
	if err != nil {
		log.Fatal("sqlx.Open error: ", err)
	}

	queries, err := goyesql.ParseFile("auth.sql")
	if err != nil {
		log.Fatal("goyesql.ParseFile error: ", err)
	}

	email := fake.EmailAddress()
	_, err = db.Exec(queries["add-user"], email, fake.SimplePassword())
	if err != nil {
		log.Fatal("db.Exec error: ", err)
	}
	log.Print("Added user: ", email)
}
