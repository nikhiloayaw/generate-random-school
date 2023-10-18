package random

import (
	"school/types"
	"sync"
)

const (
	minHouseNum = 1
	maxHouseNum = 666
)

type RandomGenerator interface {
	GetStudents(count int) []types.Student
	getAddress() types.Address
	getAllSubjects() []types.Subject
	getSubjectScore() uint
	getGender() string
	isADisabilityPerson() bool
}

type randomGenerator struct {
	mu     sync.RWMutex
	states []types.State
	names  []string
}

func NewRandomGenerator(states []types.State, names []string) RandomGenerator {

	return &randomGenerator{
		mu:     sync.RWMutex{},
		states: states,
		names:  names,
	}
}
