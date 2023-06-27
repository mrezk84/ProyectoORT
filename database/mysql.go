package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)


type LoginData struct {
    Username string
    Password string
}


func createConnection() (*sql.DB, error) {
    db, err := sql.Open("mysql", "admin:admin1234@tcp(ceaosadb.ccfkgyxkbhdd.us-east-1.rds.amazonaws.com:3306)/ceaosadb")
    if err != nil {
        return nil, err
    }
    return db, nil
}

func InsertLoginData(loginData LoginData) error {
    db, err := createConnection()
    if err != nil {
        return err
    }
    defer db.Close()

    query := "INSERT INTO usuarios (username, password) VALUES (?, ?)"
    _, err = db.Exec(query, loginData.Username, loginData.Password)
    if err != nil {
        return err
    }

    return nil
}