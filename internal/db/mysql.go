package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlStorage struct {
	DB *sql.DB
}

type MysqlConfig struct {
	Username string
	Password string
	DbName   string
	Port     uint
	Host     string
}

func DbConnect(conf MysqlConfig) (*MysqlStorage, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.DbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQL connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping SQL database: %w", err)
	}

	return &MysqlStorage{
		DB: db,
	}, nil

}
