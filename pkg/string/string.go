package string

import "strings"

func Reverse(s string) string {
	var builder strings.Builder
	runes := []rune(s)
	for i := len(runes) - 1; 0 <= i; i-- {
		builder.WriteRune(runes[i])
	}
	return builder.String()
}
