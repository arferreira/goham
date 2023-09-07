package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Antonio", "arfs.antonio@gmail.com", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Antonio", user.Name)
	assert.Equal(t, "arfs.antonio@gmail.com", user.Email)
}

func TestValidatePassword(t *testing.T) {
	user, err := NewUser("Antonio", "arfs.antonio@gmail.com", "123456")

	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, user.Password, "123456")
}
