package model

import "fmt"

type Usuario struct {
	ID       int
	Username string
	Password string
	Rol      Rol
}

type UserModel struct {
	Users []Usuario
}

func (m *UserModel) GetUserByUsername(username string) (Usuario, error) {
	for _, user := range m.Users {
		if user.Username == username {
			return user, nil
		}
	}
	return Usuario{}, fmt.Errorf("Usuario no encontrado")
}
