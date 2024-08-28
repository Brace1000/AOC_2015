package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file and read it as a string
	file, err := os.ReadFile("dimensions.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Calculate the total surface area required for the wrapping paper
	dimensions := string(file)
	totalSurfaceArea := wrapperPaperSA(dimensions)
	fmt.Printf("The elves need %v square feet of wrapping paper.\n", totalSurfaceArea)
}

func wrapperPaperSA(dimensions string) int {
	var totalArea int

	// Separate the dimensions using whitespace
	for _, dimension := range strings.Fields(dimensions) {
		dimParts := strings.Split(dimension, "x")
		if len(dimParts) != 3 {
			log.Printf("Invalid dimension format: %s\n", dimension)
			continue
		}

		// Convert dimensions to integers
		length, err := strconv.Atoi(dimParts[0])
		if err != nil {
			log.Printf("Invalid length value: %s\n", dimParts[0])
			continue
		}
		width, err := strconv.Atoi(dimParts[1])
		if err != nil {
			log.Printf("Invalid width value: %s\n", dimParts[1])
			continue
		}
		height, err := strconv.Atoi(dimParts[2])
		if err != nil {
			log.Printf("Invalid height value: %s\n", dimParts[2])
			continue
		}

		// Calculate surface area and smallest side area
		surfaceArea := 2*length*width + 2*width*height + 2*height*length
		area1 := length * width
		area2 := width * height
		area3 := height * length

		// Find the smallest area
		smallestArea := area1
		if area2 < smallestArea {
			smallestArea = area2
		}
		if area3 < smallestArea {
			smallestArea = area3
		}

		// Add the surface area and smallest area to the total
		totalArea += surfaceArea + smallestArea
	}

	return totalArea
}
