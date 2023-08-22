package repository

import (
	"context"
	"proyectoort/utils/entity"
)

const (
	qryInsertAudit = `
		INSERT INTO AUDITORIA (formulario_id, version, fecha) VALUES (?,?,?,?);`

	qryGetAuditByVersion = `
		SELECT
			id,
			formulario_id, 
			version, 
			fecha,
		FROM AUDITORIA
		WHERE version = ?;`

	qryGetAuditByDate = `
		SELECT
			id,
			estado,
			observaciones,
			version,
			fecha,
		FROM AUDITORIA
		WHERE fecha = ?;`

	qryGetAuditForm = `
		SELECT f.id,f.nombre,f.informacion,f.fecha, a.version as version
		FROM FORMULARIO f INNER JOIN AUDITORIA a
		ON f.id=a.id
		WHERE f.id=a.frmulario_id`
)

func (r *repo) SaveAudit(ctx context.Context, formulario_id int, version string, fecha string) error {
	_, err := r.db.ExecContext(ctx, qryInsertAudit, formulario_id, version, fecha)

	return err

}
func (r *repo) GetAuditByVersion(ctx context.Context, version string) (*entity.Auditoria, error) {
	a := &entity.Auditoria{}
	err := r.db.GetContext(ctx, a, qryGetAuditByVersion, version)
	if err != nil {
		return nil, err
	}

	return a, nil
}
func (r *repo) GetAuditByDate(ctx context.Context, fecha string) (*entity.Auditoria, error) {
	a := &entity.Auditoria{}
	err := r.db.GetContext(ctx, a, qryGetAuditByDate, fecha)
	if err != nil {
		return nil, err
	}

	return a, nil
}
func (r *repo) GetAuditFormVersion(ctx context.Context, version string) (*entity.Auditoria, error) {
	a := &entity.Auditoria{}
	err := r.db.GetContext(ctx, a, qryGetAuditForm, version)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (r *repo) GetAuditFormDate(ctx context.Context, fecha string) (*entity.Auditoria, error) {
	a := &entity.Auditoria{}
	err := r.db.GetContext(ctx, a, qryGetAuditForm, fecha)
	if err != nil {
		return nil, err
	}

	return a, nil
}
