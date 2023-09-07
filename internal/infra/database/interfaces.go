package database

import "github.com/arferreira/goham/internal/entity"

type UserInterface interface {
	CreateUser(user *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
}

type LeadInterface interface {
	CreateLead(lead *entity.Lead) error
	GetLeadByEmail(email string) (*entity.Lead, error)
}
