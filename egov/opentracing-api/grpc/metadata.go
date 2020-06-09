package grpc

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strings"
)

func PrepareContext(Ctx context.Context, md map[string][]string) context.Context {

	Md := metadata.MD{}

	for k, v := range md {
		Md[strings.ToLower(k)] = v
	}

	//MD["ChipId"] = []string{strconv.FormatInt(dbe.SystemConfig.ChipId, 10)}

	ctx := metadata.NewOutgoingContext(Ctx, Md)

	return ctx
}

func AddSpanFromContext(oCtx, Ctx context.Context) context.Context {
	span := opentracing.SpanFromContext(Ctx)

	req, _ := http.NewRequest("", "", nil)

	carrier := opentracing.HTTPHeadersCarrier(req.Header)

	span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, carrier)

	md := metadata.MD(req.Header)

	Md := metadata.MD{}

	for k, v := range md {
		Md[strings.ToLower(k)] = v
	}

	ctx := metadata.NewOutgoingContext(oCtx, Md)

	return ctx
}
