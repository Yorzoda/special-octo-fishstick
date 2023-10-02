package main

import (
	"errors"
	"github.com/opentracing/opentracing-go/log"
	"open-tracing/pkg"
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

	//multiple log field initialization
	span.LogFields(
		log.String("event-log-fields", "useless log"),
		log.String("another-event", "lorem where's ipsum?"),
	)

	//single log field initialization
	span.LogKV("single-event", "gift of nothing")

	span.Finish()
}
