package database

import (
	"github.com/arferreira/goham/internal/entity"
	"gorm.io/gorm"
)

type Lead struct {
	DB *gorm.DB
}

func NewLead(db *gorm.DB) *Lead {
	return &Lead{DB: db}
}

func (l *Lead) CreateLead(lead *entity.Lead) error {
	return l.DB.Create(lead).Error
}

func (l *Lead) GetLeadByEmail(email string) (*entity.Lead, error) {
	var lead entity.Lead
	err := l.DB.Where("email = ?", email).First(&lead).Error
	if err != nil {
		return nil, err
	}
	return &lead, nil
}
