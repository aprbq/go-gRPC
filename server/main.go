package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"server/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	var s *grpc.Server

	tls := flag.Bool("tls", false, "use a secure TLS connection")
	flag.Parse()

	if *tls {
		certFile := "../tls/server.crt"
		keyFile := "../tls/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatal(err)
		}
		s = grpc.NewServer(grpc.Creds(creds))
	} else {
		s = grpc.NewServer()
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterCalculatorServer(s, services.NewCalculatorServer())

	fmt.Print("gRPC server listening on port 50051")

	if *tls {
		fmt.Println(" with TLS")
	} else {
		fmt.Println()
	}

	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
