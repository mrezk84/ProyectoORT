package service

import (
	"context"
	"errors"
)

var (
	ErrAuditAlreadyExists = errors.New("la auditoria ya existe")
	ErrInvalidAudit       = errors.New("auditoria inválida")
	ErrAuditAlreadyAdded  = errors.New("ya se cuenta con version de la auditoria para el formulario")
	ErrAuditNotFound      = errors.New("error al generar auditoria")
)

func (s *serv) RegisterAudit(ctx context.Context, formularioID int64, version string, fecha string) error {

	f, _ := s.repo.GetFormsById(ctx, formularioID)
	if f != nil {
		return ErrAuditAlreadyExists
	}

	fv, _ := s.repo.GetFormByVersion(ctx, version)
	if fv != nil {
		return ErrAuditAlreadyAdded
	}

	fe, _ := s.repo.GetFormByDate(ctx, fecha)
	if fe != nil {
		return ErrFomEtapaAlreadyAdded
	}

	return s.repo.SaveAudit(ctx, f.ID, fv.Version, fe.Fecha)
}
