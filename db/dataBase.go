package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase interface {
	Connect(login, password string) error
	GetDB() *gorm.DB
}

type PostgreSQL struct {
	Host   string
	Port   string
	Dbname string
	db     *gorm.DB
}

func CreatePostgresDataBase(host, port, dbname string) *PostgreSQL {
	return &PostgreSQL{
		Host:   host,
		Port:   port,
		Dbname: dbname,
	}
}

func (p *PostgreSQL) Connect(login, password string) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.Host, p.Port, login, password, p.Dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database with GORM: %v", err)
	}

	p.db = db
	fmt.Println("Successfully connected to database (GORM)")
	return nil
}

func (p *PostgreSQL) GetDB() *gorm.DB {
	return p.db
}
