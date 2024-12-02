package scrumdb

import (
	"fmt"
	"strings"
	"time"

	"github.com/angrieralien/scrumdinger/business/domain/scrumbus"
	"github.com/google/uuid"
)

type scrum struct {
	ID     uuid.UUID `db:"scrum_id"`
	UserID uuid.UUID `db:"user_id"`

	Name      string `db:"name"`
	Time      int    `db:"time"`
	Color     string `db:"color"`
	Attendees string `db:"attendees"`

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

		DateCreated: bus.DateCreated.UTC(),
		DateUpdated: bus.DateUpdated.UTC(),
	}

	return db
}

func toBusScrum(db scrum) (scrumbus.Scrum, error) {
	attendees := []string{}
	if len(db.Attendees) > 0 {
		attendees = strings.Split(db.Attendees, "\t")
	}

	bus := scrumbus.Scrum{
		ID:        db.ID,
		Name:      db.Name,
		Time:      db.Time,
		Color:     db.Color,
		Attendees: attendees,
		UserID:    db.UserID,

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
