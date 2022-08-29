package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "github.com/mrbelka12000/grpc-gateway/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
)

type routeGuide struct {
	check []*pb.Feature
	pb.UnimplementedRouteGuideServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50511))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = "x509/server_cert.pem"
		}
		if *keyFile == "" {
			*keyFile = "x509/server_key.pem"
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRouteGuideServer(grpcServer, newRoute())
	fmt.Println("grpc server started")
	grpcServer.Serve(lis)
}

func newRoute() *routeGuide {

	return &routeGuide{
		check: []*pb.Feature{{
			Name:     1,
			Location: &pb.Point{},
		},
			{
				Name:     2,
				Location: nil,
			}},
	}
}

func (rg *routeGuide) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	var lastID int32
	for _, feature := range rg.check {
		if proto.Equal(feature.Location, point) {
			log.Println("poimal")
			return feature, nil
		}
		lastID = feature.Name
	}

	feature := &pb.Feature{
		Name:     lastID + 1,
		Location: point,
	}
	rg.check = append(rg.check, feature)

	return feature, nil
}
