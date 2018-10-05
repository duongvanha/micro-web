package main

import (
	"context"
	pb "micro-web/proto"
)

type TransferServiceServer struct{}

func (t *TransferServiceServer) Send(ctx context.Context, data *pb.Data) (*pb.Data, error) {

	return &pb.Data{Id: 123, Text: "something"}, nil

}
