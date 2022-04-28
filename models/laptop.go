package models

import (
	"encoding/json"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"

	//"github.com/gofrs/uuid"
	"time"

	"github.com/gobuffalo/validate/v3/validators"
)

// Laptop is used by pop to map your laptops database table to your go code.
type Laptop struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`

	Os_name  string `json:"os_name" db:"os_name"`
	Asset_id int    `json:"asset_id" db:"asset_id"`

	//student Student `has_one:"laptop"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (l Laptop) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Laptops is not required by pop and may be deleted
type Laptops []Laptop

// String is not required by pop and may be deleted
func (l Laptops) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (l *Laptop) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: l.Name, Name: "Name"},
		&validators.StringIsPresent{Field: l.Email, Name: "Email"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (l *Laptop) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (l *Laptop) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
