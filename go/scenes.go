package main

import (
	"fmt"
) 

func helloWorldScene(width, height int) {
	fmt.Printf("P3\n%v %v\n255\n", width, height)
	for j := range height {
		for i := range width {
			r := uint8(float64(i) / float64(width - 1) * 255)
			g := uint8(float64(j) / float64(height - 1) * 255)
			b := 0

			fmt.Printf("%v %v %v\n", r, g, b)
		}
	}
}
