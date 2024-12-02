package scrumbus

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/angrieralien/scrumdinger/business/types/scrumtype"
	"github.com/google/uuid"
)

// TestGenerateNewScrums is a helper method for testing.
func TestGenerateNewScrums(n int, userID uuid.UUID) []NewScrum {
	newScrums := make([]NewScrum, n)

	idx := rand.Intn(10000)
	for i := 0; i < n; i++ {
		idx++

		t := scrumtype.Single
		if v := (idx + i) % 2; v == 0 {
			t = scrumtype.Condo
		}

		nh := NewScrum{
			Type:      t,
			Name:      fmt.Sprintf("Name%d", idx),
			Time:      idx,
			Color:     fmt.Sprintf("Color%d", idx),
			Attendees: []string{fmt.Sprintf("Attendee%d", idx), fmt.Sprintf("Attendee%d%d", idx, idx)},
			Address: Address{
				Address1: fmt.Sprintf("Address%d", idx),
				Address2: fmt.Sprintf("Address%d", idx),
				ZipCode:  fmt.Sprintf("%05d", idx),
				City:     fmt.Sprintf("City%d", idx),
				State:    fmt.Sprintf("State%d", idx),
				Country:  fmt.Sprintf("Country%d", idx),
			},
			UserID: userID,
		}

		newScrums[i] = nh
	}

	return newScrums
}

// TestGenerateSeedScrums is a helper method for testing.
func TestGenerateSeedScrums(ctx context.Context, n int, api *Business, userID uuid.UUID) ([]Scrum, error) {
	newScrums := TestGenerateNewScrums(n, userID)

	scrums := make([]Scrum, len(newScrums))
	for i, nh := range newScrums {
		scrum, err := api.Create(ctx, nh)
		if err != nil {
			return nil, fmt.Errorf("seeding scrum: idx: %d : %w", i, err)
		}

		scrums[i] = scrum
	}

	return scrums, nil
}

// ParseAddress is a helper function to create an address value.
func ParseAddress(address1 string, address2 string, zipCode string, city string, state string, country string) Address {
	return Address{
		Address1: address1,
		Address2: address2,
		ZipCode:  zipCode,
		City:     city,
		State:    state,
		Country:  country,
	}
}
