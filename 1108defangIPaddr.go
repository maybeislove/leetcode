package main

import "strings"

func main() {

}
func defangIPaddr(address string) string {

	return strings.ReplaceAll(address, ".", "[.]")
}
