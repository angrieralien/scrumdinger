package scrumapp

import (
	"net/http"

	"github.com/angrieralien/scrumdinger/app/sdk/auth"
	"github.com/angrieralien/scrumdinger/app/sdk/authclient"
	"github.com/angrieralien/scrumdinger/app/sdk/mid"
	"github.com/angrieralien/scrumdinger/business/domain/scrumbus"
	"github.com/angrieralien/scrumdinger/business/domain/userbus"
	"github.com/angrieralien/scrumdinger/foundation/logger"
	"github.com/angrieralien/scrumdinger/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log        *logger.Logger
	UserBus    *userbus.Business
	ScrumBus   *scrumbus.Business
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.AuthClient)
	ruleAuthorizeScrum := mid.AuthorizeScrum(cfg.AuthClient, cfg.ScrumBus)
	ruleUserOnly := mid.Authorize(cfg.AuthClient, auth.RuleUserOnly)

	api := newApp(cfg.ScrumBus)

	app.HandlerFunc(http.MethodGet, version, "/scrums", api.queryByUserID, authen, ruleUserOnly)
	app.HandlerFunc(http.MethodGet, version, "/scrums/{scrum_id}", api.queryByUserID, authen, ruleAuthorizeScrum)
	app.HandlerFunc(http.MethodPost, version, "/scrums", api.create, authen, ruleUserOnly)
	app.HandlerFunc(http.MethodPut, version, "/scrums/{scrum_id}", api.update, authen, ruleAuthorizeScrum)
	app.HandlerFunc(http.MethodDelete, version, "/scrums/{scrum_id}", api.delete, authen, ruleAuthorizeScrum)
}
