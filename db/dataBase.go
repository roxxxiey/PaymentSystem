package db

import (
	"database/sql"
	"fmt"
)

type DataBase interface {
	Connect(login, password string) error
	Close() error
}

type PostgreSQL struct {
	Host   string
	Port   string
	Dbname string
	db     *sql.DB
}

func CreatePostgresDataBase(host string, port, dbname string) (*PostgreSQL, error) {

	return &PostgreSQL{
		Host:   host,
		Port:   port,
		Dbname: dbname,
	}, nil
}

func (p *PostgreSQL) Connect(login, password string) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", p.Host, p.Port, login, password, p.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("create connction to database failed: %v", err)
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return fmt.Errorf("ping database failed: %v", err)
	}
	p.db = db
	fmt.Println("Successfully connected to database")
	return nil
}

func (p *PostgreSQL) Close() error {
	if p.db != nil {
		err := p.db.Close()
		if err != nil {
			return fmt.Errorf("close database failed: %v", err)
		}
		p.db = nil
		fmt.Println("Successfully closed database")
	}
	return nil
}
