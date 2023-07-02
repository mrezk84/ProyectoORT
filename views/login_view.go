package views

import (
	"fmt"
	"log"
)

func DisplayLoginView() (string, string) {
    var username, password string

    fmt.Print("Ingrese usuario: ")
    _, err := fmt.Scan(&username)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print("Ingrese password: ")
    _, err = fmt.Scan(&password)
    if err != nil {
        log.Fatal(err)
    }

    return username, password
}