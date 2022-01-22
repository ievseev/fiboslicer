package main

import (
	"context"
	"fiboslicer/internal/server"
	pb "fiboslicer/pkg/github.com/ievseev/fiboslicer"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	go func() {
		// mux
		mux := runtime.NewServeMux()

		//register
		pb.RegisterFiboSlicerHandlerServer(context.Background(), mux, &server.FiboSlicerServer{})

		//http server
		log.Fatalln(http.ListenAndServe("localhost:81", mux))
	}()

	listner, err := net.Listen("tcp", "localhost:80")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterFiboSlicerServer(grpcServer, &server.FiboSlicerServer{})

	err = grpcServer.Serve(listner)
	if err != nil {
		log.Fatalln(err)
	}
}
