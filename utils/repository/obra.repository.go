package repository

import (
	"context"
	"fmt"
	"proyectoort/utils/entity"

	"github.com/labstack/gommon/log"
)

const (
	qryInsertObra = `
		INSERT INTO OBRA (Nombre)
		VALUES (?);`

	qryGetObrabyID = `
		SELECT
			ID,
			Nombre
		FROM OBRA
		WHERE ID = ?;`

	qryGetObrabyName = `
		SELECT
			ID,
			Nombre
		FROM OBRA
		WHERE Nombre = ?;`

	qryGetObras = `
		SELECT
			ID,
			Nombre
		FROM OBRA`

	qryGetObraDePiso = `
		SELECT
		o.ID ,
		o.Nombre
		FROM OBRA o 
		inner join OBRA_PISOS op on o.ID = op.obra_id
		WHERE op.piso_id = %v;`

	qryGetPisosDeObra = `
		SELECT
		p.id,
		p.numero
		FROM PISO p
		inner join OBRA_PISOS op on p.id = op.piso_id
		WHERE op.obra_id = ?;`

	qryUpdateObra = `
		update OBRA
	set Nombre = '%v'
	where ID = %v	
	`

	qryDeleteObra = `
		DELETE FROM OBRA where id = ?`
)

func (r *repo) SaveObra(ctx context.Context, nombre string) error {
	_, err := r.db.ExecContext(ctx, qryInsertObra, nombre)
	return err
}

func (r *repo) GetObras(ctx context.Context) ([]entity.Obra, error) {
	oo := []entity.Obra{}

	err := r.db.SelectContext(ctx, &oo, qryGetObras)
	if err != nil {
		return nil, err
	}

	return oo, nil
}

func (r *repo) GetObrabyName(ctx context.Context, name string) (*entity.Obra, error) {
	o := &entity.Obra{}
	err := r.db.GetContext(ctx, o, qryGetObrabyName, name)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (r *repo) GetObrabyID(ctx context.Context, obraID int64) (*entity.Obra, error) {
	o := &entity.Obra{}
	err := r.db.GetContext(ctx, o, qryGetObrabyID, obraID)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (r *repo) GetobraP(ctx context.Context, pisoID int64) (*entity.Obra, error) {
	obrap := &entity.Obra{}

	err := r.db.GetContext(ctx, &obrap, qryGetObraDePiso, pisoID)
	if err != nil {
		return nil, err
	}

	return obrap, nil

}

func (r *repo) GetPisosDeObra(ctx context.Context, obraID int64) ([]entity.Piso, error) {
	pp := []entity.Piso{}

	err := r.db.SelectContext(ctx, &pp, qryGetPisosDeObra, obraID)
	if err != nil {
		return nil, err
	}

	return pp, nil
}

func (r *repo) DeleteObra(ctx context.Context, obraID int64) error {
	_, err := r.db.ExecContext(ctx, qryDeleteObra, obraID)
	return err
}

func (r *repo) UpdateObra(ctx context.Context, obraID int64, nombre string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Println(err)
		log.Error(err.Error())
		return err
	}
	_, err = tx.ExecContext(ctx, fmt.Sprintf(qryUpdateObra, nombre, obraID))
	if err != nil {
		fmt.Println(err)
		fmt.Println("qdas")
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}
