package config

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error){
	conf := InitEnvs()
	connStr := conf.DatabaseDSN 

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db, nil
}