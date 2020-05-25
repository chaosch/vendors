package main

import (
	"context"
	"fmt"
	pb "getsequence-grpc/proto-api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

const(
	address="192.168.6.12:9028"
)
func main(){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
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
