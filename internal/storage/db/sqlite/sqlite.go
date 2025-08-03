package sqlite

import (
	sqlite3 "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqlite(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite3.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
