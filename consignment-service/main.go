package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/agonper/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/agonper/shipper/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
)

const (
	defaultDatabaseHost = "localhost:27017"
)

func main() {
	databaseHost := os.Getenv("DB_HOST")

	if databaseHost == "" {
		databaseHost = defaultDatabaseHost
	}

	session, err := CreateSession(databaseHost)
	defer session.Close()

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", databaseHost, err)
	}

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
