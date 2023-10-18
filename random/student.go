package random

import (
	"school/types"
)

const (
	MinAge = 12
	MaxAge = 18
)

func (r *randomGenerator) GetStudents(count int) []types.Student {

	var students []types.Student

	// in a class all students age should be almost same(like : 12,13,14)
	randomAge := GetIntBetween(MinAge, MaxAge)

	// new min and max age so in this call all students age will be in this new min and max age range
	newMinAge, newMaxAge := randomAge-1, randomAge+1

	// read names and create add it until the name channel close
	for i := 1; i <= count; i++ {

		student := types.Student{
			Name:           r.getName(),
			Age:            uint(GetIntBetween(newMinAge, newMaxAge)),
			Gender:         r.getGender(),
			Scores:         r.getAllSubjects(),
			HaveDisability: r.isADisabilityPerson(),
			Address:        r.getAddress(),
		}

		students = append(students, student)
	}

	// update the students role number according to name
	updateRoleNumber(students)

	return students
}

func updateRoleNumber(students []types.Student) {

	// sort and update the role number according to the sorted name
	sortAndUpdateRoleNumber(students, 0, len(students)-1)

	//shuffle the sorted students
	for i := range students {
		// select a random index from the previous the slice and swap with it
		j := GetIntBetween(0, i)
		students[i], students[j] = students[j], students[i]
	}
}

// using quick sort for sorting; reason each time finding pivot it's actually same for roll number
func sortAndUpdateRoleNumber(arr []types.Student, start, end int) {

	if start < end {
		pivotIdx := partition(arr, start, end)
		//  founded pivot index + 1 is the name's roll number
		arr[pivotIdx].RollNumber = uint(pivotIdx) + 1
		sortAndUpdateRoleNumber(arr, start, pivotIdx-1)
		sortAndUpdateRoleNumber(arr, pivotIdx+1, end)
	}

	if start == end {
		// if start and end is same; the roll number start or end + 1
		arr[end].RollNumber = uint(end) + 1
	}
}

func partition(arr []types.Student, start, end int) int {
	pivot := arr[end].Name
	pIndex := start

	for i := start; i < end; i++ {
		if arr[i].Name < pivot {
			arr[i], arr[pIndex] = arr[pIndex], arr[i]
			pIndex++
		}
	}

	arr[pIndex], arr[end] = arr[end], arr[pIndex]
	return pIndex
}
