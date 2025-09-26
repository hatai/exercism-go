package chance

import (
	"math/rand"
	"time"
)

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(19) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Float64() * 12.0
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
var animals = []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

func ShuffleAnimals() []string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	animalsCopy := make([]string, len(animals))
	copy(animalsCopy, animals)
	rand.Shuffle(len(animals), func(i, j int) {
		animalsCopy[i], animalsCopy[j] = animalsCopy[j], animalsCopy[i]
	})
	return animalsCopy
}
