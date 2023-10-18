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

func MakeStudents(nameChan <-chan string, studentChan chan<- []types.Student) {

	var students []types.Student

	// in a class all students age should be almost same(like : 12,13,14)
	randomAge := GetIntBetween(MinAge, MaxAge)

	// new min and max age so in this call all students age will be in this new min and max age range
	newMinAge, newMaxAge := randomAge-1, randomAge+1

	// read until the name channel close
	for name := range nameChan {

		student := types.Student{
			Name:   name,
			Age:    uint(GetIntBetween(newMinAge, newMaxAge)),
			Gender: GetGender(),
			Scores: GetAllSubjects(),
		}

		students = append(students, student)
	}

	// update the students role number
	updateRoleNumber(students)

	studentChan <- students
}

func updateRoleNumber(students []types.Student) {

	// call sort helper to sort and the same time set the appropriate roll number
	sortAndUpdateRoleNumber(students, 0, len(students)-1)

	// then shuffle the slice again

	for i := range students {
		// select a random index from the previous the slice and swap with it
		j := GetIntBetween(0, i)
		students[i], students[j] = students[j], students[i]
	}
}

func sortAndUpdateRoleNumber(arr []types.Student, start, end int) {
	if start < end {
		pivotIdx := partition(arr, start, end)
		arr[pivotIdx].RollNumber = uint(pivotIdx) + 1
		sortAndUpdateRoleNumber(arr, start, pivotIdx-1)
		sortAndUpdateRoleNumber(arr, pivotIdx+1, end)
	}

	if start == end {
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
