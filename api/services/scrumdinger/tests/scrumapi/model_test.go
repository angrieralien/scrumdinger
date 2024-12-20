package scrum_test

import (
	"time"

	"github.com/angrieralien/scrumdinger/app/domain/scrumapp"
	"github.com/angrieralien/scrumdinger/business/domain/scrumbus"
)

func toAppScrum(scrum scrumbus.Scrum) scrumapp.Scrum {
	return scrumapp.Scrum{
		ID:        scrum.ID.String(),
		UserID:    scrum.UserID.String(),
		Name:      scrum.Name,
		Time:      scrum.Time,
		Color:     scrum.Color,
		Attendees: scrum.Attendees,

		DateCreated: scrum.DateCreated.Format(time.RFC3339),
		DateUpdated: scrum.DateUpdated.Format(time.RFC3339),
	}
}

func toAppScrums(scrums []scrumbus.Scrum) []scrumapp.Scrum {
	items := make([]scrumapp.Scrum, len(scrums))
	for i, scrum := range scrums {
		items[i] = toAppScrum(scrum)
	}

	return items
}

func toAppScrumPtr(scrum scrumbus.Scrum) *scrumapp.Scrum {
	appHme := toAppScrum(scrum)
	return &appHme
}
