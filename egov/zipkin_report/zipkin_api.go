package zipkin_report

import (
	"github.com/openzipkin/zipkin-go"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
)

func GetTracer(serviceName string, ip string, enpoitUrl string) *zipkin.Tracer {
	// create a reporter to be used by the tracer
	//if EnpoitUrl == "" {
	//	return nil
	//}
	reporter := httpreporter.NewReporter(enpoitUrl)

	// set-up the local endpoint for our service
	endpoint, _ := zipkin.NewEndpoint(serviceName, ip)

	// set-up our sampling strategy
	sampler := zipkin.NewModuloSampler(1)

	// initialize the tracer
	tracer, _ := zipkin.NewTracer(
		reporter,
		zipkin.WithLocalEndpoint(endpoint),
		zipkin.WithSampler(sampler),
	)
	return tracer
}

func GetTracerWithIps(serviceName string, ips []string, enpoitUrl string) *zipkin.Tracer {
	// create a reporter to be used by the tracer
	//if EnpoitUrl == "" {
	//	return nil
	//}
	reporter := httpreporter.NewReporter(enpoitUrl)

	// set-up the local endpoint for our service
	//endpoint, _ := zipkin.NewEndpoint(serviceName, ip)

	// set-up our sampling strategy
	sampler := zipkin.NewModuloSampler(1)
	//zipkin.WithLocalEndpoint(endpoint),
	//	zipkin.WithSampler(sampler),

	// initialize the tracer
	var opts []zipkin.TracerOption
	opts = append(opts, zipkin.WithSampler(sampler))
	for _, ip := range ips {
		endpoint, _ := zipkin.NewEndpoint(serviceName, ip)
		opts = append(opts, zipkin.WithLocalEndpoint(endpoint))

	}
	tracer, _ := zipkin.NewTracer(
		reporter,
		opts...
	)
	return tracer
}
