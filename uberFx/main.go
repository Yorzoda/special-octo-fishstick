package main

import (
	"context"
	"github.com/special-octo-fishstick/uberFx/serviceA"
	"github.com/special-octo-fishstick/uberFx/serviceA/serviceC"
	"github.com/special-octo-fishstick/uberFx/serviceB"
	"github.com/special-octo-fishstick/uberFx/serviceB/ServiceD"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		provider(),
		fx.Invoke(invoker),
		fx.NopLogger,
	)

	app.Run()
}

func provider() fx.Option {
	return fx.Provide(
		serviceA.NewFirstSrv,
		serviceB.NewSecondSrv,
		serviceC.NewThirdService,
		ServiceD.NewForthService,
	)
}

func invoker(
	lc fx.Lifecycle,
	s1 *serviceA.FirsService,
	s2 *serviceB.SecondService,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			s1.FirstStartService()
			s2.SecondServerStart()
			s1.ThirdServiceStart()
			s2.ForthServiceStart()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
