package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)


func ConnectDB() (*sql.DB, error) {
    db, err := sql.Open("mysql", "admin:admin1234@tcp(ceaosadb.ccfkgyxkbhdd.us-east-1.rds.amazonaws.com:3306)/ceaosadb")
    if err != nil {
        return nil, err
    }
    return db, nil
}

