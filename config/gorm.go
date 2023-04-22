package config

import (
	"fmt"
	"os"
	"simpleapi/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Gorm struct {
	Username string
	Password string
	Port     string
	Address  string
	Database string

	DB *gorm.DB
}

type GormDB struct {
	*Gorm
}

var GORM *GormDB

func InitGorm() error {
	GORM = new(GormDB)

	GORM.Gorm = &Gorm{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Address:  os.Getenv("POSTGRES_ADDRESS"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	err := GORM.Gorm.OpenConnection()
	if err != nil {
		return err
	}

	return nil
}

func (p *Gorm) OpenConnection() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		p.Address, p.Port, p.Username, p.Database)

	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "showcase.",
			SingularTable: false,
		},
	})

	if err != nil {
		return err
	}

	p.DB = dbConnection

	err = p.DB.Debug().AutoMigrate(model.Book{})
	if err != nil {
		return err
	}

	fmt.Println("connected to database:", p.Database, "using gorm")

	return nil
}
