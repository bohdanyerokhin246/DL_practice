package main

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1"
	dbname   = "DL_test"
)

func insertRecord(name, address string) error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	query := `INSERT INTO "checkAddresses" (name, addresses) VALUES ($1, $2)`
	_, err = db.Exec(query, name, address)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			log.Printf("PostgreSQL error with code %s: %v\n", pgErr.Code, pgErr.Message)
		}
		return err
	}
	return nil
}

func getRecordsByName(name string) ([]Address, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `SELECT id, name, addresses FROM "checkAddresses" WHERE name = $1`
	rows, err := db.Query(query, name)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			log.Printf("PostgreSQL error with code %s: %v\n", pgErr.Code, pgErr.Message)
		}
		return nil, err
	}
	defer rows.Close()

	addresses := []Address{}
	for rows.Next() {
		var address Address
		err := rows.Scan(&address.ID, &address.Name, &address.Address)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}
func updateRecord(id, name, address string) error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	query := `UPDATE "checkAddresses" SET name = $1, addresses = $2 WHERE id = $3`
	_, err = db.Exec(query, name, address, id)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			log.Printf("PostgreSQL error with code %s: %v\n", pgErr.Code, pgErr.Message)
		}
		return err
	}

	return nil
}

func deleteRecord(id string) error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	query := `DELETE FROM "checkAddresses" WHERE id = $1`
	_, err = db.Exec(query, id)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			log.Printf("PostgreSQL ошибка с кодом %s: %v\n", pgErr.Code, pgErr.Message)
		}
		return err
	}

	return nil
}
