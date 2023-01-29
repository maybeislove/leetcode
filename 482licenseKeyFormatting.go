package main

import (
	"strings"
)

func main() {
	//println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
	//println(licenseKeyFormatting("2-5g-3-J", 2))
	println(licenseKeyFormatting("2-4A0r7-4k", 4))

}

func licenseKeyFormatting(s string, k int) string {
	result := ""
	splits := strings.Split(s, "-")
	for i := 0; i < len(splits); i++ {
		result += splits[i]
	}

	if len(result) == 0 {
		return ""
	}
	for i := len(result) - k; i > 0; i -= k {
		result = result[0:i] + "-" + result[i:]
	}

	return strings.ToUpper(result)

}
func licenseKeyFormatting2(s string, k int) string {
	result := strings.ReplaceAll(s, "-", "")
	if len(result) == 0 {
		return ""
	}
	for i := len(result) - k; i > 0; i -= k {
		result = result[0:i] + "-" + result[i:]
	}

	return strings.ToUpper(result)

}
