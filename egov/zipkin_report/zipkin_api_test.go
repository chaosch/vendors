package zipkin_report

import "testing"

func TestGetTracer(t *testing.T) {
	trace := FileCore.GetTracer("FileService", GetAddress(), y.EndpointUrl)
	mid := zipkinhttp.NewServerMiddleware(trace, zipkinhttp.TagResponseSize(true))
	fmt.Println(mid)
}