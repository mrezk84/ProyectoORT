package controller

import (
	"proyectoort/database"
	"proyectoort/model"
)



func LoginController(username, password string) error {
    loginData := model.Usuario{
    	Username: username,
    	Password: password,
    }

    err := database.LoginData(loginData)
    if err != nil {
        return err
    }

    return nil
}

