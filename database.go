package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database interface {
	CreateAccount(*Account) error
	GetAccountByID(id int) (*Account, error)
	GetAccounts() ([]*Account, error)
	UpdateAccount(*Account) error
	DeleteAccount(id int) error
}

type PostgresDB struct {
	psql *sql.DB

}

func PostgresDBInstance() (*PostgresDB, error) {
// connStr := "postgresql://postgres:developer@localhost:5432/chopwella?schema=public" 
connStr := "user=postgres dbname=bankgoals password=developer sslmode=disable host=localhost port=5432"
db, err := sql.Open("postgres", connStr)

if err != nil {
	return nil, err
}
if err = db.Ping(); err != nil {
	return nil, err
}

return &PostgresDB{
	psql: db,}, nil
}

func (db *PostgresDB) Init() error {
	return db.createAccountTable()
}

func (db *PostgresDB) createAccountTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS accounts (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		number BIGINT UNIQUE,
		balance BIGINT,
		createdat TIMESTAMP
		)`

	_, err := db.psql.Exec(query)
	return err
}

func (pg *PostgresDB) CreateAccount(acc *Account) error {
	query := `INSERT INTO accounts (first_name, last_name, number, balance, createdat) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id`

	resp, err := pg.psql.Query(
		query, 
		acc.FirstName, 
		acc.LastName, 
		acc.Number, 
		acc.Balance,
		acc.CreatedAt,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

func (pg *PostgresDB) UpdateAccount(*Account) error {
	return nil
}

func (pg *PostgresDB) DeleteAccount(id int) error {
	return nil
}

func (pg *PostgresDB) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}

func (pg *PostgresDB) GetAccounts() ([]*Account, error) {
	rows, err := pg.psql.Query("SELECT * FROM accounts")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt,
		); 
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)   	
	}
	return accounts, nil
}