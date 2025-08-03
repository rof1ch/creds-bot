package typestorage

import "errors"

var (
	ErrCreateType = errors.New("Ошибка при добавлении типа")
    ErrDeleteType = errors.New("Ошибка при удалении типа")
    ErrListTypes = errors.New("Ошибка при получении списка типов")
)