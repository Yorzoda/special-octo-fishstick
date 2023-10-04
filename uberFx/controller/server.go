package controller

import (
	"context"
	"go.uber.org/fx"
	"net"
	"net/http"
)

func NewHTTPServer(lc fx.Lifecycle, h *echoHandler) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: h}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			go func() {
				err = srv.Serve(ln)
				if err != nil {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
