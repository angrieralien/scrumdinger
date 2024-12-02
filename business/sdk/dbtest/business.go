package dbtest

import (
	"time"

	"github.com/angrieralien/scrumdinger/business/domain/scrumbus"
	"github.com/angrieralien/scrumdinger/business/domain/scrumbus/stores/scrumdb"

	"github.com/angrieralien/scrumdinger/business/domain/userbus"
	"github.com/angrieralien/scrumdinger/business/domain/userbus/stores/usercache"
	"github.com/angrieralien/scrumdinger/business/domain/userbus/stores/userdb"

	"github.com/angrieralien/scrumdinger/business/sdk/delegate"
	"github.com/angrieralien/scrumdinger/foundation/logger"
	"github.com/jmoiron/sqlx"
)

// BusDomain represents all the business domain apis needed for testing.
type BusDomain struct {
	Delegate *delegate.Delegate
	Scrum    *scrumbus.Business
	User     *userbus.Business
}

func newBusDomains(log *logger.Logger, db *sqlx.DB) BusDomain {
	delegate := delegate.New(log)
	userBus := userbus.NewBusiness(log, delegate, usercache.NewStore(log, userdb.NewStore(log, db), time.Hour))
	scrumBus := scrumbus.NewBusiness(log, userBus, delegate, scrumdb.NewStore(log, db))

	return BusDomain{
		Delegate: delegate,
		Scrum:    scrumBus,
		User:     userBus,
	}
}
