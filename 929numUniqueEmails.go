package main

import (
	"strings"
)

func main() {
	println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
	println(numUniqueEmails([]string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}))

}
func numUniqueEmails(emails []string) int {
	m := make(map[string]struct{}, len(emails))
	for _, email := range emails {
		e := strings.Split(email, "@")
		if len(e) < 2 {
			return 0
		}
		e[0] = strings.ReplaceAll(e[0], ".", "")
		if strings.ContainsAny(e[0], "+") {
			e[0] = strings.Split(e[0], "+")[0]
		}
		builder := strings.Builder{}
		builder.WriteString(e[0])
		builder.WriteString("@")
		builder.WriteString(e[1])

		if _, ok := m[builder.String()]; !ok {
			m[builder.String()] = struct{}{}
		}
	}
	return len(m)
}
