package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}

	var results []string
	for i := range rhyme {
		if i == len(rhyme)-1 {
			results = append(results, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
		} else {
			results = append(results, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		}
	}

	return results
}
