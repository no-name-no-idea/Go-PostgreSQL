package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0000"
	dbname   = "no_idea"
)

type User struct {
	Id   int
	Name string
	Idea string
}

func main() {
	r := gin.Default()

	r.GET("/User", GetData)

	r.Run()
}

func GetData(c *gin.Context) {

	gSlice := User{}
	gData := make([]User, 0)

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check database
	err = db.Ping()
	CheckError(err)

	// select
	selectStmt := `select * from User`
	rows, e := db.Query(selectStmt)
	CheckError(e)

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&gSlice.Id, &gSlice.Name, &gSlice.Idea)
		gData = append(gData, gSlice)
		CheckError(err)

	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    gData,
	})
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
