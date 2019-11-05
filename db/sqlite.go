package db

import (
	"database/sql"
	"github.com/catcherwong/rest-api-sample/config"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("sqlite3", config.AppCfg.DB)

	DB = db

	DB.SetMaxIdleConns(100)
	DB.SetMaxOpenConns(150)

	cSql := `CREATE TABLE IF NOT EXISTS userinfo(
   id INTEGER PRIMARY KEY,
   name TEXT,
   gender INTEGER 
);
`

	DB.Exec(cSql)

	DB.Ping()

	if err != nil {
		panic(err)
	}
}
