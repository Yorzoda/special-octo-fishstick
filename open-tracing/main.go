package main

import (
	"context"
	"errors"
	"open-tracing/pkg"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func main() {
	tracer, closer := pkg.InitJager("trainee-trace")
	defer func() {
		if err := closer.Close(); err != nil {
			err = errors.New("closer.Close err:" + err.Error())
			log.Error(err)
		}
	}()

	span := tracer.StartSpan("UwU-trace")
	span.SetTag("event", "UwU")
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	//multiple log field initialization
	span.LogFields(
		log.String("event-log-fields", "useless log"),
		log.String("another-event", "lorem where's ipsum?"),
	)
	printFromCtx(ctx)
	//single log field initialization
	span.LogKV("single-event", "gift of nothing")

	span.Finish()
}

func printFromCtx(ctx context.Context) {
	spam, _ := opentracing.StartSpanFromContext(ctx, "print-from-context")
	defer spam.Finish()

	spam.LogFields(
		log.String("event-info", "this is very informative"),
		log.Event("event without key"),
	)
}
