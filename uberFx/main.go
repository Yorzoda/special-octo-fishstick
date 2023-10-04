package main

import (
	"github.com/special-octo-fishtick/uberFx/controller"
	"go.uber.org/fx"
	"net/http"
)

func main() {
	fx.New(
		fx.Provide(controller.NewHandler),
		fx.Provide(controller.NewHTTPServer),
		fx.Invoke(func(server *http.Server) {}),
	).Run()

}
