package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xcheck/config"
)

var ConnectionString = ""

func getConn() *sql.DB {
	if ConnectionString == "" {
		ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=90s&collation=utf8_general_ci",
			config.CONFIG.DB.Username,
			config.CONFIG.DB.Pass,
			config.CONFIG.DB.Host,
			config.CONFIG.DB.Port,
			config.CONFIG.DB.DbName)
	}

	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func Version() string {
	var db = getConn()
	defer func() {
		e := db.Close()
		if e != nil {
			fmt.Println(e)
		}
	}()

	// Connect and check the server version
	var version string
	err := db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		fmt.Println(err)
		return "DB CONN ERR"
	}
	return version
}
