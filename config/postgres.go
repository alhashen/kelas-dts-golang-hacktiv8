package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Username string
	Password string
	Port     string
	Address  string
	Database string

	DB *sql.DB
}

type Psqldb struct {
	*Postgres
}

var PSQL *Psqldb

func InitProgress() error {
	PSQL = new(Psqldb)

	PSQL.Postgres = &Postgres{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Address:  os.Getenv("POSTGRES_ADDRESS"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	err := PSQL.Postgres.OpenConnection()
	if err != nil {
		return err
	}

	err = PSQL.DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) OpenConnection() error {

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		p.Address, p.Port, p.Username, p.Database)
	dbConnection, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	p.DB = dbConnection

	err = p.DB.Ping()
	if err != nil {
		return err
	}

	fmt.Println("connected to database:", p.Database)

	return nil
}
