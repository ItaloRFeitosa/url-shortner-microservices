package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/italorfeitosa/url-shortner-mvp/libs/event"
	"github.com/italorfeitosa/url-shortner-mvp/services/link_management/core"
	"github.com/italorfeitosa/url-shortner-mvp/services/link_management/external/database"
	"github.com/italorfeitosa/url-shortner-mvp/services/link_management/external/encoder"
	g "github.com/italorfeitosa/url-shortner-mvp/services/link_management/external/grpc"
	pb "github.com/italorfeitosa/url-shortner-mvp/services/link_management/external/grpc/proto"
	grpc "google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func setupHandler() *g.Handler {
	db := database.NewDB()
	repo := database.NewLinkRepository(db)
	broker := event.NewSNSPublisher(event.SNSPublisherConfig{
		Topics: map[string]string{
			core.LinkCreatedTopic: "",
		},
	})
	u := core.NewLinkService(repo, encoder.NewEncoder(), broker)
	return g.NewHandler(u)
}

func main() {
	flag.Parse()

	h := setupHandler()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterLinkManagementServiceServer(grpcServer, h)
	log.Println("grpc server running in ", fmt.Sprintf("localhost:%d", *port))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
