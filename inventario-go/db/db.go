package db

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("sqlite3", "./inventario.db")
	if err != nil {
		log.Fatal(err)
	}
	createTables()
}

func createTables() {

	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS usuarios (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nombre TEXT,
		email TEXT UNIQUE,
		clave TEXT
	);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS productos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nombre TEXT,
		precio REAL,
		cantidad INTEGER
	);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS proveedores (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nombre TEXT,
		contacto TEXT
	);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS ventas (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		producto_id INTEGER,
		cantidad INTEGER,
		fecha TEXT,
		FOREIGN KEY(producto_id) REFERENCES productos(id)
	);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS configuraciones (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		clave TEXT,
		valor TEXT
	);`)
	if err != nil {
		log.Fatal(err)
	}
}

func Close() {
	DB.Close()
}
