package link_management

import (
	"log"

	pb "github.com/italorfeitosa/url-shortner-mvp/services/link_management/external/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient() pb.LinkManagementServiceClient {
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Could not connect:", err)
	}

	return pb.NewLinkManagementServiceClient(cc)
}
