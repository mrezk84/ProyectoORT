package controller

import (
	"proyectoort/model"

	"proyectoort/mysql"
)

func LoginController(username, password string) error {
    loginData := model.Usuario{
        Username: username,
        Password: password,
    }

    err := mysql.InsertLoginData(loginData)
    if err != nil {
        return err
    }

    return nil
}