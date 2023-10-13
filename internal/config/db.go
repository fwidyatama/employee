package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"strconv"
)

var (
	mysqlLog = log.WithField("config", "DBConnector")
)

func GetDBInstance(cfg Config) *sql.DB {

	mysqlLog.Println("init postgresql instance")

	dbHost := cfg.DBHost
	port := cfg.DBPort
	dbUser := cfg.DBUser
	dbPassword := cfg.DBPassword
	dbName := cfg.DBName

	dbPort, _ := strconv.Atoi(port)

	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	log.Println("Success connect to db")

	return db

}
