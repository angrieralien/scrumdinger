package scrum_test

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/angrieralien/scrumdinger/app/domain/scrumapp"
	"github.com/angrieralien/scrumdinger/app/sdk/apitest"
	"github.com/angrieralien/scrumdinger/app/sdk/query"
	"github.com/angrieralien/scrumdinger/business/domain/scrumbus"
	"github.com/google/go-cmp/cmp"
)

func query200(sd apitest.SeedData) []apitest.Table {
	scrums := make([]scrumbus.Scrum, 0, len(sd.Admins[0].Scrums)+len(sd.Users[0].Scrums))
	scrums = append(scrums, sd.Admins[0].Scrums...)
	scrums = append(scrums, sd.Users[0].Scrums...)

	sort.Slice(scrums, func(i, j int) bool {
		return scrums[i].ID.String() <= scrums[j].ID.String()
	})

	table := []apitest.Table{
		{
			Name:       "basic",
			URL:        "/v1/scrums?page=1&rows=10&orderBy=scrum_id,ASC",
			Token:      sd.Admins[0].Token,
			StatusCode: http.StatusOK,
			Method:     http.MethodGet,
			GotResp:    &query.Result[scrumapp.Scrum]{},
			ExpResp: &query.Result[scrumapp.Scrum]{
				Page:        1,
				RowsPerPage: 10,
				Total:       len(scrums),
				Items:       toAppScrums(scrums),
			},
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
	}

	return table
}

func queryByID200(sd apitest.SeedData) []apitest.Table {
	table := []apitest.Table{
		{
			Name:       "basic",
			URL:        fmt.Sprintf("/v1/scrums/%s", sd.Users[0].Scrums[0].ID),
			Token:      sd.Users[0].Token,
			StatusCode: http.StatusOK,
			Method:     http.MethodGet,
			GotResp:    &scrumapp.Scrum{},
			ExpResp:    toAppScrumPtr(sd.Users[0].Scrums[0]),
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
	}

	return table
}
