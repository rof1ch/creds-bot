package credintialservice

import (
	"passwordbot/internal/domain/models"
)

func (s *CredintialService) ByTypeId(typeId uint) ([]models.Credintial, error) {
	return s.repo.List(typeId)
}
