package credintialstorage

import (
	"passwordbot/internal/domain/models"
	"passwordbot/pkg/logger/sl"
)

func (r *CredintialStorage) Delete(credId uint) error {
	log := r.log.With(sl.Op(op, "Delete"))

	err := r.db.Where("id = ?").Delete(&models.Credintial{}).Error

	if err != nil {
		log.Error(ErrDelete.Error(), sl.Err(err))
		return ErrDelete
	}

	return nil
}
