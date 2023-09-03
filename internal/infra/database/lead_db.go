package database

import "gorm.io/gorm"

type Lead struct {
	DB *gorm.DB
}

func NewLead(db *gorm.DB) *Lead {
	return &Lead{DB: db}
}
