package scrumbus

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

// TestGenerateNewScrums is a helper method for testing.
func TestGenerateNewScrums(n int, userID uuid.UUID) []NewScrum {
	newScrums := make([]NewScrum, n)

	idx := rand.Intn(10000)
	for i := 0; i < n; i++ {
		idx++

		nh := NewScrum{
			Name:      fmt.Sprintf("Name%d", idx),
			Time:      idx,
			Color:     fmt.Sprintf("Color%d", idx),
			Attendees: []string{fmt.Sprintf("Attendee%d", idx), fmt.Sprintf("Attendee%d%d", idx, idx)},
			UserID:    userID,
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
