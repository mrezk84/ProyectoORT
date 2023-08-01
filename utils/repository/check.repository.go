package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertCheck = `
		INSERT INTO Check (estado, fecha_control, observaciones, version)
		VALUES (?,?,?,?);`

	qryGetCheckByVersion = `
		SELECT
			ID
			estado
			fecha_control
			observaciones
			version
		FROM CHECKS
		WHERE version = ?;`

	// qryInsert = `
	// 	INSERT INTO OBRA_PISOS (obra_id, piso_id) VALUES (:obra_id, :piso_id);`
)

func (r *repo) SaveCheck(ctx context.Context, estado, fecha, observaciones string, version int) error {
	_, err := r.db.ExecContext(ctx, qryInsertCheck, estado, fecha, observaciones, version)
	return err
}

func (r *repo) GetCheckByVersion(ctx context.Context, version int) (*entity.Check, error) {
	c := &entity.Check{}
	err := r.db.GetContext(ctx, c, qryGetCheckByVersion, version)
	if err != nil {
		return nil, err
	}

	return c, nil
}
