package models

import "time"

type Credintial struct {
	Id                uint
	Name              string
	LoginEncrypted    string
	PasswordEncrypted string
	KeyHash           string
	LoginNonce        string
	PasswordNonce     string
	Description       string
	TypeId            uint
	UserId            int64
	CreatedAt         time.Time
	UpdatedAt         time.Time

	Type Type `gorm:"foreignKey:TypeId"`
}

func (c *Credintial) TableName() string {
	return "credentials"
}
