package main

import (
	"context"
	"github.com/special-octo-fishstick/uberFx/serviceA"
	"github.com/special-octo-fishstick/uberFx/serviceB"
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

// We also can provide third and fourth service by constructors
// But we want implement them from firs two services
func provider() fx.Option {
	return fx.Provide(
		serviceA.NewFirstSrv,
		serviceB.NewSecondSrv,
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
