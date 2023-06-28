package model

import (
	"database/sql"
)


type Usuario struct {
	ID       int
	Username string
	Password string
	Rol      Rol
}
type UserModel struct {
	Users []Usuario
}

func (u *Usuario) Save(db *sql.DB) error {
    //Inserta el usuario y lo inserta en la base de datos
    _, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", u.Username, u.Password)
    return err
}

func GetUserByUsername(db *sql.DB, username string) (*Usuario, error) {
    //Selecciona en la base el usaurio y lo devuelve
    var user Usuario
    err := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
    if err != nil {
        return nil, err
    }
    return &user, nil
}
