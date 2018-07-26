package main

import (
	"log"
	"google.golang.org/grpc"
	pb "getsequence-grpc/proto-api"
	"context"
	"fmt"
)

const(
	address="192.168.5.245: 9025"
)
func main(){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		fmt.Println(err)
	}
	defer conn.Close()
	c:=pb.NewGetSequenceServiceClient(conn)
	rq:=&pb.Request{OutBuf:"tab_user"}
	res,err1:=c.GrpcGetSequence(context.Background(),rq)
	if err1!=nil{
		fmt.Println(err1)
	}
	fmt.Println(string(res.OutBuf))
}
