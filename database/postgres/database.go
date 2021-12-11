package postgres

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	MAX_RETRIES  = 10
	MAX_TIMEWAIT = 3
)

type Database struct{}

func (d *Database) GetConnection(dns string) (interface{}, error) {
	retries := 0
retry:
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		if retries < MAX_RETRIES {
			fmt.Println("Retrying to connect to the postgress database-->", retries+1)
			retries++
			time.Sleep(time.Second * time.Duration(MAX_TIMEWAIT))
			goto retry
		} else {
			return nil, err
		}
	}
	return db, nil // Here it is *gorm.DB
}
