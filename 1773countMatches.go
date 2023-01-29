package main

func main() {
	println(countMatches([][]string{
		{"phone", "blue", "pixel"},
		{"computer", "silver"},
		{"lenovo", "phone", "gold", "iphone"}},
		"color", "silver"))
}

var map1773 = map[string]int{
	"type":  0,
	"color": 1,
	"name":  2,
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	if _, ok := map1773[ruleKey]; !ok {
		return 0
	}
	if len(ruleValue) == 0 {
		return 0
	}
	count, keyIndex := 0, map1773[ruleKey]
	for _, item := range items {
		if item[keyIndex] == ruleValue {
			count++
		}
	}
	return count
}
