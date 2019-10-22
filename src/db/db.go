package db

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "root"
	DBNAME   = "dbasego"
)

func Open() (*sql.DB, error) {
	//
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatalf("database.Connect ERROR: %s", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db, nil
}
