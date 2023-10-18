package school

import (
	"fmt"
	"school/random"
	"school/types"
)

func (s *schoolMaker) createClass(name string, classChan chan<- types.Class) {

	fmt.Println("create class : ", name, "started....")

	// select a random student count between min and max student count
	studentCount := random.GetIntBetween(s.minStudentsPerClass, s.maxStudentsPerClass)

	// students := s.getStudents(studentCount)
	students := s.getStudentsNew(studentCount)

	fmt.Println("create class : ", name, "completed....")

	classChan <- types.Class{
		Name:          name,
		Students:      students,
		TotalStudents: uint(len(students)),
	}
}
