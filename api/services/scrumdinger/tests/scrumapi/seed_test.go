package scrum_test

import (
	"context"
	"fmt"

	"github.com/angrieralien/scrumdinger/app/sdk/apitest"
	"github.com/angrieralien/scrumdinger/app/sdk/auth"
	"github.com/angrieralien/scrumdinger/business/domain/scrumbus"
	"github.com/angrieralien/scrumdinger/business/domain/userbus"
	"github.com/angrieralien/scrumdinger/business/sdk/dbtest"
	"github.com/angrieralien/scrumdinger/business/types/role"
)

func insertSeedData(db *dbtest.Database, ath *auth.Auth) (apitest.SeedData, error) {
	ctx := context.Background()
	busDomain := db.BusDomain

	usrs, err := userbus.TestSeedUsers(ctx, 1, role.User, busDomain.User)
	if err != nil {
		return apitest.SeedData{}, fmt.Errorf("seeding users : %w", err)
	}

	scrums, err := scrumbus.TestGenerateSeedScrums(ctx, 2, busDomain.Scrum, usrs[0].ID)
	if err != nil {
		return apitest.SeedData{}, fmt.Errorf("seeding scrums : %w", err)
	}

	tu1 := apitest.User{
		User:   usrs[0],
		Scrums: scrums,
		Token:  apitest.Token(db.BusDomain.User, ath, usrs[0].Email.Address),
	}

	// -------------------------------------------------------------------------

	usrs, err = userbus.TestSeedUsers(ctx, 1, role.User, busDomain.User)
	if err != nil {
		return apitest.SeedData{}, fmt.Errorf("seeding users : %w", err)
	}

	tu2 := apitest.User{
		User:  usrs[0],
		Token: apitest.Token(db.BusDomain.User, ath, usrs[0].Email.Address),
	}

	// -------------------------------------------------------------------------

	usrs, err = userbus.TestSeedUsers(ctx, 1, role.Admin, busDomain.User)
	if err != nil {
		return apitest.SeedData{}, fmt.Errorf("seeding users : %w", err)
	}

	scrums, err = scrumbus.TestGenerateSeedScrums(ctx, 2, busDomain.Scrum, usrs[0].ID)
	if err != nil {
		return apitest.SeedData{}, fmt.Errorf("seeding scrums : %w", err)
	}

	tu3 := apitest.User{
		User:   usrs[0],
		Scrums: scrums,
		Token:  apitest.Token(db.BusDomain.User, ath, usrs[0].Email.Address),
	}

	// -------------------------------------------------------------------------

	usrs, err = userbus.TestSeedUsers(ctx, 1, role.Admin, busDomain.User)
	if err != nil {
		return apitest.SeedData{}, fmt.Errorf("seeding users : %w", err)
	}

	tu4 := apitest.User{
		User:  usrs[0],
		Token: apitest.Token(db.BusDomain.User, ath, usrs[0].Email.Address),
	}

	// -------------------------------------------------------------------------

	sd := apitest.SeedData{
		Users:  []apitest.User{tu1, tu2},
		Admins: []apitest.User{tu3, tu4},
	}

	return sd, nil
}
