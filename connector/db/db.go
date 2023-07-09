package db

import (
	"Message-Listener-Service/commons"
	"database/sql"
	_ "github.com/lib/pq"
)

type DataBaseConnection struct {
	db *sql.DB
}

func NewDataBaseConnection() (*DataBaseConnection,error) {
	
	// create db connection
	db,err := sql.Open("postgres", "postgres://postgres:Welcome4$@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return nil,err
	}

	return &DataBaseConnection{db: db},nil
}

func (dbc *DataBaseConnection) Close() error {
	
	// close db connection
	err := dbc.db.Close()
	if err != nil {
		return err
	}

	return nil
}

func (dbc *DataBaseConnection) GetDB() *sql.DB {
	return dbc.db
}

func InsertMessage(message commons.KafkaMessage) error {

	// insert message into db

	return nil
}

func InsertTransaction(transaction commons.Transaction) error {

	// insert transaction into db

	return nil
}

func InsertMerchant(merchant commons.Merchant) error {

	// insert merchant into db

	return nil
}

