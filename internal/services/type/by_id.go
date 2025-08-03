package typeservice

import "passwordbot/internal/domain/models"

func (s *TypeService) ById(typeId uint) (models.Type, error) {
    return s.repo.ById(typeId)
}