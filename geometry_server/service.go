package main

import (
	"context"
	"log"
	"net"

	gs "github.com/KernelGamut32/grpc-geometry/geometry"
	"google.golang.org/grpc"
)

type server struct {
	gs.UnimplementedGeometryServer
}

func (s *server) ComputeArea(ctx context.Context, r *gs.AreaRequest) (*gs.AreaResponse, error) {
	log.Printf("received request to compute area with radius value: %f", r.Circle.Radius)
	area := ComputeArea(r.Circle.Radius)
	return &gs.AreaResponse{Area: area}, nil
}

func (s *server) ComputePerimeter(ctx context.Context, r *gs.PerimeterRequest) (*gs.PerimeterResponse, error) {
	log.Printf("received request to compute perimeter with radius value: %f", r.Circle.Radius)
	perim := ComputePerimeter(r.Circle.Radius)
	return &gs.PerimeterResponse{Perimeter: perim}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	gs.RegisterGeometryServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
