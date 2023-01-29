package main

import "time"

func main() {

}

func dayOfYear(date string) int {
	parse, _ := time.Parse("2006-01-02", date)
	return parse.YearDay()
}
