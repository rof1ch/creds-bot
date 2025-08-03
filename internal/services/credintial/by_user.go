package credintialservice

import "passwordbot/internal/domain/models"

func (s *CredintialService) ByUserId(userId int64) ([]models.Credintial, error) {
    return s.repo.ByUserId(userId)
}