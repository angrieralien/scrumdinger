package home_test

import (
	"time"

	"github.com/angrieralien/scrumdinger/app/domain/scrumapp"
	"github.com/angrieralien/scrumdinger/business/domain/scrumbus"
)

func toAppHome(hme scrumbus.Home) scrumapp.Home {
	return scrumapp.Home{
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

func toAppHomes(homes []scrumbus.Home) []scrumapp.Home {
	items := make([]scrumapp.Home, len(homes))
	for i, hme := range homes {
		items[i] = toAppHome(hme)
	}

	return items
}

func toAppHomePtr(hme scrumbus.Home) *scrumapp.Home {
	appHme := toAppHome(hme)
	return &appHme
}
