package listops

type IntList []int
type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

func (s IntList) Foldl(fn binFunc, initial int) int {
	if len(s) == 0 {
		return initial
	}

	result := initial
	for _, v := range s {
		result = fn(result, v)
	}

	return result
}

func (s IntList) Foldr(fn binFunc, initial int) int {
	if len(s) == 0 {
		return initial
	}

	result := initial
	for _, v := range s {
		result = fn(v, result)
		if result == 0 {
			result = v
		}
	}

	return result
}

func (s IntList) Filter(fn predFunc) IntList {
	result := []int{}
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn unaryFunc) IntList {
	result := []int{}
	for _, v := range s {
		result = append(result, fn(v))
	}

	return result
}

func (s IntList) Reverse() IntList {
	result := []int{}
	for j := range s {
		index := len(s) - 1 - j
		result = append(result, s[index])
	}

	return result
}

func (s IntList) Append(x IntList) IntList {
	result := s
	for _, v := range x {
		result = append(result, v)
	}

	return result
}

func (s IntList) Concat(args []IntList) IntList {
	result := s
	for _, v := range args {
		for _, vv := range v {
			result = append(result, vv)
		}
	}

	return result
}

