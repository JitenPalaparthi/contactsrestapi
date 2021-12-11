package models

import "fmt"

type Contact struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	Status       string `json:"status"`
	LastModified string `json:"lastModified"`
}

func (c *Contact) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("invalid Name or Name is not provided")
	}
	if c.Email == "" {
		return fmt.Errorf("invalid Email or Email is not provided")
	}
	if c.Mobile == "" {
		return fmt.Errorf("invalid Mobile or Mobile is not provided")
	}
	return nil
}

func Add(i, j int) int {
	return i + j
}
