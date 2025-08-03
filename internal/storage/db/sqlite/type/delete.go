package typestorage

import (
	"passwordbot/internal/domain/models"
	"passwordbot/pkg/logger/sl"
)

func (r *TypeStorage) Delete(credId uint) error {
	log := r.log.With(sl.Op(op, "Delete"))

	err := r.db.Where("id = ?", credId).Delete(&models.Type{}).Error
	if err != nil {
		log.Error(ErrDeleteType.Error(), sl.Err(err))
		return ErrDeleteType
	}

	return nil
}
