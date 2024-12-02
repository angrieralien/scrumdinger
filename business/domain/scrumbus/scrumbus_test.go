package scrumbus_test

import (
	"context"
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/angrieralien/scrumdinger/business/domain/scrumbus"
	"github.com/angrieralien/scrumdinger/business/domain/userbus"
	"github.com/angrieralien/scrumdinger/business/sdk/dbtest"
	"github.com/angrieralien/scrumdinger/business/sdk/page"
	"github.com/angrieralien/scrumdinger/business/sdk/unitest"
	"github.com/angrieralien/scrumdinger/business/types/role"
	"github.com/google/go-cmp/cmp"
)

func Test_Scrum(t *testing.T) {
	t.Parallel()

	db := dbtest.New(t, "Test_Scrum")

	sd, err := insertSeedData(db.BusDomain)
	if err != nil {
		t.Fatalf("Seeding error: %s", err)
	}

	// -------------------------------------------------------------------------

	unitest.Run(t, query(db.BusDomain, sd), "query")
	unitest.Run(t, create(db.BusDomain, sd), "create")
	unitest.Run(t, update(db.BusDomain, sd), "update")
	unitest.Run(t, delete(db.BusDomain, sd), "delete")
}

// =============================================================================

func insertSeedData(busDomain dbtest.BusDomain) (unitest.SeedData, error) {
	ctx := context.Background()

	usrs, err := userbus.TestSeedUsers(ctx, 1, role.User, busDomain.User)
	if err != nil {
		return unitest.SeedData{}, fmt.Errorf("seeding users : %w", err)
	}

	scrums, err := scrumbus.TestGenerateSeedScrums(ctx, 2, busDomain.Scrum, usrs[0].ID)
	if err != nil {
		return unitest.SeedData{}, fmt.Errorf("seeding scrums : %w", err)
	}

	tu1 := unitest.User{
		User:   usrs[0],
		Scrums: scrums,
	}

	// -------------------------------------------------------------------------

	usrs, err = userbus.TestSeedUsers(ctx, 1, role.User, busDomain.User)
	if err != nil {
		return unitest.SeedData{}, fmt.Errorf("seeding users : %w", err)
	}

	tu2 := unitest.User{
		User: usrs[0],
	}

	// -------------------------------------------------------------------------

	usrs, err = userbus.TestSeedUsers(ctx, 1, role.Admin, busDomain.User)
	if err != nil {
		return unitest.SeedData{}, fmt.Errorf("seeding users : %w", err)
	}

	scrums, err = scrumbus.TestGenerateSeedScrums(ctx, 2, busDomain.Scrum, usrs[0].ID)
	if err != nil {
		return unitest.SeedData{}, fmt.Errorf("seeding scrums : %w", err)
	}

	tu3 := unitest.User{
		User:   usrs[0],
		Scrums: scrums,
	}

	// -------------------------------------------------------------------------

	usrs, err = userbus.TestSeedUsers(ctx, 1, role.Admin, busDomain.User)
	if err != nil {
		return unitest.SeedData{}, fmt.Errorf("seeding users : %w", err)
	}

	tu4 := unitest.User{
		User: usrs[0],
	}

	// -------------------------------------------------------------------------

	sd := unitest.SeedData{
		Users:  []unitest.User{tu1, tu2},
		Admins: []unitest.User{tu3, tu4},
	}

	return sd, nil
}

// =============================================================================

