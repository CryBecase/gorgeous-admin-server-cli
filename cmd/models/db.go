package models

import (
	"github.com/xormplus/xorm"
	"github.com/xormplus/xorm/log"
)

var (
	db *xorm.Engine
)

func initDB() {
	if db != nil {
		return
	}

	dsn := _config.DB.User + ":" + _config.DB.Password + "@tcp(" + _config.DB.Address + ")/" + _config.DB.Name + "?charset=utf8"

	var err error
	db, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.ShowSQL(false)
	db.SetLogLevel(log.LOG_OFF)

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
