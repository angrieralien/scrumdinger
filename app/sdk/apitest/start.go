package apitest

import (
	"net/http/httptest"
	"testing"

	authbuild "github.com/angrieralien/scrumdinger/api/services/auth/build/all"
	scrumbuild "github.com/angrieralien/scrumdinger/api/services/scrums/build/all"
	"github.com/angrieralien/scrumdinger/app/sdk/auth"
	"github.com/angrieralien/scrumdinger/app/sdk/authclient"
	"github.com/angrieralien/scrumdinger/app/sdk/mux"
	"github.com/angrieralien/scrumdinger/business/sdk/dbtest"
)

// New initialized the system to run a test.
func New(t *testing.T, testName string) *Test {
	db := dbtest.New(t, testName)

	// -------------------------------------------------------------------------

	auth, err := auth.New(auth.Config{
		Log:       db.Log,
		DB:        db.DB,
		KeyLookup: &KeyStore{},
	})
	if err != nil {
		t.Fatal(err)
	}

	// -------------------------------------------------------------------------

	server := httptest.NewServer(mux.WebAPI(mux.Config{
		Log: db.Log,
		DB:  db.DB,
		AuthConfig: mux.AuthConfig{
			Auth: auth,
		},
	}, authbuild.Routes()))

	authClient := authclient.New(db.Log, server.URL)

	// -------------------------------------------------------------------------

	mux := mux.WebAPI(mux.Config{
		Log: db.Log,
		DB:  db.DB,
		ScrumdingerConfig: mux.ScrumdingerConfig{
			AuthClient: authClient,
		},
	}, scrumbuild.Routes())

	return &Test{
		DB:   db,
		Auth: auth,
		mux:  mux,
	}
}
