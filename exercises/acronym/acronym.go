package acronym

import (
	"strings"
	"regexp"
)

func Abbreviate(s string) string {
	rep := regexp.MustCompile("[-_]")
	s = rep.ReplaceAllString(s, " ")

	result := ""
	for _, word := range strings.Split(s, " ") {
		if word != "" {
			result += strings.ToUpper(string(word[0]))
		}
	}

	return result
}
