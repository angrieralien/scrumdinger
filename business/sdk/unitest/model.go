package unitest

import (
	"context"

	"github.com/angrieralien/scrumdinger/business/domain/homebus"
	"github.com/angrieralien/scrumdinger/business/domain/userbus"
)

// User represents an app user specified for the test.
type User struct {
	userbus.User
	Homes []homebus.Home
}

// SeedData represents data that was seeded for the test.
type SeedData struct {
	Users  []User
	Admins []User
}

// Table represent fields needed for running an unit test.
type Table struct {
	Name    string
	ExpResp any
	ExcFunc func(ctx context.Context) any
	CmpFunc func(got any, exp any) string
}
