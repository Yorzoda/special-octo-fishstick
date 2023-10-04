package main

import (
	"log"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/special-octo-fishstick/open-tracing/pkg"
)

func main() {
	http.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
		tracer, err := pkg.InitJager("formatter?")
		if err != nil {
			panic(err)
		}
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		span := tracer.StartSpan("publish", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		helloStr := r.FormValue("useless")
		println(helloStr)
	})

	log.Fatal(http.ListenAndServe(":8082", nil))
}
