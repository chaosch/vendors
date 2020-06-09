package main

import (
	"context"
	"egov/zipkin_report"
	"fmt"
	pb "getsequence-grpc/proto-api"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

const(
	address="192.168.6.12:9028"
)
func main(){

	tracergrpcclient:=zipkin_report.GetTracer("dal","192.168.6.12:9028","http://192.168.4.160:9411/api/v2/spans")

	conn, err := grpc.Dial(address, grpc.WithInsecure(),grpc.WithStatsHandler(zipkingrpc.NewClientHandler(tracergrpcclient)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		fmt.Println(err)
	}
	defer conn.Close()
	c:=pb.NewGetSequenceServiceClient(conn)
	rq:=&pb.Request{OutBuf:"tab_affairs"}
	type MD map[string][]string
	md := metadata.New(map[string]string{"ChipId": "181"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	res,err1:=c.GrpcGetSequence(ctx,rq)
	if err1!=nil{
		fmt.Println(err1)
	}else {
		fmt.Println(string(res.OutBuf))
	}
	cc:=pb.NewGetSequenceIntServiceClient(conn)
	reqs:=&pb.RequestInt{OutInt:"tab_user"}
	resint,err2:=cc.GrpcGetSequenceInt(context.Background(),reqs)
	if err2!=nil{
		fmt.Println(err2)
	}else{
		fmt.Println(resint.OutInt)
	}

}
