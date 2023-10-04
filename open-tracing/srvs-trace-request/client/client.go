package main

import (
	"context"
	"github.com/special-octo-fishstick/open-tracing/pkg"
	"net/http"
	"net/url"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	xhttp "github.com/yurishkuro/opentracing-tutorial/go/lib/http"
)

func main() {

	tracer, closer := pkg.InitJager("inportant-srv")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	flag := os.Args[1]

	span := tracer.StartSpan("question")
	span.SetTag("inportant-srv", flag)
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	helloStr := formatString(ctx, flag)
	printHello(ctx, helloStr)
}

func formatString(ctx context.Context, flag string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "formatString")
	defer span.Finish()

	v := url.Values{}
	v.Set("useless", flag)
	url := "http://localhost:8082/format?" + v.Encode()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	resp, err := xhttp.Do(req)
	if err != nil {
		ext.LogError(span, err)
		panic(err.Error())
	}

	uselessResp := string(resp)

	span.LogFields(
		log.String("event", "useless msg"),
		log.String("value", uselessResp),
	)

	return uselessResp
}

func printHello(ctx context.Context, helloStr string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "printHello")
	defer span.Finish()

	v := url.Values{}
	v.Set("helloFellaKids", helloStr)
	url := "http://localhost:8082/publish?" + v.Encode()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

	if _, err := xhttp.Do(req); err != nil {
		panic(err.Error())
	}
}
