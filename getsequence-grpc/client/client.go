package main

import (
	"log"
	"google.golang.org/grpc"
	pb "getsequence-grpc/proto-api"
	"context"
	"fmt"
)

const(
	address="localhost: 9028"
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
	res,err1:=c.GrpcGetSequence(context.Background(),rq)
	if err1!=nil{
		fmt.Println(err1)
	}
	fmt.Println(string(res.OutBuf))
	cc:=pb.NewGetSequenceIntServiceClient(conn)
	reqs:=&pb.RequestInt{OutInt:"tab_user"}
	resint,err2:=cc.GrpcGetSequenceInt(context.Background(),reqs)
	if err2!=nil{
		fmt.Println(err2)
	}else{
		fmt.Println(resint.OutInt)
	}

}
