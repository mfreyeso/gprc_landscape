package main

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "health/client/clinical"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewMeasureServiceClient(conn)

	diastole := float32(102.0)
	systole := float32(89.0)

	idMeasure := "20002"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CalculateBloodPressure(ctx, &pb.BloodMeasure{Diastole: diastole, Systole: systole,
		Datetime: ptypes.TimestampNow(), Id: idMeasure})

	if err != nil {
		log.Fatalf("Fail trying to calculate Blood Pressure")
	}
	log.Printf("Sucessful operation, the blood presure is: %v", r.Measure)

}
