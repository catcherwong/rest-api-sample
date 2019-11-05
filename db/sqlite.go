package db

import (
	"database/sql"
	"github.com/catcherwong/rest-api-sample/config"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func init() {
	DB, _ = sql.Open("sqlite3", config.AppCfg.DB)

	DB.SetMaxIdleConns(100)
	DB.SetMaxOpenConns(150)

	DB.Ping()

	if err != nil {
		panic(err)
	}
}
