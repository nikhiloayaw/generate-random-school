package school

import (
	"fmt"
	"school/random"
	"school/types"
)

// school maker functionalities
type SchoolMaker interface {
	CreateSchool(schoolName string) types.School
}

type schoolMaker struct {
	maxClasses          int
	minStudentsPerClass int
	maxStudentsPerClass int

	startWorkChan chan struct{}
	randGenerator random.RandomGenerator
}

func NewSchoolMaker(minStudentPerClass, maxStudentPerClass,
	maxClasses int, randGenerator random.RandomGenerator) SchoolMaker {

	return &schoolMaker{
		maxClasses:          maxClasses,
		minStudentsPerClass: minStudentPerClass,
		maxStudentsPerClass: maxStudentPerClass,

		startWorkChan: make(chan struct{}, 4),
		randGenerator: randGenerator,
	}
}

func (s *schoolMaker) CreateSchool(schoolName string) types.School {

	totalClasses := random.GetIntBetween(1, s.maxClasses)

	var (
		classChan = make(chan types.Class, 3)
		className string
		classes   = make([]types.Class, totalClasses)
	)

	go func() {
		// fire create class in separate goroutines to do concurrent
		for i := 1; i <= totalClasses; i++ {

			// limit the active goroutines
			s.startWorkChan <- struct{}{} // if the buffered channel it's full then the wait for others to complete their work and then fire new goroutine

			className = fmt.Sprintf("class-%d", i)
			go s.createClass(className, classChan)
		}

	}()

	// receive all classes
	for i := 1; i <= totalClasses; i++ {

		classes[i-1] = <-classChan
		// release  values from channel whenever goroutines work completed by sending class and receive on here
		<-s.startWorkChan
	}

	return types.School{
		Name:    schoolName,
		Classes: classes,
	}
}
