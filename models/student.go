package models

import (
	"encoding/json"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"

	"time"
)

// Student is used by pop to map your students database table to your go code.
type Student struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

// String is not required by pop and may be deleted
func (s Student) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Students is not required by pop and may be deleted
type Students []Student

// String is not required by pop and may be deleted
func (s Students) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Student) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Student) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Student) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
