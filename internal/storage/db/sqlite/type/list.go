package typestorage

import (
	"passwordbot/internal/domain/models"
	"passwordbot/pkg/logger/sl"
)

func (r *TypeStorage) List(userId int64) ([]models.Type, error) {
	log := r.log.With(sl.Op(op, "List"))

	var types []models.Type
	err := r.db.Where("user_id = ?", userId).Find(&types).Error
	if err != nil {
		log.Error(ErrListTypes.Error(), sl.Err(err))
		return nil, ErrListTypes
	}

	return types, nil
}
