package database

import (
	"contacts/models"

	"gorm.io/gorm"
)

type ContactDB struct {
	DBClient interface{}
}

func (c *ContactDB) Create(contact *models.Contact) (id interface{}, err error) {
	c.DBClient.(*gorm.DB).AutoMigrate(&models.Contact{})

	result := c.DBClient.(*gorm.DB).Create(contact)
	if result.Error != nil {
		return nil, result.Error
	}
	return contact.ID, nil
}
