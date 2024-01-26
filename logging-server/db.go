package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	pb "github.com/markbekhet/central-logging/log"
)

var db = &sql.DB{}

func ConnectDB(host, user, password, dbname string, port int) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dbCopy, err := sql.Open("postgres", connStr)
	db = dbCopy
	if err != nil {
		log.Fatal(err)
	}
}

func CreateHTTPLogTable() error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS HTTPLOG(
		id SERIAL PRIMARY KEY,
		statusCode INT,
		userIP VARCHAR(255),
		serverHost VARCHAR(255),
		day INT,
		month INT,
		year INT,
		hour INT,
		minute INT,
		second INT,
		nano INT
	)`)
	return err
}

func AddHTTPLOG(log *pb.HTTPLog) error {
	_, err := db.Exec(`
	INSERT INTO HTTPLOG 
	(statusCode, userIP, serverHost, day, month, year, hour, minute, second, nano)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		log.GetStatusCode(), log.GetUserIP(), log.GetServerHost(), log.Date.GetDay(),
		log.Date.GetMonth(), log.Date.GetYear(), log.Date.GetHour(), log.Date.GetMinute(),
		log.Date.GetSecond(), log.Date.GetNano())
	return err
}

func CloseDB() {
	db.Close()
}
