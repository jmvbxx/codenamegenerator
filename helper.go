package codenamegenerator

import (
	"fmt"
	"strings"
)

func getSingleWord(s string) string {
	if strings.Contains(s, " ") || strings.Contains(s, "-") {
		s1 := strings.ReplaceAll(s, " ", "-")
		s2 := strings.Split(s1, "-")
		string := fmt.Sprintf(s2[len(s2)-1])
		return string
	} else {
		return s
	}
}
