package main

func main() {

}

func checkStraightLine(coordinates [][]int) bool {
	//check （y-y0）× (x1-x0)-（x-x0)×(y1 - y0) == 0
	n := len(coordinates)
	x0, y0 := coordinates[0][0], coordinates[0][1]
	x, y := coordinates[1][0]-x0, coordinates[1][1]-y0
	for i := 2; i < n; i++ {
		if (coordinates[i][1]-y0)*x-(coordinates[i][0]-x0)*y != 0 {
			return false
		}
	}
	return true
}
