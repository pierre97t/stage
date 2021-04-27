package model

import (
	"encoding/json"
	"fmt"
	"stagemain/entity"
)

// address is the struct used to manage address data
type address struct {
	Lines *map[uint]string `json:"lines,omitempty"`
	ZIP   *string          `json:"zip,omitempty"`
	City  *string          `json:"city,omitempty"`
	Alias *string          `json:"alias,omitempty"`
	Type  *string          `json:"type,omitempty"` // contact, delivery, invoice, ...
}

// Contact struct is used to manage contact data
type Contact struct {
	Std
	FirstName *string           `json:"firstName,omitempty"`
	LastName  *string           `json:"lastName,omitempty"`
	Email     *string           `json:"email,omitempty"`
	Phone     *string           `json:"phone,omitempty"`
	Addresses *map[uint]address `json:"addresses,omitempty"` // Contact addresses
}

// NewContact loads json byte data into a Contact struct
func NewContact(data []byte) (*Contact, error) {
	c := &Contact{}
	return c, handleJSONErr(TypeContact, data, json.Unmarshal(data, c))
}

// NewContacts loads json byte data into a *[]User struct
func NewContacts(data []byte) (*[]*Contact, error) {
	c := []*Contact{}
	return &c, handleJSONErr(TypeContact, data, json.Unmarshal(data, &c))
}

// New returns an instance of Contact with type settled
func (c *Contact) New() entity.Entity {
	tmp := &Contact{}
	t := TypeContact
	tmp.Type = &t
	return tmp
}

// Init initializes Contact's default parameter
func (c *Contact) Init(sessionID *string) error {
	if c.Type == nil {
		t := TypeContact
		c.Type = &t
	}
	return nil
}

// Validate will check that Contact data fit validation criteria
func (c *Contact) Validate(sessionID *string, checkRelConsist func(map[string]string) error) error {
	return nil
}

// Interface returns an interface, which will be a copy of the Contact whithout template fields
func (c *Contact) Interface() interface{} {
	return *c
}

// String returns object as string
func (c *Contact) String() string {
	return fmt.Sprintf("Contact entity content: %v", *c)
}

// List Returns an array of USer has interface
func (c *Contact) List() interface{} {
	return []*Contact{}
}
