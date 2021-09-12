package pkg

import (
	"google.golang.org/grpc"
	"log"
)

func Conn() *grpc.ClientConn{
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("discount:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	return conn
}
