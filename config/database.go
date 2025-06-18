package config

import (
	"database/sql"
	
	_ "assets/mysql"
)

func Default() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:P@ssw0rd@tcp(172.18.135.223)/uniform")

	if ENV == "development" {
		db, err = sql.Open("mysql", "root:P@ssw0rd@tcp(172.18.135.223)/uniform")
	} else if ENV == "production" {
		db, err = sql.Open("mysql", "ccp_own:Un1f0rM@ccP#@tcp(172.18.67.76)/uniform")
	}
	return db, err
}
