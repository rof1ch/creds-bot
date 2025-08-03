package typestorage

import (
	"passwordbot/internal/domain/models"
	"passwordbot/pkg/logger/sl"
)

func (r *TypeStorage) Create(input models.Type) (uint, error) {
	log := r.log.With(sl.Op(op, "Create"))

	err := r.db.Create(&input).Error
	if err != nil {
		log.Error(ErrCreateType.Error(), sl.Err(err))
		return 0, ErrCreateType
	}

	return input.Id, nil
}
