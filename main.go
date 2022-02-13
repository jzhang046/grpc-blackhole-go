package main

import (
	"flag"
	"log"
	"net"
	"os"

	pb "github.com/jzhang046/grpc-blackhole/blackhole"
	"github.com/jzhang046/grpc-blackhole/server"
	"google.golang.org/grpc"
)

const (
	defaultListenAddress = "localhost:9010"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

func main() {

	listenAddress := flag.String(
		"listen-addr",
		defaultListenAddress,
		"address that the gRPC service would listen to",
	)

	flag.Parse()

	lis, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("Starting gRPC server on [%s] with PID [%d]", *listenAddress, os.Getpid())
	}

	var ops []grpc.ServerOption

	grpcServer := grpc.NewServer(ops...)
	pb.RegisterBlackHoleServer(grpcServer, server.New())
	grpcServer.Serve(lis)
}
