package database

import (
	"testing"

	"github.com/arferreira/goham/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	t.Run("should create a user", func(t *testing.T) {
		db, err := gorm.Open(sqlite.Open("file:memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}
		db.AutoMigrate(&entity.User{})
		user, _ := entity.NewUser("Antonio", "arfs.antonio@gmail.com", "123456")
		userDB := NewUser(db)
		err = userDB.CreateUser(user)
		assert.Nil(t, err)

		var userFromDB entity.User
		err = db.First(&userFromDB, "id = ?", user.ID).Error
		assert.Nil(t, err)

		assert.Equal(t, user.ID, userFromDB.ID)
		assert.Equal(t, user.Name, userFromDB.Name)
		assert.Equal(t, user.Email, userFromDB.Email)

	})
}

func TestGetUserByEmail(t *testing.T) {
	t.Run("should get a user by email", func(t *testing.T) {
		db, err := gorm.Open(sqlite.Open("file:memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}
		db.AutoMigrate(&entity.User{})
		user, _ := entity.NewUser("Antonio", "arfs.antonio@gmail.com", "123456")
		userDB := NewUser(db)
		err = userDB.CreateUser(user)
		assert.Nil(t, err)

		userFromDB, err := userDB.GetUserByEmail("arfs.antonio@gmail.com")
		assert.Nil(t, err)

		assert.Equal(t, user.ID, userFromDB.ID)
		assert.Equal(t, user.Name, userFromDB.Name)
		assert.Equal(t, user.Email, userFromDB.Email)
	})
}