func query(busDomain dbtest.BusDomain, sd unitest.SeedData) []unitest.Table {
	scrums := make([]scrumbus.Scrum, 0, len(sd.Admins[0].Scrums)+len(sd.Users[0].Scrums))
	scrums = append(scrums, sd.Admins[0].Scrums...)
	scrums = append(scrums, sd.Users[0].Scrums...)

	sort.Slice(scrums, func(i, j int) bool {
		return scrums[i].ID.String() <= scrums[j].ID.String()
	})

	table := []unitest.Table{
		{
			Name:    "all",
			ExpResp: scrums,
			ExcFunc: func(ctx context.Context) any {
				resp, err := busDomain.Scrum.Query(ctx, scrumbus.QueryFilter{}, scrumbus.DefaultOrderBy, page.MustParse("1", "10"))
				if err != nil {
					return err
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				gotResp, exists := got.([]scrumbus.Scrum)
				if !exists {
					return "error occurred"
				}

				expResp := exp.([]scrumbus.Scrum)

				for i := range gotResp {
					if gotResp[i].DateCreated.Format(time.RFC3339) == expResp[i].DateCreated.Format(time.RFC3339) {
						expResp[i].DateCreated = gotResp[i].DateCreated
					}

					if gotResp[i].DateUpdated.Format(time.RFC3339) == expResp[i].DateUpdated.Format(time.RFC3339) {
						expResp[i].DateUpdated = gotResp[i].DateUpdated
					}
				}

				return cmp.Diff(gotResp, expResp)
			},
		},
		{
			Name:    "byid",
			ExpResp: sd.Users[0].Scrums[0],
			ExcFunc: func(ctx context.Context) any {
				resp, err := busDomain.Scrum.QueryByID(ctx, sd.Users[0].Scrums[0].ID)
				if err != nil {
					return err
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				gotResp, exists := got.(scrumbus.Scrum)
				if !exists {
					return "error occurred"
				}

				expResp := exp.(scrumbus.Scrum)

				if gotResp.DateCreated.Format(time.RFC3339) == expResp.DateCreated.Format(time.RFC3339) {
					expResp.DateCreated = gotResp.DateCreated
				}

				if gotResp.DateUpdated.Format(time.RFC3339) == expResp.DateUpdated.Format(time.RFC3339) {
					expResp.DateUpdated = gotResp.DateUpdated
				}

				return cmp.Diff(gotResp, expResp)
			},
		},
	}

	return table
}

func create(busDomain dbtest.BusDomain, sd unitest.SeedData) []unitest.Table {
	table := []unitest.Table{
		{
			Name: "basic",
			ExpResp: scrumbus.Scrum{
				UserID:    sd.Users[0].ID,
				Name:      "Scrumdinger",
				Time:      5,
				Color:     "blue",
				Attendees: []string{"stephen", "mark", "matthew"},
			},
			ExcFunc: func(ctx context.Context) any {
				nh := scrumbus.NewScrum{
					UserID:    sd.Users[0].ID,
					Name:      "Scrumdinger",
					Time:      5,
					Color:     "blue",
					Attendees: []string{"stephen", "mark", "matthew"},
				}

				resp, err := busDomain.Scrum.Create(ctx, nh)
				if err != nil {
					return err
				}

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				gotResp, exists := got.(scrumbus.Scrum)
				if !exists {
					return "error occurred"
				}

				expResp := exp.(scrumbus.Scrum)

				expResp.ID = gotResp.ID
				expResp.DateCreated = gotResp.DateCreated
				expResp.DateUpdated = gotResp.DateUpdated

				return cmp.Diff(gotResp, expResp)
			},
		},
	}

	return table
}

func update(busDomain dbtest.BusDomain, sd unitest.SeedData) []unitest.Table {
	table := []unitest.Table{
		{
			Name: "basic",
			ExpResp: scrumbus.Scrum{
				ID:          sd.Users[0].Scrums[0].ID,
				UserID:      sd.Users[0].ID,
				Name:        "update",
				Time:        sd.Users[0].Scrums[0].Time,
				Color:       sd.Users[0].Scrums[0].Color,
				Attendees:   sd.Users[0].Scrums[0].Attendees,
				DateCreated: sd.Users[0].Scrums[0].DateCreated,
				DateUpdated: sd.Users[0].Scrums[0].DateCreated,
			},
			ExcFunc: func(ctx context.Context) any {
				uh := scrumbus.UpdateScrum{
					Name: dbtest.StringPointer("update"),
				}

				resp, err := busDomain.Scrum.Update(ctx, sd.Users[0].Scrums[0], uh)
				if err != nil {
					return err
				}

				resp.DateUpdated = resp.DateCreated

				return resp
			},
			CmpFunc: func(got any, exp any) string {
				gotResp, exists := got.(scrumbus.Scrum)
				if !exists {
					return "error occurred"
				}

				expResp := exp.(scrumbus.Scrum)

				expResp.DateUpdated = gotResp.DateUpdated

				return cmp.Diff(gotResp, expResp)
			},
		},
	}

	return table
}

func delete(busDomain dbtest.BusDomain, sd unitest.SeedData) []unitest.Table {
	table := []unitest.Table{
		{
			Name:    "user",
			ExpResp: nil,
			ExcFunc: func(ctx context.Context) any {
				if err := busDomain.Scrum.Delete(ctx, sd.Users[0].Scrums[1]); err != nil {
					return err
				}

				return nil
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
		{
			Name:    "admin",
			ExpResp: nil,
			ExcFunc: func(ctx context.Context) any {
				if err := busDomain.Scrum.Delete(ctx, sd.Admins[0].Scrums[1]); err != nil {
					return err
				}

				return nil
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
	}

	return table
}
