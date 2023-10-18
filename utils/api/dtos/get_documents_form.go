package dtos

type GetDocumentsForm struct {
	ID     int64 `param:"id"`
	UserID int64 `json:"usuario_id"`
}
