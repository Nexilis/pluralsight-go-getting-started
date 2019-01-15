package main

import (
	"fmt"
)

func reportFirst() {
	var plantCapacities []float64

	plantCapacities = []float64{30,30,30,60,60,100}

	var capacity float64
	for i := 0; i < len(plantCapacities); i++ {
		capacity += plantCapacities[i]
	}

	var gridLoad = 75.

	utilization := gridLoad / capacity

	fmt.Printf("%-20s%.0f\n", "Capacity", capacity)
	fmt.Printf("%-20s%.0f\n", "Load", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization", utilization)
}