package credintialstorage

import "errors"

var (
	ErrCreateCredintial = errors.New("Ошибка при добавлении записи")
	ErrCredNotFound     = errors.New("Данной записи не существует")
	ErrGetByIdCred      = errors.New("Не удалось получить данные по Id")
	ErrGetListByTypeId  = errors.New("Не удалось получить данные по Id типа")
	ErrDelete           = errors.New("Не удалось удалить запись")
)
