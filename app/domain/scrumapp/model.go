package scrumapp

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/angrieralien/scrumdinger/app/sdk/errs"
	"github.com/angrieralien/scrumdinger/app/sdk/mid"
	"github.com/angrieralien/scrumdinger/business/domain/scrumbus"
)

type queryParams struct {
	Page             string
	Rows             string
	OrderBy          string
	ID               string
	UserID           string
	StartCreatedDate string
	EndCreatedDate   string
}

// =============================================================================

// Scrum represents information about an individual scrum.
type Scrum struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Time        int      `json:"time"`
	Color       string   `json:"color"`
	Attendees   []string `json:"attendees"`
	UserID      string   `json:"userID"`
	DateCreated string   `json:"dateCreated"`
	DateUpdated string   `json:"dateUpdated"`
}

// Encode implements the encoder interface.
func (app Scrum) Encode() ([]byte, string, error) {
	data, err := json.Marshal(app)
	return data, "application/json", err
}

func toAppScrum(scrum scrumbus.Scrum) Scrum {
	return Scrum{
		ID:          scrum.ID.String(),
		UserID:      scrum.UserID.String(),
		Name:        scrum.Name,
		Time:        scrum.Time,
		Color:       scrum.Color,
		Attendees:   scrum.Attendees,
		DateCreated: scrum.DateCreated.Format(time.RFC3339),
		DateUpdated: scrum.DateUpdated.Format(time.RFC3339),
	}
}

func toAppScrums(scrums []scrumbus.Scrum) []Scrum {
	app := make([]Scrum, len(scrums))
	for i, scrum := range scrums {
		app[i] = toAppScrum(scrum)
	}

	return app
}

// =============================================================================

// NewScrum defines the data needed to add a new scrum.
type NewScrum struct {
	Name      string   `json:"name" validate:"required"`
	Time      int      `json:"time" validate:"required"`
	Color     string   `json:"color" validate:"required"`
	Attendees []string `json:"attendees" validate:"required"`
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

	bus := scrumbus.NewScrum{
		UserID:    userID,
		Name:      app.Name,
		Time:      app.Time,
		Color:     app.Color,
		Attendees: app.Attendees,
	}

	return bus, nil
}

// =============================================================================

// UpdateScrum defines the data needed to update a scrum.
type UpdateScrum struct {
	Name      *string  `json:"name"`
	Time      *int     `json:"time"`
	Color     *string  `json:"color"`
	Attendees []string `json:"attendees"`
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
	bus := scrumbus.UpdateScrum{
		Name:      app.Name,
		Time:      app.Time,
		Color:     app.Color,
		Attendees: app.Attendees,
	}

	return bus, nil
}
