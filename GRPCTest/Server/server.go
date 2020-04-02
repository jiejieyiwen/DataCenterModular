package Server

import (
	"DataCenterModular/GRPCTest/Message"
	"context"
	"fmt"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	Message.UnimplementedDC_NotificationServer
}

func (pThis *server) Notify(ctx context.Context, req *Message.DC_Request) (*Message.DC_Respon, error) {
	fmt.Println(req.BGetNew)
	strData := "已经收到通知"
	fmt.Println(strData)
	res := &Message.DC_Respon{StrRespon: 1}
	return res, nil
}

type GRpcServer struct {
}

func (pThis *GRpcServer) Service(strPort string) error {
	lis, err := net.Listen("tcp", strPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gRpcServer := grpc.NewServer()
	Message.RegisterDC_NotificationServer(gRpcServer, &server{})
	reflection.Register(gRpcServer)
	if err := gRpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return err
}
