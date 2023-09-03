package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLead(t *testing.T) {
	lead, err := NewLead("Antonio", "arfs.antonio@gmail.com", "16892696132")

	assert.Nil(t, err)
	assert.NotNil(t, lead)
	assert.NotEmpty(t, lead.ID)
	assert.Equal(t, "Antonio", lead.Name)
	assert.Equal(t, "arfs.antonio@gmail.com", lead.Email)
	assert.Equal(t, "16892696132", lead.Phone)
}

func TestLeadWhenNameIsEmpty(t *testing.T) {
	_, err := NewLead("", "arfs.antonio@gmail.com", "16892696132")

	assert.NotNil(t, err)
	assert.Equal(t, ErrorNameIsRequired, err)
}

func TestLeadWhenEmailIsEmpty(t *testing.T) {
	_, err := NewLead("Antonio", "", "16892696132")

	assert.NotNil(t, err)
	assert.Equal(t, ErrorEmailIsRequired, err)
}
