package random

import (
	"math/rand"
)

var (
	// used only for read purpose so no need of mutex
	genders = [...]string{"male", "female"} // genders list
)

// To get a random gender from list of genders
func (r *randomGenerator) getGender() string {

	index := rand.Intn(len(genders))

	return genders[index]
}
