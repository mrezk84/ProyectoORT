package dtos

type RegisterCheck struct {
	ID            int    `json:"id_check" `
	Estado        string `json:"estado" `
	Observaciones string `json:"observaciones"`
	Version       string `json:"version"`
	Fecha         string `json:"fecha_control"`
}
