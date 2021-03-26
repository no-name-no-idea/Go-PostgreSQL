package main

import (
	"db"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	db.Init()

	// insert
	insertStmt := `insert into ideas ("id","name", "idea") values (0, 'no-name', 'no-idea')`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	// dynamic insert
	insertDynStmt := `insert into idea ("id","name", "idea") values($1, $2, $3)`
	_, e = db.Exec(insertDynStmt, 1, "no-name", "no-idea")
	CheckError(e)

	// update
	updateStmt := `update idea set "name" = $1, "idea" = $2 where "id" = $3`
	_, e = db.Exec(updateStmt, "Med", "hum....", 0)
	CheckError(e)

	// delete
	deleteStmt := `delete from idea where id = $1`
	_, e = db.Exec(deleteStmt, 1)
	CheckError(e)

	// select
	selectStmt := `select * from idea`
	rows, e := db.Query(selectStmt)
	CheckError(e)

	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var idea string

		err := rows.Scan(&id, &name, &idea)
		CheckError(err)

		fmt.Println(id, name, idea)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
