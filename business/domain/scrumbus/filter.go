package scrumbus

import (
	"time"

	"github.com/angrieralien/scrumdinger/business/types/scrumtype"
	"github.com/google/uuid"
)

// QueryFilter holds the available fields a query can be filtered on.
// We are using pointer semantics because the With API mutates the value.
type QueryFilter struct {
	ID               *uuid.UUID
	UserID           *uuid.UUID
	Name             *string
	Type             *scrumtype.ScrumType
	StartCreatedDate *time.Time
	EndCreatedDate   *time.Time
}