package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/pop/v5/slices"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/v3/validators"
)
// Thing is used by pop to map your things database table to your go code.
type Thing struct {
    ID uuid.UUID `json:"id" db:"id"`
    Name string `json:"name" db:"name"`
    Notes string `json:"notes" db:"notes"`
    Secret string `json:"secret" db:"secret"`
    Status string `json:"status" db:"status"`
    Tags slices.String `json:"tags" db:"tags"`
    UserID uuid.UUID `json:"user_id" db:"user_id"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (t Thing) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Things is not required by pop and may be deleted
type Things []Thing

// String is not required by pop and may be deleted
func (t Things) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Thing) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.Name, Name: "Name"},
		&validators.StringIsPresent{Field: t.Notes, Name: "Notes"},
		&validators.StringIsPresent{Field: t.Secret, Name: "Secret"},
		&validators.StringIsPresent{Field: t.Status, Name: "Status"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Thing) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Thing) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
