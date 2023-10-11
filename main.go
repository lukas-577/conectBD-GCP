package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "34.27.242.101"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "bdtest"
)

var db *sql.DB

type Aeropuerto struct {
	Column1 string `json:"Column1"`
	Column2 string `json:"Column2"`
	Column3 string `json:"Column3"`
}

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexión exitosa")

	r := mux.NewRouter()
	r.HandleFunc("/aeropuertos", GetAeropuertos).Methods("GET")

	http.Handle("/", r)

	fmt.Println("Servidor escuchando en el puerto 8080...")
	http.ListenAndServe(":8080", nil)
}

func GetAeropuertos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * from Aeropuerto")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var aeropuertos []Aeropuerto

	for rows.Next() {
		var a Aeropuerto
		if err := rows.Scan(&a.Column1, &a.Column2, &a.Column3); err != nil {
			log.Fatal(err)
		}
		aeropuertos = append(aeropuertos, a)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Convertir el resultado en JSON
	jsonData, err := json.Marshal(aeropuertos)
	if err != nil {
		log.Fatal(err)
	}

	// Configurar las cabeceras HTTP para indicar que la respuesta es JSON
	w.Header().Set("Content-Type", "application/json")

	// Enviar el JSON como respuesta
	w.Write(jsonData)
}
