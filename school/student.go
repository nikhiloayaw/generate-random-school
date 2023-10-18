package school

import (
	"school/types"
)

// new way of getting names
func (s *schoolMaker) getStudentsNew(count int) []types.Student {

	// get the random student lists and return it
	return s.randGenerator.GetStudents(count)
}
