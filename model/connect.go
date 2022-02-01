package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dsn   = "super:YOUR_PASSWD@tcp(localhost:3306)/StoreOL"
	db, _ = sql.Open("mysql", dsn)
)
