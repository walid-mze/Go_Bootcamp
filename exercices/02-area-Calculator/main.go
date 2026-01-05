package main

import "log"

func calculateArea(length float64, width float64) float64 {
	if length < 0 || width < 0 {
		log.Fatal("length and width must be positive values")
	}
	return length * width
}
func main() {
	area := calculateArea(5.0, 3.0)
	log.Println("Area in meters :", area)
	log.Println("Area in foot :", area*3.28)
}
