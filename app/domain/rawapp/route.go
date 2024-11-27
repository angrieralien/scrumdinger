package rawapp

import (
	"net/http"

	"github.com/angrieralien/scrumdinger/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(app *web.App) {
	const version = "v1"

	app.RawHandlerFunc(http.MethodGet, version, "/raw", rawHandler)
}
