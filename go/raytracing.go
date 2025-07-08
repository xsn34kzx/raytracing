package main 

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Image 
	imageWidth := 256;
	imageHeight := 256;

	// Render
	renderStart := time.Now()

	fmt.Printf("P3\n%v %v\n255\n", imageWidth, imageHeight)
	for j := range imageHeight {
		for i := range imageWidth {
			r := uint8(float64(i) / float64(imageWidth - 1) * 255)
			g := uint8(float64(j) / float64(imageHeight - 1) * 255)
			b := 0

			fmt.Printf("%v %v %v\n", r, g, b)
		}
	}

	renderEnd := time.Since(renderStart)

	log.Printf("Finished in %v", renderEnd.String())
}
