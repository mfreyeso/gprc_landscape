package main

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	pb "health/service/clinical"
	"log"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	meanMeasure = 120.0 / 80.0
	highMeasure = 140.0 / 90

	goodMeasure   = "GOOD MEASURE"
	preHypMeasure = "PRE-HYPERTENSION"
	hypMeasure    = "HYPERTENSION"
)

type server struct {
	pb.UnimplementedMeasureServiceServer
}

func (s *server) CalculateBloodPressure(ctx context.Context,
	in *pb.BloodMeasure) (*pb.BloodPressure, error) {

	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Operation ID", err)
	}
	var operationId = out.String()
	measure := in.Systole / in.Diastole
	var result = ""
	if measure < meanMeasure {
		result = goodMeasure
	}
	if measure >= meanMeasure && measure < highMeasure {
		result = preHypMeasure
	} else {
		result = hypMeasure
	}
	log.Printf("Operation %v : %v - calculated.", operationId, result)
	return &pb.BloodPressure{Id: out.String(), Measure: result, Datetime: ptypes.TimestampNow()},
		status.New(codes.OK, "").Err()
}
