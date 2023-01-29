package main

import "fmt"

func main() {
	fmt.Printf("%v", convertTemperature(36.5))
	fmt.Printf("%v", convertTemperature(122.11))
}
func convertTemperature(celsius float64) []float64 {
	kelvin := celsius + 273.15
	return []float64{kelvin, celsius*1.8 + 32}
}
