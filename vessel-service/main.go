package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/agonper/shipper/vessel-service/proto/vessel"
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
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	vesselService := &service{session}
	vesselService.GetRepo().Create(&pb.Vessel{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500})

	pb.RegisterVesselServiceHandler(srv.Server(), vesselService)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
