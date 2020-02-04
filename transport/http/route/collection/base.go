package collection

import (
	"github.com/payfazz/go-apt/pkg/fazzrouter"
	"github.com/payfazz/tango/transport/container"
	"github.com/payfazz/tango/transport/http/controller/base"
)

// RouteBase is a base router
func RouteBase(r *fazzrouter.Route, app *container.AppContainer) {
	r.Get("ping", base.Ping())
}