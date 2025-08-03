package typeservice

import "passwordbot/internal/domain/models"

func (s *TypeService) List(userID int64) ([]models.Type, error) {
	return s.repo.List(userID)
}
