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

func InsertGasto(fecha, concepto string, cantidad float64) {
	db, err := sql.Open("sqlite3", "./gastos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `insert into gastos (fecha, concepto, cantidad) values (?, ?, ?);`
	_, err = db.Exec(sqlStmt, fecha, concepto, cantidad)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	} else {
		fmt.Println("Gasto insertado")
	}
}

func GetGastos() (string, error) {
	db, err := sql.Open("sqlite3", "./gastos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from gastos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println("ID | Fecha | Concepto | Cantidad")
	for rows.Next() {
		var id int
		var fecha string
		var concepto string
		var cantidad float64
		rows.Scan(&id, &fecha, &concepto, &cantidad)
		fmt.Printf("%d | %s | %s | %.2f\n", id, fecha, concepto, cantidad)
	}
	return "Gastos", nil
}

func selectGastoId(id int) {
	db, err := sql.Open("sqlite3", "./gastos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from gastos where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println("ID | Fecha | Concepto | Cantidad")
	for rows.Next() {
		var id int
		var fecha string
		var concepto string
		var cantidad float64
		rows.Scan(&id, &fecha, &concepto, &cantidad)
		fmt.Printf("%d | %s | %s | %.2f\n", id, fecha, concepto, cantidad)
	}
}

func UpdateGasto(id int, fecha, concepto string, cantidad float64) (string, error) {
	db, err := sql.Open("sqlite3", "./gastos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `update gastos set fecha = ?, concepto = ?, cantidad = ? where id = ?;`
	_, err = db.Exec(sqlStmt, fecha, concepto, cantidad, id)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return "Error al actualizar el gasto", err
	} else {
		fmt.Println("Gasto actualizado")
	}
	return "Gasto actualizado", nil
}

func DeleteGasto(id int) (int, error) {
	db, err := sql.Open("sqlite3", "./gastos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `delete from gastos where id = ?;`
	_, err = db.Exec(sqlStmt, id)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return id, err
	} else {
		fmt.Println("Gasto borrado")
	}
	return id, nil
}

func InsertIngreso(fecha, concepto string, cantidad float64) (string, error) {
	db, err := sql.Open("sqlite3", "./gastos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `insert into ingresos (fecha, concepto, cantidad) values (?, ?, ?);`
	_, err = db.Exec(sqlStmt, fecha, concepto, cantidad)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return "", err
	} else {
		fmt.Println("Ingreso insertado")
		return "Ingreso insertado", nil
	}
}

func GetIngresos() (string, error) {
	db, err := sql.Open("sqlite3", "./gastos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from ingresos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println("ID | Fecha | Concepto | Cantidad")
	for rows.Next() {
		var id int
		var fecha string
		var concepto string
		var cantidad float64
		rows.Scan(&id, &fecha, &concepto, &cantidad)
		fmt.Printf("%d | %s | %s | %.2f\n", id, fecha, concepto, cantidad)
	}
	return "Ingresos", nil
}

func selectIngresoId(id int) (string, error) {
	db, err := sql.Open("sqlite3", "./gastos.db")
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer db.Close()
	rows, err := db.Query("select * from ingresos where id = ?", id)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer rows.Close()
	fmt.Println("ID | Fecha | Concepto | Cantidad")
	for rows.Next() {
		var id int
		var fecha string
		var concepto string
		var cantidad float64
		rows.Scan(&id, &fecha, &concepto, &cantidad)
		fmt.Printf("%d | %s | %s | %.2f\n", id, fecha, concepto, cantidad)
	}
	return "Ingreso", nil
}

func UpdateIngreso(id int, fecha, concepto string, cantidad float64) (string, error) {
	db, err := sql.Open("sqlite3", "./gastos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `update ingresos set fecha = ?, concepto = ?, cantidad = ? where id = ?;`
	_, err = db.Exec(sqlStmt, fecha, concepto, cantidad, id)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return "Error al actualizar el ingreso", err
	} else {
		fmt.Println("Ingreso actualizado")
	}
	return "Ingreso actualizado", nil
}

func DeleteIngreso(id int) (int, error) {
	db, err := sql.Open("sqlite3", "./gastos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `delete from ingresos where id = ?;`
	_, err = db.Exec(sqlStmt, id)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return id, err
	} else {
		fmt.Println("Ingreso borrado")
	}
	return id, nil
}
