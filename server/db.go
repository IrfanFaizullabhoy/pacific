package main;

import (
	_ "github.com/lib/pq"

)

func ConnectToPG(dbName string) *sql.DB {
	db, err := sql.Open("postgres", "postgres://pacific:password@"+os.Getenv("DB_PORT_5432_TCP_ADDR")+"/objects?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func SetupTable(db *sql.DB, tableName string) {
	_, err := db.Exec("CREATE TABLE objects (name text PRIMARY KEY, size integer)")
	if err != nil {
		log.Printf("Error inserting into DB: %+v", err)
		return
	}
}
