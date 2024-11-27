package productapp

import (
	"net/http"

	"github.com/angrieralien/scrumdinger/app/sdk/auth"
	"github.com/angrieralien/scrumdinger/app/sdk/authclient"
	"github.com/angrieralien/scrumdinger/app/sdk/mid"
	"github.com/angrieralien/scrumdinger/business/domain/productbus"
	"github.com/angrieralien/scrumdinger/business/domain/userbus"
	"github.com/angrieralien/scrumdinger/foundation/logger"
	"github.com/angrieralien/scrumdinger/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log        *logger.Logger
	UserBus    *userbus.Business
	ProductBus *productbus.Business
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.AuthClient)
	ruleAny := mid.Authorize(cfg.AuthClient, auth.RuleAny)
	ruleUserOnly := mid.Authorize(cfg.AuthClient, auth.RuleUserOnly)
	ruleAuthorizeProduct := mid.AuthorizeProduct(cfg.AuthClient, cfg.ProductBus)

	api := newApp(cfg.ProductBus)

	app.HandlerFunc(http.MethodGet, version, "/products", api.query, authen, ruleAny)
	app.HandlerFunc(http.MethodGet, version, "/products/{product_id}", api.queryByID, authen, ruleAuthorizeProduct)
	app.HandlerFunc(http.MethodPost, version, "/products", api.create, authen, ruleUserOnly)
	app.HandlerFunc(http.MethodPut, version, "/products/{product_id}", api.update, authen, ruleAuthorizeProduct)
	app.HandlerFunc(http.MethodDelete, version, "/products/{product_id}", api.delete, authen, ruleAuthorizeProduct)
}
