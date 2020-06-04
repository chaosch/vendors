package opentracing_api

import (
	"github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go"
	zipkinot "github.com/openzipkin/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go/reporter"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

///////////////////////////////////////////////

func NewTracer(reportUrl, serviceName, endPoint *string) (reporter.Reporter, opentracing.Tracer) {
	if reportUrl == nil || endPoint == nil {
		return nil,nil
	}
	reporter := zipkinhttp.NewReporter(*reportUrl)

	// create our local service endpoint
	endpoint, err := zipkin.NewEndpoint(*serviceName, *endPoint)
	if err != nil {
		return nil,nil
	}

	// initialize our tracer
	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		return nil,nil
	}

	// use zipkin-go-opentracing to wrap our tracer
	tracer := zipkinot.Wrap(nativeTracer)

	// optionally set as Global OpenTracing tracer instance
	opentracing.SetGlobalTracer(tracer)

	return reporter,tracer

}

/////////////////////////////////
