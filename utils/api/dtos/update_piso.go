package dtos

type UpdatePiso struct {
	PisoID int64 `param:"id"`
	Numero int   `json:"numero"`
}
