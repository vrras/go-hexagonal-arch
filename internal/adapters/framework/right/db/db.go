package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// connect to database
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}

	// test db connection
	if err = db.Ping(); err != nil {
		log.Fatalf("db ping failure: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (da Adapter) CloseDBConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

func (da Adapter) AddToHistory(answer int32, operation string) error {
	stmt, err := da.db.Prepare("INSERT INTO arith_history(answer, date, operation) VALUES(?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(answer, time.Now(), operation)
	if err != nil {
		return err
	}

	return nil
}
