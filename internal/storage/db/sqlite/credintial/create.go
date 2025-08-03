package credintialstorage

import (
	"passwordbot/internal/domain/models"
	"passwordbot/pkg/logger/sl"
)

func (r *CredintialStorage) Create(input models.Credintial) (uint, error) {
	log := r.log.With(sl.Op(op, "Create"))

	err := r.db.Create(&input).Error
	if err != nil {
		log.Error(ErrCreateCredintial.Error(), sl.Err(err))
		return 0, ErrCreateCredintial
	}

	return input.Id, nil
}
