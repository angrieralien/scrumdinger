package scrumapp

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/angrieralien/scrumdinger/app/sdk/errs"
	"github.com/angrieralien/scrumdinger/app/sdk/mid"
	"github.com/angrieralien/scrumdinger/business/domain/scrumbus"
	"github.com/angrieralien/scrumdinger/business/types/scrumtype"
)

type queryParams struct {
	Page             string
	Rows             string
	OrderBy          string
	ID               string
	UserID           string
	Type             string
	StartCreatedDate string
	EndCreatedDate   string
}

// =============================================================================

// Address represents information about an individual address.
type Address struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	ZipCode  string `json:"zipCode"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
}

// Scrum represents information about an individual scrum.
type Scrum struct {
	ID          string  `json:"id"`
	UserID      string  `json:"userID"`
	Type        string  `json:"type"`
	Address     Address `json:"address"`
	DateCreated string  `json:"dateCreated"`
	DateUpdated string  `json:"dateUpdated"`
}

// Encode implements the encoder interface.
func (app Scrum) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}

func toAppScrum(hme scrumbus.Scrum) Scrum {
	return Scrum{
		ID:     hme.ID.String(),
		UserID: hme.UserID.String(),
		Type:   hme.Type.String(),
		Address: Address{
			Address1: hme.Address.Address1,
			Address2: hme.Address.Address2,
			ZipCode:  hme.Address.ZipCode,
			City:     hme.Address.City,
			State:    hme.Address.State,
			Country:  hme.Address.Country,
		},
		DateCreated: hme.DateCreated.Format(time.RFC3339),
		DateUpdated: hme.DateUpdated.Format(time.RFC3339),
	}
}

func toAppScrums(scrums []scrumbus.Scrum) []Scrum {
	app := make([]Scrum, len(scrums))
	for i, hme := range scrums {
		app[i] = toAppScrum(hme)
	}

	return app
}

// =============================================================================

// NewAddress defines the data needed to add a new address.
type NewAddress struct {
	Address1 string `json:"address1" validate:"required,min=1,max=70"`
	Address2 string `json:"address2" validate:"omitempty,max=70"`
	ZipCode  string `json:"zipCode" validate:"required,numeric"`
	City     string `json:"city" validate:"required"`
	State    string `json:"state" validate:"required,min=1,max=48"`
	Country  string `json:"country" validate:"required,iso3166_1_alpha2"`
}

// NewScrum defines the data needed to add a new scrum.
type NewScrum struct {
	Name      string     `json:"type" validate:"required"`
	Time      int        `json:"time" validate:"required"`
	Color     string     `json:"color" validate:"required"`
	Attendees []string   `json:"attendees" validate:"required"`
	Address   NewAddress `json:"address"`
}

// Decode implements the decoder interface.
func (app *NewScrum) Decode(data []byte) error {
	return json.Unmarshal(data, app)
}

// Validate checks if the data in the model is considered clean.
func (app NewScrum) Validate() error {
	if err := errs.Check(app); err != nil {
		return errs.Newf(errs.InvalidArgument, "validate: %s", err)
	}

	return nil
}

func toBusNewScrum(ctx context.Context, app NewScrum) (scrumbus.NewScrum, error) {
	userID, err := mid.GetUserID(ctx)
	if err != nil {
		return scrumbus.NewScrum{}, fmt.Errorf("getuserid: %w", err)
	}

	typ, err := scrumtype.Parse(app.Name)
	if err != nil {
		return scrumbus.NewScrum{}, fmt.Errorf("parse: %w", err)
	}

	bus := scrumbus.NewScrum{
		UserID: userID,
		Type:   typ,
		Address: scrumbus.Address{
			Address1: app.Address.Address1,
			Address2: app.Address.Address2,
			ZipCode:  app.Address.ZipCode,
			City:     app.Address.City,
			State:    app.Address.State,
			Country:  app.Address.Country,
		},
	}

	return bus, nil
}

// =============================================================================

// UpdateAddress defines the data needed to update an address.
type UpdateAddress struct {
	Address1 *string `json:"address1" validate:"omitempty,min=1,max=70"`
	Address2 *string `json:"address2" validate:"omitempty,max=70"`
	ZipCode  *string `json:"zipCode" validate:"omitempty,numeric"`
	City     *string `json:"city"`
	State    *string `json:"state" validate:"omitempty,min=1,max=48"`
	Country  *string `json:"country" validate:"omitempty,iso3166_1_alpha2"`
}

// UpdateScrum defines the data needed to update a scrum.
type UpdateScrum struct {
	Type    *string        `json:"type"`
	Address *UpdateAddress `json:"address"`
}

// Decode implements the decoder interface.
func (app *UpdateScrum) Decode(data []byte) error {
	return json.Unmarshal(data, app)
}

// Validate checks the data in the model is considered clean.
func (app UpdateScrum) Validate() error {
	if err := errs.Check(app); err != nil {
		return errs.Newf(errs.InvalidArgument, "validate: %s", err)
	}

	return nil
}

func toBusUpdateScrum(app UpdateScrum) (scrumbus.UpdateScrum, error) {
	var t scrumtype.ScrumType
	if app.Type != nil {
		var err error
		t, err = scrumtype.Parse(*app.Type)
		if err != nil {
			return scrumbus.UpdateScrum{}, fmt.Errorf("parse: %w", err)
		}
	}

	bus := scrumbus.UpdateScrum{
		Type: &t,
	}

	if app.Address != nil {
		bus.Address = &scrumbus.UpdateAddress{
			Address1: app.Address.Address1,
			Address2: app.Address.Address2,
			ZipCode:  app.Address.ZipCode,
			City:     app.Address.City,
			State:    app.Address.State,
			Country:  app.Address.Country,
		}
	}

	return bus, nil
}
