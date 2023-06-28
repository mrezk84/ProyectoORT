package view

import (
	"fmt"
	"log"
)

func DisplayLoginView() (string, string) {
    var username, password string

    fmt.Print("Enter username: ")
    _, err := fmt.Scan(&username)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print("Enter password: ")
    _, err = fmt.Scan(&password)
    if err != nil {
        log.Fatal(err)
    }

    return username, password
}