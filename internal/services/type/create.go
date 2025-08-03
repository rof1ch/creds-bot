package typeservice

import "passwordbot/internal/domain/models"

func (s *TypeService) Create(name, icon string, userId int64) error {
	typeCred := models.Type{
		Name:   name,
		Icon:   icon,
		UserId: userId,
	}
	_, err := s.repo.Create(typeCred)
	if err != nil {
		return err
	}
	return nil
}
