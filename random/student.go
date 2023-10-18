package random

import (
	"math/rand"
	"school/types"
	"sort"
)

const (
	MinAge = 12
	MaxAge = 18
)

// To get random students by names
func GetStudentsByNames(names []string) []types.Student {

	// first sort the names
	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})

	// create slice of int for random indexes
	randIndexes := make([]int, len(names))

	// fill the index and make it randomized
	for i := range randIndexes {
		randIndexes[i] = i
		// swap this number with its any of previous number to make the slice randomize
		j := rand.Intn(i + 1)
		randIndexes[i], randIndexes[j] = randIndexes[j], randIndexes[i]
	}

	students := make([]types.Student, len(names))

	// in a class all students age should be almost same(like : 12,13,14)
	randomAge := GetIntBetween(MinAge, MaxAge)

	// new min and max age so in this call all students age will be in this new min and max age range
	newMinAge, newMaxAge := randomAge-1, randomAge+1

	for idx, randIdx := range randIndexes {

		student := types.Student{
			Name:       names[randIdx],
			Age:        uint(GetIntBetween(newMinAge, newMaxAge)),
			RollNumber: uint(randIdx) + 1,
			Gender:     GetGender(),
			Scores:     GetAllSubjects(),
		}

		students[idx] = student
	}

	return students
}
