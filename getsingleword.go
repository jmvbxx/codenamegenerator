package codenamegenerator

import (
	"fmt"
	"strings"
)

// getSingleWord function returns a single word in cases where the original
// response is multi-word either separated by a space or a hyphen
func getSingleWord(s string) string {
	s1 := strings.ReplaceAll(s, " ", "-")
	s2 := strings.Split(s1, "-")
	string := fmt.Sprintf(s2[len(s2)-1])
	return string
}
