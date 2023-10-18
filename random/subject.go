package random

import (
	"math/rand"
	"school/types"
)

const (
	MaxScore = 100
)

var (
	// used only for read purpose so no need of mutex
	subjectNames = [...]string{"English", "Hindi", "Malayalam", "Maths", "Social", "Science"}
)

// To get a random score between 1 and maxScore
func getSubjectScore() uint {

	score := rand.Intn(MaxScore) + 1 // Intn func return (0 to maxMark-1); so adding 1 to it

	return uint(score)
}

// To get all subjects with random scores
func GetAllSubjects() []types.Subject {

	subjects := make([]types.Subject, len(subjectNames))

	for i := range subjectNames {

		subject := types.Subject{
			Name: subjectNames[i],
		}
		subject.SetScore(getSubjectScore())

		subjects[i] = subject
	}

	return subjects
}
