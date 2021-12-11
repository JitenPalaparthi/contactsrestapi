package filedb

import (
	"contacts/models"
	"fmt"
	"os"
)

type ContactFileDB struct {
	DBClient interface{}
}

func (c *ContactFileDB) Create(contact *models.Contact) (id interface{}, err error) {
	var f *os.File
	f, err = os.Open("contacts.fdb")
	if err != nil {
		f, err = os.Create("contacts.fdb")
		if err != nil {
			return nil, err
		}
	} else {
		defer f.Close()
	}
	return f.WriteString(fmt.Sprint(contact))
}

func (c *ContactFileDB) Delete(id int) (records interface{}, err error) {
	if err := os.Remove("contacts.fdb"); err != nil {
		return nil, err
	}
	return 1, nil
}
