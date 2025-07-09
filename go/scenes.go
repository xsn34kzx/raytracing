package main

import (
	"fmt"
) 

func helloWorldScene(width, height int) {
	fmt.Printf("P3\n%v %v\n255\n", width, height)
	for j := range height {
		v := float64(j) / float64(height - 1)
		for i := range width {
			u := float64(i) / float64(width - 1)

			r := uint8(u * 255)
			g := uint8(v * 255)
			b := 0

			fmt.Printf("%v %v %v\n", r, g, b)
		}
	}
}

func gradientScene(width, height int) {
	aspectRatio := float64(width) / float64(height)
	focalLength := 1.0

	viewportHeight := 2.0
	viewportWidth := viewportHeight * aspectRatio

	var cameraCenter Vec3

	pixelDeltaU := viewportWidth / float64(width)
	pixelDeltaV := -viewportHeight / float64(height)

	viewportUpperLeft := Vec3{
		cameraCenter.X - viewportWidth / 2,
		cameraCenter.Y + viewportHeight / 2,
		cameraCenter.Z - focalLength,
	}

	pixelOriginLocation := viewportUpperLeft
	pixelOriginLocation.X += pixelDeltaU / 2
	pixelOriginLocation.Y += pixelDeltaV / 2

	ray := Ray{
		Origin: cameraCenter,
	}

	fmt.Printf("P3\n%v %v\n255\n", width, height)
	for j := range height {
		v := float64(j) * pixelDeltaV
		for i := range width {
			u := float64(i) * pixelDeltaU 

			pixelCenter := pixelOriginLocation
			pixelCenter.X += u
			pixelCenter.Y += v

			ray.Direction = Subtract(pixelCenter, cameraCenter)
			a := (UnitVector(ray.Direction).Y + 1) / 2

			startColor := Vec3{1, 1, 1}
			endColor := Vec3{0.5, 0.7, 1}

			startColor.Scale(1 - a)
			endColor.Scale(a)

			finalColor := Add(startColor, endColor)

			r := uint8(finalColor.X * 255)
			g := uint8(finalColor.Y * 255)
			b := uint8(finalColor.Z * 255)

			fmt.Printf("%v %v %v\n", r, g, b)
		}
	}
}
