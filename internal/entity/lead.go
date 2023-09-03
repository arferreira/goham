package entity

import (
	"errors"
	"time"

	"github.com/arferreira/goham/pkg/entity"
)

var (
	ErrorIdIsRequired    = errors.New("Ops! ID is required")
	ErrorNameIsRequired  = errors.New("Ops! Name is required")
	ErrorEmailIsRequired = errors.New("Ops! Email is required")
)

type Lead struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewLead(name, email, phone string) (*Lead, error) {
	lead := &Lead{
		ID:        entity.NewID(),
		Name:      name,
		Email:     email,
		Phone:     phone,
		CreatedAt: time.Now(),
	}
	err := lead.Validate()
	if err != nil {
		return nil, err
	}
	return lead, nil
}

func (l *Lead) Validate() error {
	if l.ID.String() == "" {
		return ErrorIdIsRequired
	}
	if _, err := entity.ParseID(l.ID.String()); err != nil {
		return err
	}
	if l.Name == "" {
		return ErrorNameIsRequired
	}
	if l.Email == "" {
		return ErrorEmailIsRequired
	}
	return nil
}
