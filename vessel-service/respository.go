package main

import (
	"errors"

	"gopkg.in/mgo.v2"

	pb "github.com/agonper/shipper/vessel-service/proto/vessel"
)

const (
	dbName           = "shippy"
	vesselCollection = "vessels"
)

type Repository interface {
	Create(*pb.Vessel) error
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Close()
}

type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessels []*pb.Vessel

	err := repo.collection().Find(nil).All(&vessels)
	if err != nil {
		return nil, err
	}

	for _, vessel := range vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}

func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}
