package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-gorp/gorp"
)

//database info
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0000"
	dbname   = "no_idea"
)

type DB struct {
	*sql.DB
}

var db *gorp.DbMap

func Init() {

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}

}

func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {

	db, err := sql.Open("postgres", dataSourceName)
	CheckError(err)

	err = db.Ping()
	CheckError(err)

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	return dbmap, nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
