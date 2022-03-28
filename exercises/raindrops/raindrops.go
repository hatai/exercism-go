package raindrops

import (
	"fmt"
	"strings"
)

func Convert(input int) string {
	var builder strings.Builder
	if input % 3 == 0 {
		builder.WriteString("Pling")
	}

	if input % 5 == 0 {
		builder.WriteString("Plang")
	}

	if input % 7 == 0 {
		builder.WriteString("Plong")
	}

	if builder.Len() == 0 {
		_, _ = fmt.Fprint(&builder, input)
	}

	return builder.String()
}