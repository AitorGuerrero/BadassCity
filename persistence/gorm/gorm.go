package gorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type persistLayer struct {
	db gorm.Connection
	tx gorm.Transaction
}

func Get(user, password, dbName string) persistLayer {
	db, _ := gorm.Open("mysql", user + ":" + password + "@/" + dbName + "?charset=utf8&parseTime=True&loc=Local")
	db.DB()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return &persistLayer {
		db: db,
		tx: db.Begin(),
	}
}

func (l persistLayer) Persist (entity *interface{}) {
	l.db.Create(entity)
}

func (l persistLayer) Flush () {
	l.db.Commit()
}
