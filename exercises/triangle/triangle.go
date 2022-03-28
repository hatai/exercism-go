package triangle

import (
	"sort"
	"math"
)

type Kind string

const (
	Equ Kind = "Equ"
	Iso Kind = "Iso"
	Sca Kind = "Sca"
	NaT Kind = "NaT"
)

func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	sort.Float64s(sides)

	for _, v := range sides {
		if math.IsNaN(v) || math.IsInf(v, -1) || math.IsInf(v, 0) || v <= 0 {
			return NaT
		}
	}

	if sides[0]+sides[1] < sides[2] {
		return NaT
	}

	if sides[0] == sides[1] && sides[0] == sides[2] {
		return Equ
	}

	if sides[0] == sides[1] || sides[1] == sides[2] {
		return Iso
	}

	return Sca
}
