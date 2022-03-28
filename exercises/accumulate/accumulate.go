package accumulate

type converter func(string) string

func Accumulate(inputs []string, fn converter) []string {
	var result []string

	for _, input := range inputs {
		result = append(result, fn(input))
	}

	return result
}