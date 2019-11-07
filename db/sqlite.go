package db

import (
	"github.com/catcherwong/rest-api-sample/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var SQLiteDb *gorm.DB

func init() {

	db, err := gorm.Open("sqlite3", config.AppCfg.DB)
	if err != nil {
		panic("failed to connect database")
	}

	db.LogMode(true)

	SQLiteDb = db

	SQLiteDb.DB().SetMaxIdleConns(100)
	SQLiteDb.DB().SetMaxIdleConns(150)

	cSql := `CREATE TABLE IF NOT EXISTS userinfo(
   id INTEGER PRIMARY KEY,
   name TEXT,
   gender INTEGER 
);
`

	SQLiteDb.Exec(cSql)
}
