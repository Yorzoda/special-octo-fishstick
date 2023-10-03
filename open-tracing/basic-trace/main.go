package main

import (
	"context"
	"errors"
	"trace-services/pkg"

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

	//Global trace initialization
	opentracing.SetGlobalTracer(tracer)

	span := tracer.StartSpan("UwU-trace")

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	span.SetTag("event", "UwU")
	//multiple log field initialization
	span.LogFields(
		log.String("event-log-fields", "useless log"),
		log.String("another-event", "lorem where's ipsum?"),
	)

	//Calling children span
	// multiPrintFromParent(span)
	// singlePrintFromParent(span)
	multiPrintFromCtx(ctx)

	//single log field initialization
	span.LogKV("single-event", "gift of nothing", "another-singl-event", "value of it")

	span.Finish()
}

// func multiPrintFromParent(rootSpan opentracing.Span) {
// 	span := rootSpan.Tracer().StartSpan("mulit-child-span", opentracing.ChildOf(rootSpan.Context()))
// 	defer span.Finish()

// 	span.LogFields(
// 		log.String("event-info", "this is very informative"),
// 		log.Event("event without key"),
// 	)
// }

// func singlePrintFromParent(rootSpan opentracing.Span) {
// 	span := rootSpan.Tracer().StartSpan("single-child-span", opentracing.ChildOf(rootSpan.Context()))
// 	defer span.Finish()

// 	span.LogKV("event-sinle", "one is the lonliest number")
// }

func multiPrintFromCtx(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "multi-child-ctx-span")
	defer span.Finish()

	span.LogFields(
		log.String("event-info", "this is very informative"),
		log.Event("event without key"),
	)
}
