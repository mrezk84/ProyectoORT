package models

type Document struct {
	ID         int64      `json:"id"`
	Formulario Formulario `json:"formulario"`
	Obra       Obra       `json:"obra"`
	Piso       Piso       `json:"piso"`
	Checks     []Check
	Status     string `json:"status"`
}
