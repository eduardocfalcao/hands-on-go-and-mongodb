package main

import (
	"fmt"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/eduardocfalcao/hands-on-go-and-mongodb/src/proto/sweatmgr"
	"github.com/eduardocfalcao/hands-on-go-and-mongodb/src/sweatmgr/config"
	"github.com/eduardocfalcao/hands-on-go-and-mongodb/src/sweatmgr/db"
	"github.com/eduardocfalcao/hands-on-go-and-mongodb/src/sweatmgr/logger"
	"github.com/eduardocfalcao/hands-on-go-and-mongodb/src/sweatmgr/service"
	"github.com/urfave/negroni"
)

func main() {
	logger.Init()
	config.Load()

	l := logger.Get()

	db, err := db.GetDB()
	if err != nil {
		panic(err)
	}

	l.Infof("Db Initialized: ", db)

	router := service.InitRouter()

	server := negroni.Classic()
	server.UseHandler(router)

	port := config.AppPort()
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	go GRPCServe()
	server.Run(addr)
}

func GRPCServe() {
	host := config.ReadEnvString("GRPC_HOST")
	port := config.ReadEnvInt("GRPC_PORT")
	tls := config.ReadEnvBool("TLS")
	certFile := config.ReadEnvString("CERT_FILE")
	keyFile := config.ReadEnvString("KEY_FILE")

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		logger.Get().Fatalf("Failed to listen on: %v", err)
	}
	var opts []grpc.ServerOption
	if tls {
		if certFile == "" {
			logger.Get().Fatalf("No certificate file specified")
		}
		if keyFile == "" {
			logger.Get().Fatalf("No key file specified")
		}

		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			logger.Get().Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	grpcServer := grpc.NewServer(opts...)
	s := service.GrpcServer{}

	pb.RegisterSweatMgrServiceServer(grpcServer, &s)

	logger.Get().Infof("Listening for gRPC on %s:%d", host, port)

	grpcServer.Serve(lis)

}
