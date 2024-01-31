package databases

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

type LogData struct {
	ID       int64     `json:"id"`
	DH       time.Time `json:"dh"`
	MF       string    `json:"mf"`
	Argument string    `json:"argument"`
	Statut   string    `json:"statut"`
}

var db *sql.DB

func ConnectDataBase() {

	cfg := mysql.Config{
		User:   "root",
		Passwd: "1234",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "go",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}

func AddLog(log LogData) (int64, error) {
	result, err := db.Exec("INSERT INTO journal (dh, mf, argument, statut) VALUES (?,?,?,?)", log.DH, log.MF, log.Argument, log.Statut)
	if err != nil {
		return 0, fmt.Errorf("AddLog %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddLog %v", err)
	}

	return id, nil
}

func LastJournal() ([]LogData, error) {

	rows, err := db.Query("SELECT * FROM journal ORDER BY id DESC LIMIT 50;")
	if err != nil {
		return nil, fmt.Errorf("LastJournal %v", err)
	}

	defer rows.Close()

	var logs []LogData
	for rows.Next() {
		var (
			log     LogData
			dhBytes []byte
		)

		if err := rows.Scan(&log.ID, &dhBytes, &log.MF, &log.Argument, &log.Statut); err != nil {
			return nil, fmt.Errorf("LastJournal %v", err)
		}

		dhString := string(dhBytes)
		layout := "2006-01-02 15:04:05"
		log.DH, err = time.Parse(layout, dhString)
		if err != nil {
			return nil, fmt.Errorf("LastJournal%v", err)
		}

		logs = append(logs, log)
	}

	return logs, nil
}
