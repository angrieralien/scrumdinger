package scrum_test

import (
	"fmt"
	"net/http"
	"time"

	"github.com/angrieralien/scrumdinger/app/domain/scrumapp"
	"github.com/angrieralien/scrumdinger/app/sdk/apitest"
	"github.com/angrieralien/scrumdinger/app/sdk/errs"
	"github.com/angrieralien/scrumdinger/business/sdk/dbtest"
	"github.com/google/go-cmp/cmp"
)

func update200(sd apitest.SeedData) []apitest.Table {
	table := []apitest.Table{
		{
			Name:       "basic",
			URL:        fmt.Sprintf("/v1/scrums/%s", sd.Users[0].Scrums[0].ID),
			Token:      sd.Users[0].Token,
			Method:     http.MethodPut,
			StatusCode: http.StatusOK,
			Input: &scrumapp.UpdateScrum{
				Name:      dbtest.StringPointer("App 1"),
				Time:      dbtest.IntPointer(10),
				Color:     dbtest.StringPointer("navy"),
				Attendees: []string{"stephen", "luke"},
			},
			GotResp: &scrumapp.Scrum{},
			ExpResp: &scrumapp.Scrum{
				ID:          sd.Users[0].Scrums[0].ID.String(),
				UserID:      sd.Users[0].ID.String(),
				Name:        "App 1",
				Time:        10,
				Color:       "navy",
				Attendees:   []string{"stephen", "luke"},
				DateCreated: sd.Users[0].Scrums[0].DateCreated.Format(time.RFC3339),
				DateUpdated: sd.Users[0].Scrums[0].DateCreated.Format(time.RFC3339),
			},
			CmpFunc: func(got any, exp any) string {
				gotResp, exists := got.(*scrumapp.Scrum)
				if !exists {
					return "error occurred"
				}

				expResp := exp.(*scrumapp.Scrum)
				gotResp.DateUpdated = expResp.DateUpdated

				return cmp.Diff(gotResp, expResp)
			},
		},
	}

	return table
}

func update401(sd apitest.SeedData) []apitest.Table {
	table := []apitest.Table{
		{
			Name:       "emptytoken",
			URL:        fmt.Sprintf("/v1/scrums/%s", sd.Users[0].Scrums[0].ID),
			Token:      "&nbsp;",
			Method:     http.MethodPut,
			StatusCode: http.StatusUnauthorized,
			GotResp:    &errs.Error{},
			ExpResp:    errs.Newf(errs.Unauthenticated, "error parsing token: token contains an invalid number of segments"),
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
		{
			Name:       "badsig",
			URL:        fmt.Sprintf("/v1/scrums/%s", sd.Users[0].Scrums[0].ID),
			Token:      sd.Users[0].Token + "A",
			Method:     http.MethodPut,
			StatusCode: http.StatusUnauthorized,
			GotResp:    &errs.Error{},
			ExpResp:    errs.Newf(errs.Unauthenticated, "authentication failed : bindings results[[{[true] map[x:false]}]] ok[true]"),
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
		{
			Name:       "wronguser",
			URL:        fmt.Sprintf("/v1/scrums/%s", sd.Admins[0].Scrums[0].ID),
			Token:      sd.Users[0].Token,
			Method:     http.MethodPut,
			StatusCode: http.StatusUnauthorized,
			Input: &scrumapp.UpdateScrum{
				Name: dbtest.StringPointer("wrong user app"),
			},
			GotResp: &errs.Error{},
			ExpResp: errs.Newf(errs.Unauthenticated, "authorize: you are not authorized for that action, claims[[USER]] rule[rule_admin_or_subject]: rego evaluation failed : bindings results[[{[true] map[x:false]}]] ok[true]"),
			CmpFunc: func(got any, exp any) string {
				return cmp.Diff(got, exp)
			},
		},
	}

	return table
}
