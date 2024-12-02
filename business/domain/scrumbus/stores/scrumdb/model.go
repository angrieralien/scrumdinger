package scrumdb

import (
	"fmt"
	"strings"
	"time"

	"github.com/angrieralien/scrumdinger/business/domain/scrumbus"
	"github.com/angrieralien/scrumdinger/business/types/scrumtype"
	"github.com/google/uuid"
)

type scrum struct {
	ID     uuid.UUID `db:"scrum_id"`
	UserID uuid.UUID `db:"user_id"`

	Name      string `db:"name"`
	Time      int    `db:"time"`
	Color     string `db:"color"`
	Attendees string `db:"attendees"`

	Type        string    `db:"type"`
	Address1    string    `db:"address_1"`
	Address2    string    `db:"address_2"`
	ZipCode     string    `db:"zip_code"`
	City        string    `db:"city"`
	Country     string    `db:"country"`
	State       string    `db:"state"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

func toDBScrum(bus scrumbus.Scrum) scrum {
	db := scrum{
		ID:     bus.ID,
		UserID: bus.UserID,

		Name:      bus.Name,
		Time:      bus.Time,
		Color:     bus.Color,
		Attendees: strings.Join(bus.Attendees, "\t"),

		Type:        bus.Type.String(),
		Address1:    bus.Address.Address1,
		Address2:    bus.Address.Address2,
		ZipCode:     bus.Address.ZipCode,
		City:        bus.Address.City,
		Country:     bus.Address.Country,
		State:       bus.Address.State,
		DateCreated: bus.DateCreated.UTC(),
		DateUpdated: bus.DateUpdated.UTC(),
	}

	return db
}

func toBusScrum(db scrum) (scrumbus.Scrum, error) {
	typ, err := scrumtype.Parse(db.Type)
	if err != nil {
		return scrumbus.Scrum{}, fmt.Errorf("parse type: %w", err)
	}

	bus := scrumbus.Scrum{
		ID:        db.ID,
		Name:      db.Name,
		Time:      db.Time,
		Color:     db.Color,
		Attendees: strings.Split(db.Attendees, "\t"),
		UserID:    db.UserID,
		Type:      typ,
		Address: scrumbus.Address{
			Address1: db.Address1,
			Address2: db.Address2,
			ZipCode:  db.ZipCode,
			City:     db.City,
			Country:  db.Country,
			State:    db.State,
		},
		DateCreated: db.DateCreated.In(time.Local),
		DateUpdated: db.DateUpdated.In(time.Local),
	}

	return bus, nil
}

func toBusScrums(dbs []scrum) ([]scrumbus.Scrum, error) {
	bus := make([]scrumbus.Scrum, len(dbs))

	for i, db := range dbs {
		var err error
		bus[i], err = toBusScrum(db)
		if err != nil {
			return nil, fmt.Errorf("parse type: %w", err)
		}
	}

	return bus, nil
}
