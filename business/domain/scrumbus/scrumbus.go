// Package scrumbus provides business access to scrum domain.
package scrumbus

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/angrieralien/scrumdinger/business/domain/userbus"
	"github.com/angrieralien/scrumdinger/business/sdk/delegate"
	"github.com/angrieralien/scrumdinger/business/sdk/order"
	"github.com/angrieralien/scrumdinger/business/sdk/page"
	"github.com/angrieralien/scrumdinger/business/sdk/sqldb"
	"github.com/angrieralien/scrumdinger/foundation/logger"
	"github.com/angrieralien/scrumdinger/foundation/otel"
	"github.com/google/uuid"
)

// Set of error variables for CRUD operations.
var (
	ErrNotFound     = errors.New("scrum not found")
	ErrUserDisabled = errors.New("user disabled")
)

// Storer interface declares the behaviour this package needs to persist and
// retrieve data.
type Storer interface {
	NewWithTx(tx sqldb.CommitRollbacker) (Storer, error)
	Create(ctx context.Context, scrum Scrum) error
	Update(ctx context.Context, scrum Scrum) error
	Delete(ctx context.Context, scrum Scrum) error
	Query(ctx context.Context, filter QueryFilter, orderBy order.By, page page.Page) ([]Scrum, error)
	Count(ctx context.Context, filter QueryFilter) (int, error)
	QueryByID(ctx context.Context, scrumID uuid.UUID) (Scrum, error)
	QueryByUserID(ctx context.Context, userID uuid.UUID) ([]Scrum, error)
}

// Business manages the set of APIs for scrum api access.
type Business struct {
	log      *logger.Logger
	userBus  *userbus.Business
	delegate *delegate.Delegate
	storer   Storer
}

// NewBusiness constructs a scrum business API for use.
func NewBusiness(log *logger.Logger, userBus *userbus.Business, delegate *delegate.Delegate, storer Storer) *Business {
	return &Business{
		log:      log,
		userBus:  userBus,
		delegate: delegate,
		storer:   storer,
	}
}

// NewWithTx constructs a new domain value that will use the
// specified transaction in any store related calls.
func (b *Business) NewWithTx(tx sqldb.CommitRollbacker) (*Business, error) {
	storer, err := b.storer.NewWithTx(tx)
	if err != nil {
		return nil, err
	}

	userBus, err := b.userBus.NewWithTx(tx)
	if err != nil {
		return nil, err
	}

	bus := Business{
		log:      b.log,
		userBus:  userBus,
		delegate: b.delegate,
		storer:   storer,
	}

	return &bus, nil
}

// Create adds a new scrum to the system.
func (b *Business) Create(ctx context.Context, ns NewScrum) (Scrum, error) {
	ctx, span := otel.AddSpan(ctx, "business.scrumbus.create")
	defer span.End()

	usr, err := b.userBus.QueryByID(ctx, ns.UserID)
	if err != nil {
		return Scrum{}, fmt.Errorf("user.querybyid: %s: %w", ns.UserID, err)
	}

	if !usr.Enabled {
		return Scrum{}, ErrUserDisabled
	}

	now := time.Now()

	s := Scrum{
		ID:        uuid.New(),
		Name:      ns.Name,
		Time:      ns.Time,
		Color:     ns.Color,
		Attendees: ns.Attendees,

		UserID:      ns.UserID,
		DateCreated: now,
		DateUpdated: now,
	}

	if err := b.storer.Create(ctx, s); err != nil {
		return Scrum{}, fmt.Errorf("create: %w", err)
	}

	return s, nil
}

// Update modifies information about a scrum.
func (b *Business) Update(ctx context.Context, scrum Scrum, uh UpdateScrum) (Scrum, error) {
	ctx, span := otel.AddSpan(ctx, "business.scrumbus.update")
	defer span.End()

	scrum.DateUpdated = time.Now()

	if uh.Name != nil {
		scrum.Name = *uh.Name
	}

	if uh.Time != nil {
		scrum.Time = *uh.Time
	}

	if uh.Color != nil {
		scrum.Color = *uh.Color
	}

	if uh.Attendees != nil {
		scrum.Attendees = uh.Attendees
	}

	if err := b.storer.Update(ctx, scrum); err != nil {
		return Scrum{}, fmt.Errorf("update: %w", err)
	}

	return scrum, nil
}

// Delete removes the specified scrum.
func (b *Business) Delete(ctx context.Context, scrum Scrum) error {
	ctx, span := otel.AddSpan(ctx, "business.scrumbus.delete")
	defer span.End()

	if err := b.storer.Delete(ctx, scrum); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

// Query retrieves a list of existing scrums.
func (b *Business) Query(ctx context.Context, filter QueryFilter, orderBy order.By, page page.Page) ([]Scrum, error) {
	ctx, span := otel.AddSpan(ctx, "business.scrumbus.query")
	defer span.End()

	scrums, err := b.storer.Query(ctx, filter, orderBy, page)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return scrums, nil
}

// Count returns the total number of scrums.
func (b *Business) Count(ctx context.Context, filter QueryFilter) (int, error) {
	ctx, span := otel.AddSpan(ctx, "business.scrumbus.count")
	defer span.End()

	return b.storer.Count(ctx, filter)
}

// QueryByID finds the scrum by the specified ID.
func (b *Business) QueryByID(ctx context.Context, scrumID uuid.UUID) (Scrum, error) {
	ctx, span := otel.AddSpan(ctx, "business.scrumbus.querybyid")
	defer span.End()

	scrum, err := b.storer.QueryByID(ctx, scrumID)
	if err != nil {
		return Scrum{}, fmt.Errorf("query: scrumID[%s]: %w", scrumID, err)
	}

	return scrum, nil
}

// QueryByUserID finds the scrums by a specified User ID.
func (b *Business) QueryByUserID(ctx context.Context, userID uuid.UUID) ([]Scrum, error) {
	ctx, span := otel.AddSpan(ctx, "business.scrumbus.querybyuserid")
	defer span.End()

	scrums, err := b.storer.QueryByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return scrums, nil
}
