package main 

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Invalid amount of arguments!")
	}

	// Image 
	imageWidth := 256;
	imageHeight := 256;

	// Render
	renderStart := time.Now()

	switch os.Args[1] {
	case "hello_world":
		helloWorldScene(imageWidth, imageHeight)
	default:
		fmt.Println("Not a valid scene!")
	}

	renderEnd := time.Since(renderStart)

	log.Printf("Finished in %v", renderEnd.String())
}
