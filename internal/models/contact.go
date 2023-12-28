package models

import "time"

type Contact struct {
	ID         int
	Subject    string
	Name       string
	Email      string
	Phone      string
	Address    string
	SetMode    string
	Created_At time.Time
}

func (c *Contact) NewContact() *Contact {
	return &Contact{}
}
