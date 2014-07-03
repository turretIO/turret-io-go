// +build appengine

package turretIO

import (
	"appengine"
    "appengine/urlfetch"
	"net/http"
)

// NewAppEngineTurretIO is used to create a new TurretIO base instance compatible
// with Google App Engine and requires a extra context parameter
func NewAppEngineTurretIO(api_key string, api_secret string, ctx appengine.Context) *AppEngineTurretIO {
    t := &AppEngineTurretIO{}
	t.GAEContext = ctx
    t.TurretIO.Apikey = api_key
    t.TurretIO.Apisecret = api_secret
    return t
}

type AppEngineTurretIO struct {
	TurretIO
    GAEContext appengine.Context
}

func (aet *AppEngineTurretIO) GetHTTPClient() (*http.Client) {
	 return urlfetch.Client(aet.GAEContext)
}


