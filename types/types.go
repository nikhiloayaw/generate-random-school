package types

import (
	"errors"
)

var (
	ErrInvalidScore = errors.New("invalid score")
)

type School struct {
	Name    string
	Classes []Class
}

type Class struct {
	Name          string    // A to J; because only max of 10 class
	Students      []Student // 30 to 40
	TotalStudents uint
}

type Student struct {
	Name       string
	Age        uint
	RollNumber uint // ascending order based on name
	Gender     string
	Scores     []Subject
}

type Subject struct {
	Name          string
	Score         uint
	Grade         string
	ClassCategory string
	Passed        bool
}

// To set score to the subject and update score related fields
func (s *Subject) SetScore(score uint) error {

	// check score is valid or not
	if score < 0 || score > 100 {
		return ErrInvalidScore
	}

	// update the score on subject
	s.Score = score

	// update other details according to score
	if score >= 80 {
		s.Grade = "O"
		s.Passed = true
		s.ClassCategory = "Distinction"
	} else if score >= 70 {
		s.Grade = "A"
		s.Passed = true
		s.ClassCategory = "Distinction"
	} else if score >= 60 {
		s.Grade = "B"
		s.Passed = true
		s.ClassCategory = "First Class"
	} else if score >= 55 {
		s.Grade = "C"
		s.Passed = true
		s.ClassCategory = "Second Class"
	} else if score >= 50 {
		s.Grade = "D"
		s.Passed = true
		s.ClassCategory = "Second Class"
	} else if score >= 40 {
		s.Grade = "E"
		s.Passed = true
		s.ClassCategory = "Pass Class"
	} else {
		s.Grade = "F"
		s.Passed = false
		s.ClassCategory = "Fail"
	}

	return nil
}
