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

type serve struct {
}

func (s *serve) CalculateBloodPressure(ctx context.Context,
	in *pb.BloodMeasure) (*pb.BloodPressure, error) {

	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Operation ID", err)
	}
	var operationId = out.String()
	var result = in.Diastole * in.Systole
	log.Printf("Operation %v : %v - calculated.", operationId, result)
	return &pb.BloodPressure{Id:out.String(), Measure:result, Datetime:ptypes.TimestampNow()},
		status.New(codes.OK, "").Err()
}
