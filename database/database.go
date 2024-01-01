package database

import (
	"database/sql"
	"task-5-pbi-btpns-almas/helper"
	"time"
)

func Database() *sql.DB {
	db, err := sql.Open("mysql", "root:Bismill@h1409@tcp(localhost:3307)/btpns_user_service")
	helper.PanicError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}