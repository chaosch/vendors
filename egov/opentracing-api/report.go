package opentracing_api

import (
	"github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go"
	zipkinot "github.com/openzipkin/zipkin-go-opentracing"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

///////////////////////////////////////////////

func NewTracer(reportUrl, serviceName, endPoint *string) opentracing.Tracer {
	if reportUrl == nil || endPoint == nil {
		return nil
	}
	reporter := zipkinhttp.NewReporter(*reportUrl)
	defer reporter.Close()

	// create our local service endpoint
	endpoint, err := zipkin.NewEndpoint(*serviceName, *endPoint)
	if err != nil {
		return nil
	}

	// initialize our tracer
	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		return nil
	}

	// use zipkin-go-opentracing to wrap our tracer
	tracer := zipkinot.Wrap(nativeTracer)

	// optionally set as Global OpenTracing tracer instance
	opentracing.SetGlobalTracer(tracer)

	return tracer

}

/////////////////////////////////
