package random

import (
	"math/rand"
)

func GetIntBetween(start, end int) int {

	max := (end - start) + 1
	
	return start + (rand.Intn(max))
}
