package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDB() {
	os.Remove("gastos.db")
	db, err := sql.Open("sqlite3", "./gastos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `create table gastos (id integer not null primary key, fecha text, concepto text, cantidad real);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	} else {
		fmt.Println("Tabla gastos creada")
	}
}
