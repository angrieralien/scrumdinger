package scrumbus

import (
	"time"

	"github.com/angrieralien/scrumdinger/business/types/scrumtype"
	"github.com/google/uuid"
)

// Address represents an individual address.
type Address struct {
	Address1 string // We should create types for these fields.
	Address2 string
	ZipCode  string
	City     string
	State    string
	Country  string
}

// Scrum represents an individual scrum.
type Scrum struct {
	ID        uuid.UUID
	Name      string
	Time      int
	Color     string
	Attendees []string

	UserID      uuid.UUID
	Type        scrumtype.ScrumType
	Address     Address
	DateCreated time.Time
	DateUpdated time.Time
}

// NewScrum is what we require from clients when adding a Scrum.
type NewScrum struct {
	UserID uuid.UUID

	Name      string
	Time      int
	Color     string
	Attendees []string

	Type    scrumtype.ScrumType
	Address Address
}

// UpdateAddress is what fields can be updated in the store.
type UpdateAddress struct {
	Address1 *string
	Address2 *string
	ZipCode  *string
	City     *string
	State    *string
	Country  *string
}

// UpdateScrum defines what information may be provided to modify an existing
// Scrum. All fields are optional so clients can send only the fields they want
// changed. It uses pointer fields so we can differentiate between a field that
// was not provided and a field that was provided as explicitly blank. Normally
// we do not want to use pointers to basic types but we make exception around
// marshalling/unmarshalling.
type UpdateScrum struct {
	Name      *string
	Time      *int
	Color     *string
	Attendees []string

	Type    *scrumtype.ScrumType
	Address *UpdateAddress
}
