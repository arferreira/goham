package database

import (
	"testing"

	"github.com/arferreira/goham/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateLead(t *testing.T) {
	t.Run("should create a lead", func(t *testing.T) {
		db, err := gorm.Open(sqlite.Open("file:memory:"), &gorm.Config{})
		if err != nil {
			t.Error(err)
		}
		db.AutoMigrate(&entity.Lead{})

		lead, _ := entity.NewLead("Antonio", "arfs.antonio@gmail.com", "123456")
		leadDB := NewLead(db)
		err = leadDB.CreateLead(lead)
		assert.Nil(t, err)

		var leadFromDB entity.Lead
		err = db.First(&leadFromDB, "email = ?", leadFromDB.Email).Error
		assert.Nil(t, err)
		assert.Equal(t, leadFromDB.Name, lead.Name)
	})
}
