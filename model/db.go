package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"tiramisu/config"
)

var (
	db *sqlx.DB
)

func init() {
	err := openDB(config.DB_HOST, config.DB_USER, config.DB_PASS, config.DB_NAME, config.DB_PORT)
	if err != nil {
		log.Fatalln("open db error: ", err)
	}
}

func openDB(dbhost, dbuser, dbpass, dbname, dbport string) error {
	var err error
	db, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpass, dbhost, dbport, dbname))
	if err != nil {
		return err
	}
	err = db.Ping()
	return err
}
