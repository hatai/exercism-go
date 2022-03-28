package hamming

type MyError struct {
	message string
}

func (e *MyError) Error() string {
	return e.message
}

func Distance(a, b string) (int, error) {
	count := 0

	if len(a) != len(b) {
		return count, &MyError{"strand length is not equal"}
	}

	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
