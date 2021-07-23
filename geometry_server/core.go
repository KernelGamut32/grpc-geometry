package main

import "math"

func ComputeArea(radius float32) float32 {
	return math.Pi * radius * radius
}

func ComputePerimeter(radius float32) float32 {
	return 2 * math.Pi * radius
}
