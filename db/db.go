package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/mustafakocatepe/find-gaming-friends/db/migrate/postgres"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB(driver, host, database, username, password string, port, maxOpenConnections int) (*sql.DB, error) {

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		username,
		password,
		host,
		port,
		database)

	db, err := sql.Open(driver, url)
	if err != nil {
		return nil, err
	}

	if host == "localhost" {
		err = postgres.Migrate(db)
		if err != nil {
			return nil, err
		}
	}

	if err := pingDatabase(db); err != nil {
		return nil, err
	}

	return db, nil
}

func pingDatabase(db *sql.DB) error {
	r := 3
	var err error
	for i := 0; i < r; i++ {
		err = db.Ping()
		if err == nil {
			return nil
		}
		time.Sleep(1 * time.Second)
	}

	return err
}

func parseDSN(driver, host, database, username, password string, port int) (string, error) {

	switch driver {
	case "postgres":
		return postgreParseDSN(host, database, username, password, port), nil
	default:
		return "", nil
	}
}

func postgreParseDSN(host, database, username, password string, port int) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)
}
