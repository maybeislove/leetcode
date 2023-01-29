package main

import (
	"math"
	"time"
)

func main() {
	println(daysBetweenDates("2019-06-29", "2019-06-30"))
}
func daysBetweenDates(date1 string, date2 string) int {
	timeFormat := "2006-01-02"
	d1, _ := time.Parse(timeFormat, date1)
	d2, _ := time.Parse(timeFormat, date2)
	sub := d2.Sub(d1)
	abs := math.Abs(sub.Hours())
	return int(abs / 24)
}
