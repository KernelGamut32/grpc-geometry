package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	gs "github.com/KernelGamut32/grpc-geometry/geometry"
	"google.golang.org/grpc"
)

func main() {
	// Use context to establish a 1-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Set up a connection to the gRPC server
	opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	conn, err := grpc.DialContext(ctx, "localhost:50051", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Get a new instance of our client
	client := gs.NewGeometryClient(conn)

	var action, r string
	var radius float32

	// Expect something like "area <radius>"
	if len(os.Args) > 1 {
		action, r = os.Args[1], os.Args[2]
		r64, _ := strconv.ParseFloat(r, 32)
		radius = float32(r64)
	}

	// Call client.ComputeArea() or client.ComputerPerimeter() as appropriate.
	switch action {
	case "area":
		result, err := client.ComputeArea(ctx, &gs.AreaRequest{Circle: &gs.Circle{Radius: radius}})
		if err != nil {
			log.Fatalf("could not calculate area for circle: %v\n", err)
		}
		log.Printf("compute area returns: %.4f\n", result.Area)
	case "perimeter":
		result, err := client.ComputePerimeter(ctx, &gs.PerimeterRequest{Circle: &gs.Circle{Radius: radius}})
		if err != nil {
			log.Fatalf("could not calculate perimeter for circle: %v\n", err)
		}
		log.Printf("computer perimeter returns: %.4f\n", result.Perimeter)
	default:
		log.Fatalf("Syntax: go run [area|perimeter] RADIUS")
	}
}
