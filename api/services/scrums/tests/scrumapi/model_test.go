package scrum_test

import (
	"time"

	"github.com/angrieralien/scrumdinger/app/domain/scrumapp"
	"github.com/angrieralien/scrumdinger/business/domain/scrumbus"
)

func toAppScrum(hme scrumbus.Scrum) scrumapp.Scrum {
	return scrumapp.Scrum{
		ID:     hme.ID.String(),
		UserID: hme.UserID.String(),
		Type:   hme.Type.String(),
		Address: scrumapp.Address{
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

func toAppScrums(scrums []scrumbus.Scrum) []scrumapp.Scrum {
	items := make([]scrumapp.Scrum, len(scrums))
	for i, hme := range scrums {
		items[i] = toAppScrum(hme)
	}

	return items
}

func toAppScrumPtr(hme scrumbus.Scrum) *scrumapp.Scrum {
	appHme := toAppScrum(hme)
	return &appHme
}
