package typestorage

import (
	"passwordbot/internal/domain/models"
	"passwordbot/pkg/logger/sl"
)

func (r *TypeStorage) ById(typeID uint) (models.Type, error) {
	log := r.log.With(sl.Op(op, "ById"))

	var typeCred models.Type
	err := r.db.Where("id = ?", typeID).First(&typeCred).Error
	if err != nil {
		log.Error(ErrListTypes.Error(), sl.Err(err))
		return models.Type{}, ErrListTypes
	}

	return typeCred, nil
}
