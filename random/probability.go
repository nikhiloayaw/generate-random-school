package random

import "math/rand"

func (r *randomGenerator)isADisabilityPerson() bool {

	probability := 0.2

	return rand.Float64() < probability
}
