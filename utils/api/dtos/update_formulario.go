package dtos

type UpdateForm struct {
	FormID      int64  `param:"id"`
	Nombre      string `json:"nombre"`
	Informacion string `json:"informacion"`
}
