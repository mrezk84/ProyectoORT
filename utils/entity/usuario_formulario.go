package entity

type UsuarioForm struct {
	ID           int64 `db:"id"`
	FormularioID int64 `db:"formulario_id"`
	UsuarioID    int64 `db:"usuario_id"`
}
