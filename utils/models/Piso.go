package models

type Piso struct {
	ID     int  `json:"id"`
	Numero int  `json:"numero"`
	Obra   Obra `json:"obra"`
}
