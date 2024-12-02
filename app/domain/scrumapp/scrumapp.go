// Package scrumapp maintains the app layer api for the scrum domain.
package scrumapp

import (
	"context"
	"net/http"

	"github.com/angrieralien/scrumdinger/app/sdk/errs"
	"github.com/angrieralien/scrumdinger/app/sdk/mid"
	"github.com/angrieralien/scrumdinger/app/sdk/query"
	"github.com/angrieralien/scrumdinger/business/domain/scrumbus"
	"github.com/angrieralien/scrumdinger/business/sdk/order"
	"github.com/angrieralien/scrumdinger/business/sdk/page"
	"github.com/angrieralien/scrumdinger/foundation/web"
)

type app struct {
	scrumBus *scrumbus.Business
}

func newApp(scrumBus *scrumbus.Business) *app {
	return &app{
		scrumBus: scrumBus,
	}
}

func (a *app) create(ctx context.Context, r *http.Request) web.Encoder {
	var app NewScrum
	if err := web.Decode(r, &app); err != nil {
		return errs.New(errs.InvalidArgument, err)
	}

	nh, err := toBusNewScrum(ctx, app)
	if err != nil {
		return errs.New(errs.InvalidArgument, err)
	}

	scrum, err := a.scrumBus.Create(ctx, nh)
	if err != nil {
		return errs.Newf(errs.Internal, "create: scrum[%+v]: %s", app, err)
	}

	return toAppScrum(scrum)
}

func (a *app) update(ctx context.Context, r *http.Request) web.Encoder {
	var app UpdateScrum
	if err := web.Decode(r, &app); err != nil {
		return errs.New(errs.InvalidArgument, err)
	}

	uh, err := toBusUpdateScrum(app)
	if err != nil {
		return errs.New(errs.InvalidArgument, err)
	}

	scrum, err := mid.GetScrum(ctx)
	if err != nil {
		return errs.Newf(errs.Internal, "scrum missing in context: %s", err)
	}

	updUsr, err := a.scrumBus.Update(ctx, scrum, uh)
	if err != nil {
		return errs.Newf(errs.Internal, "update: scrumID[%s] uh[%+v]: %s", scrum.ID, uh, err)
	}

	return toAppScrum(updUsr)
}

func (a *app) delete(ctx context.Context, _ *http.Request) web.Encoder {
	scrum, err := mid.GetScrum(ctx)
	if err != nil {
		return errs.Newf(errs.Internal, "scrumID missing in context: %s", err)
	}

	if err := a.scrumBus.Delete(ctx, scrum); err != nil {
		return errs.Newf(errs.Internal, "delete: scrumID[%s]: %s", scrum.ID, err)
	}

	return nil
}

func (a *app) query(ctx context.Context, r *http.Request) web.Encoder {
	qp := parseQueryParams(r)

	page, err := page.Parse(qp.Page, qp.Rows)
	if err != nil {
		return errs.NewFieldsError("page", err)
	}

	filter, err := parseFilter(qp)
	if err != nil {
		return err.(errs.FieldErrors)
	}

	orderBy, err := order.Parse(orderByFields, qp.OrderBy, scrumbus.DefaultOrderBy)
	if err != nil {
		return errs.NewFieldsError("order", err)
	}

	scrums, err := a.scrumBus.Query(ctx, filter, orderBy, page)
	if err != nil {
		return errs.Newf(errs.Internal, "query: %s", err)
	}

	total, err := a.scrumBus.Count(ctx, filter)
	if err != nil {
		return errs.Newf(errs.Internal, "count: %s", err)
	}

	return query.NewResult(toAppScrums(scrums), total, page)
}

func (a *app) queryByID(ctx context.Context, _ *http.Request) web.Encoder {
	scrum, err := mid.GetScrum(ctx)
	if err != nil {
		return errs.Newf(errs.Internal, "querybyid: %s", err)
	}
	return toAppScrum(scrum)
}

func (a *app) queryByUserID(ctx context.Context, r *http.Request) web.Encoder {
	qp := parseQueryParams(r)

	page, err := page.Parse(qp.Page, qp.Rows)
	if err != nil {
		return errs.NewFieldsError("page", err)
	}

	filter, err := parseFilter(qp)
	if err != nil {
		return err.(errs.FieldErrors)
	}

	userID, err := mid.GetUserID(ctx)
	if err != nil {
		return errs.Newf(errs.Internal, "user missing in context: %s", err)
	}

	scrums, err := a.scrumBus.QueryByUserID(ctx, userID)

	if err != nil {
		return errs.Newf(errs.Internal, "query: %s", err)
	}

	total, err := a.scrumBus.Count(ctx, filter)
	if err != nil {
		return errs.Newf(errs.Internal, "count: %s", err)
	}

	return query.NewResult(toAppScrums(scrums), total, page)
}
