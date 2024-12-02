// Package scrumdb contains scrum related CRUD functionality.
package scrumdb

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/angrieralien/scrumdinger/business/domain/scrumbus"
	"github.com/angrieralien/scrumdinger/business/sdk/order"
	"github.com/angrieralien/scrumdinger/business/sdk/page"
	"github.com/angrieralien/scrumdinger/business/sdk/sqldb"
	"github.com/angrieralien/scrumdinger/foundation/logger"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Store manages the set of APIs for scrum database access.
type Store struct {
	log *logger.Logger
	db  sqlx.ExtContext
}

// NewStore constructs the api for data access.
func NewStore(log *logger.Logger, db *sqlx.DB) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}

// NewWithTx constructs a new Store value replacing the sqlx DB
// value with a sqlx DB value that is currently inside a transaction.
func (s *Store) NewWithTx(tx sqldb.CommitRollbacker) (scrumbus.Storer, error) {
	ec, err := sqldb.GetExtContext(tx)
	if err != nil {
		return nil, err
	}

	store := Store{
		log: s.log,
		db:  ec,
	}

	return &store, nil
}

// Create inserts a new scrum into the database.
func (s *Store) Create(ctx context.Context, hme scrumbus.Scrum) error {
	const q = `
    INSERT INTO scrums
        (scrum_id, user_id, name, time, color, attendees, type, address_1, address_2, zip_code, city, state, country, date_created, date_updated)
    VALUES
        (:scrum_id, :user_id, :name, :time, :color, :attendees, :type, :address_1, :address_2, :zip_code, :city, :state, :country, :date_created, :date_updated)`

	if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, toDBScrum(hme)); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Delete removes a scrum from the database.
func (s *Store) Delete(ctx context.Context, hme scrumbus.Scrum) error {
	data := struct {
		ID string `db:"scrum_id"`
	}{
		ID: hme.ID.String(),
	}

	const q = `
    DELETE FROM
	    scrums
	WHERE
	  	scrum_id = :scrum_id`

	if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Update replaces a scrum document in the database.
func (s *Store) Update(ctx context.Context, hme scrumbus.Scrum) error {
	const q = `
    UPDATE
        scrums
    SET
        "address_1"     = :address_1,
        "address_2"     = :address_2,
        "zip_code"      = :zip_code,
        "city"          = :city,
        "state"         = :state,
        "country"       = :country,
        "type"          = :type,
        "date_updated"  = :date_updated
    WHERE
        scrum_id = :scrum_id`

	if err := sqldb.NamedExecContext(ctx, s.log, s.db, q, toDBScrum(hme)); err != nil {
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Query retrieves a list of existing scrums from the database.
func (s *Store) Query(ctx context.Context, filter scrumbus.QueryFilter, orderBy order.By, page page.Page) ([]scrumbus.Scrum, error) {
	data := map[string]any{
		"offset":        (page.Number() - 1) * page.RowsPerPage(),
		"rows_per_page": page.RowsPerPage(),
	}

	const q = `
    SELECT
	    scrum_id, user_id, name, time, color, attendees, type, address_1, address_2, zip_code, city, state, country, date_created, date_updated
	FROM
	  	scrums`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	orderByClause, err := orderByClause(orderBy)
	if err != nil {
		return nil, err
	}

	buf.WriteString(orderByClause)
	buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	var dbHmes []scrum
	if err := sqldb.NamedQuerySlice(ctx, s.log, s.db, buf.String(), data, &dbHmes); err != nil {
		return nil, fmt.Errorf("namedqueryslice: %w", err)
	}

	hmes, err := toBusScrums(dbHmes)
	if err != nil {
		return nil, err
	}

	return hmes, nil
}

// Count returns the total number of scrums in the DB.
func (s *Store) Count(ctx context.Context, filter scrumbus.QueryFilter) (int, error) {
	data := map[string]any{}

	const q = `
    SELECT
        count(1)
    FROM
        scrums`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	var count struct {
		Count int `db:"count"`
	}
	if err := sqldb.NamedQueryStruct(ctx, s.log, s.db, buf.String(), data, &count); err != nil {
		return 0, fmt.Errorf("db: %w", err)
	}

	return count.Count, nil
}

// QueryByID gets the specified scrum from the database.
func (s *Store) QueryByID(ctx context.Context, scrumID uuid.UUID) (scrumbus.Scrum, error) {
	data := struct {
		ID string `db:"scrum_id"`
	}{
		ID: scrumID.String(),
	}

	const q = `
    SELECT
	  	scrum_id, user_id, name, time, color, attendees, type, address_1, address_2, zip_code, city, state, country, date_created, date_updated
    FROM
        scrums
    WHERE
        scrum_id = :scrum_id`

	var dbHme scrum
	if err := sqldb.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbHme); err != nil {
		if errors.Is(err, sqldb.ErrDBNotFound) {
			return scrumbus.Scrum{}, fmt.Errorf("db: %w", scrumbus.ErrNotFound)
		}
		return scrumbus.Scrum{}, fmt.Errorf("db: %w", err)
	}

	return toBusScrum(dbHme)
}

// QueryByUserID gets the specified scrum from the database by user id.
func (s *Store) QueryByUserID(ctx context.Context, userID uuid.UUID) ([]scrumbus.Scrum, error) {
	data := struct {
		ID string `db:"user_id"`
	}{
		ID: userID.String(),
	}

	const q = `
	SELECT
	    scrum_id, user_id, name, time, color, attendees, type, address_1, address_2, zip_code, city, state, country, date_created, date_updated
	FROM
		scrums
	WHERE
		user_id = :user_id`

	var dbHmes []scrum
	if err := sqldb.NamedQuerySlice(ctx, s.log, s.db, q, data, &dbHmes); err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}

	return toBusScrums(dbHmes)
}
