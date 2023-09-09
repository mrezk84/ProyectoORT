package entity

type UsuarioRol struct {
	ID     int `db:"id"`
	UserID int `db:"usuario_id"`
	RoleID int `db:"rol_id"`
}
