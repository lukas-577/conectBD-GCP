package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "34.27.242.101"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "bdtest"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexión exitosa")

	// solo para mostrar xD

	rows, err := db.Query("SELECT * from Aeropuerto")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Printf("%-8s %-40s %s\n", "Column1", "Column2", "Column3")

	for rows.Next() {
		var column1, column2, column3 string
		if err := rows.Scan(&column1, &column2, &column3); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-8s %-40s %s\n", column1, column2, column3)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}
