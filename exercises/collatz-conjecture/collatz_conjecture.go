package collatzconjecture

type myError struct {
	What string
}

func (m myError) Error() string {
	return m.What
}

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return -1, myError{"N should be n > 0"}
	}

	count := 0
	for n != 1 {
		if n % 2 == 0 {
			n = n / 2
		} else {
			n = n * 3 + 1
		}

		count++
	}

	return count, nil
}