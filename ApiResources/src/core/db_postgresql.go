package core

import (
	"fmt"
	"log"
	"os"
	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" 
)

type Conn_PostgreSQL struct {
	DB  *sql.DB
	Err string
}

func GetDBPool() *Conn_PostgreSQL {

	error := ""
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbSchema := os.Getenv("DB_SCHEMA")
	dbPort := os.Getenv("DB_PORT") 


	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbSchema)

	db, err := sql.Open("postgres", dsn) 

	if err != nil {
		error = fmt.Sprintf("error al abrir la base de datos: %v", err)
	}
	fmt.Println("Conexión establecida correctamente")
	db.SetMaxOpenConns(10)

	if err := db.Ping(); err != nil {
		db.Close()
		error = fmt.Sprintf("error al verificar la conexion a la base de datos: %v", err)
	}
	return &Conn_PostgreSQL{DB: db, Err: error}
}

func (conn *Conn_PostgreSQL) ExecutePreparedQuery(query string, values ...interface{}) (sql.Result, error) {
	stmt, err := conn.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error al preparar la consulta: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta preparada: %w", err)
	}

	return result, nil
}

func (conn *Conn_PostgreSQL) FetchRows(query string, values ...interface{}) (*sql.Rows) {
	rows, err := conn.DB.Query(query, values...)
	if err != nil {
		fmt.Printf("error al ejecutar la consulta SELECT: %v", err)
	}
	return rows
}
