package dto

import "passwordbot/internal/domain/models"

type CredintialInput struct {
	Name        string
	Login       string
	Password    string
	Key         string
	Description string
	TypeId      uint
	UserId      int64
}

type Credintial struct {
	Name        string
	Login       string
	Password    string
	Description string

	Type models.Type
}