package interfaces

import (
	"contacts/models"
)

type IContact interface {
	Create(contact *models.Contact) (id interface{}, err error)
}
