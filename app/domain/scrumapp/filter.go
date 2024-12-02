package scrumapp

import (
	"net/http"
	"time"

	"github.com/angrieralien/scrumdinger/app/sdk/errs"
	"github.com/angrieralien/scrumdinger/business/domain/scrumbus"
	"github.com/google/uuid"
)

func parseQueryParams(r *http.Request) queryParams {
	values := r.URL.Query()

	filter := queryParams{
		Page:             values.Get("page"),
		Rows:             values.Get("row"),
		OrderBy:          values.Get("orderBy"),
		ID:               values.Get("scrum_id"),
		UserID:           values.Get("user_id"),
		StartCreatedDate: values.Get("start_created_date"),
		EndCreatedDate:   values.Get("end_created_date"),
	}

	return filter
}

func parseFilter(qp queryParams) (scrumbus.QueryFilter, error) {
	var filter scrumbus.QueryFilter

	if qp.ID != "" {
		id, err := uuid.Parse(qp.ID)
		if err != nil {
			return scrumbus.QueryFilter{}, errs.NewFieldsError("scrum_id", err)
		}
		filter.ID = &id
	}

	if qp.UserID != "" {
		id, err := uuid.Parse(qp.UserID)
		if err != nil {
			return scrumbus.QueryFilter{}, errs.NewFieldsError("user_id", err)
		}
		filter.UserID = &id
	}

	if qp.StartCreatedDate != "" {
		t, err := time.Parse(time.RFC3339, qp.StartCreatedDate)
		if err != nil {
			return scrumbus.QueryFilter{}, errs.NewFieldsError("start_created_date", err)
		}
		filter.StartCreatedDate = &t
	}

	if qp.EndCreatedDate != "" {
		t, err := time.Parse(time.RFC3339, qp.EndCreatedDate)
		if err != nil {
			return scrumbus.QueryFilter{}, errs.NewFieldsError("end_created_date", err)
		}
		filter.EndCreatedDate = &t
	}

	return filter, nil
}
