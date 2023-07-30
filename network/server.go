package network

import (
	"fmt"
	pb "github.com/qcodelabsllc/exag/email/gen"
	svc "github.com/qcodelabsllc/exag/email/services"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func StartServer() {
	// create grpc server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)

	// register services
	pb.RegisterEmailServiceServer(grpcServer, &svc.EmailServiceImpl{})

	// setup reflection
	reflection.Register(grpcServer)

	// define address to listen on
	address := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))

	// listen on address
	if listener, err := net.Listen("tcp", address); err != nil {
		log.Fatalf("unable to listen on %s: %v", address, err)
	} else {
		// start server
		log.Printf("gRPC server started on %s", address)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("unable to serve: %v", err)
		}
	}
}
