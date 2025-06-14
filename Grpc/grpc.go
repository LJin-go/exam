package Grpc

//import (
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials/insecure"
//	"log"
//)
//
//func Grpc() *grpc.ClientConn {
//	conn, err := grpc.NewClient("", grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		log.Fatalf("did not connect: %v", err)
//	}
//	return conn
//}

//func InitGrpc() {
//	conn := Grpc("127.0.0.1:10002")
//	global.GreeterClient := pb.NewGreeterClient(conn)
//}
