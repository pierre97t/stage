package model

import (
	"encoding/json"
	"fmt"

	"stagemain/entity"
)

// User is the struc used to manage application's users
type User struct {
	Std
	Email     *string `json:"email,omitempty" validate:"alpha"`
	Username  *string `json:"username,omitempty" validate:"alpha"`
	FirstName *string `json:"firstName,omitempty" validate:"alpha"`
	LastName  *string `json:"lastName,omitempty" validate:"alpha"`
}

// NewUser loads json byte data into a User struct
func NewUser(data []byte) (*User, error) {
	u := &User{}
	return u, handleJSONErr(TypeUser, data, json.Unmarshal(data, u))
}

// NewUsers loads json byte data into a *[]User struct
func NewUsers(data []byte) (*[]*User, error) {
	u := []*User{}
	return &u, handleJSONErr(TypeUser, data, json.Unmarshal(data, &u))
}

// New returns an instance of user with type settled
func (u *User) New() entity.Entity {
	tmp := &User{}
	t := TypeUser
	tmp.Type = &t
	return tmp
}

// Init initializes user's default parameter
func (u *User) Init(sessionID *string) error {
	if u.Type == nil {
		t := TypeUser
		u.Type = &t
	}
	return nil
}

// Validate will check that user data fit validation criteria
func (u *User) Validate(sessionID *string, checkRelConsist func(map[string]string) error) error {
	return nil
}

// Interface returns an interface, which will be a copy of the User whithout template fields
func (u *User) Interface() interface{} {
	return *u
}

// String returns object as string
func (u *User) String() string {
	return fmt.Sprintf("User entity content: %v", *u)
}

// List Returns an array of USer has interface
func (u *User) List() interface{} {
	return []*User{}
}
