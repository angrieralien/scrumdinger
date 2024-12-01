// Package all binds all the routes into the specified app.
package all

import (
	"time"

	"github.com/angrieralien/scrumdinger/app/domain/checkapp"
	"github.com/angrieralien/scrumdinger/app/domain/homeapp"
	"github.com/angrieralien/scrumdinger/app/domain/userapp"
	"github.com/angrieralien/scrumdinger/app/sdk/mux"
	"github.com/angrieralien/scrumdinger/business/domain/homebus"
	"github.com/angrieralien/scrumdinger/business/domain/homebus/stores/homedb"
	"github.com/angrieralien/scrumdinger/business/domain/userbus"
	"github.com/angrieralien/scrumdinger/business/domain/userbus/stores/usercache"
	"github.com/angrieralien/scrumdinger/business/domain/userbus/stores/userdb"
	"github.com/angrieralien/scrumdinger/business/sdk/delegate"
	"github.com/angrieralien/scrumdinger/foundation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) {

	// Construct the business domain packages we need here so we are using the
	// sames instances for the different set of domain apis.
	delegate := delegate.New(cfg.Log)
	userBus := userbus.NewBusiness(cfg.Log, delegate, usercache.NewStore(cfg.Log, userdb.NewStore(cfg.Log, cfg.DB), time.Minute))
	homeBus := homebus.NewBusiness(cfg.Log, userBus, delegate, homedb.NewStore(cfg.Log, cfg.DB))

	checkapp.Routes(app, checkapp.Config{
		Build: cfg.Build,
		Log:   cfg.Log,
		DB:    cfg.DB,
	})

	homeapp.Routes(app, homeapp.Config{
		Log:        cfg.Log,
		UserBus:    userBus,
		HomeBus:    homeBus,
		AuthClient: cfg.AuthClient,
	})

	userapp.Routes(app, userapp.Config{
		Log:        cfg.Log,
		UserBus:    userBus,
		AuthClient: cfg.AuthClient,
	})
}
