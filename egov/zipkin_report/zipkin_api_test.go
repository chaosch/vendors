package zipkin_report

import (
	"fmt"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	"net/http"
	"testing"
)

func TestGetTracer(t *testing.T) {
	trace := GetTracer("dal", "localhost", "http://192.168.4.160:9411/api/v2/spans")
	mid := zipkinhttp.NewServerMiddleware(trace, zipkinhttp.TagResponseSize(true))
	c:=&http.Client{}

	fmt.Println(mid(c))
}