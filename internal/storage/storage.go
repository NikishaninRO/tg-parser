package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New(pgCfg string) (*Storage, error) {
	const op = "storage.New"

	db, err := gorm.Open(postgres.Open(pgCfg), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
