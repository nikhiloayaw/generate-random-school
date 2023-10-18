// package main

// import (
// 	"encoding/json"
// 	"encoding/xml"
// 	"fmt"
// 	"log"
// 	"os"
// 	"school/random"
// 	"school/types"
// 	"sync"
// 	"time"

// 	"github.com/xuri/excelize/v2"
// )

// const (
// 	MinStudents  = 10000 // minimum students in a class
// 	MaxStudents  = 10000 // maximum students in a class
// 	TotalClasses = 10000 // total classes in a school
// )

// var (
// 	OutDir = "./out"
// )

// func ConvertToJSON(data any) ([]byte, error) {

// 	b, err := json.Marshal(data)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to marshal data to json")
// 	}

// 	return b, nil
// }

// func ConvertToXML(data any) ([]byte, error) {

// 	b, err := xml.Marshal(data)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to marshal data to xml")
// 	}

// 	return b, nil
// }

// func WriteToFile(name string, data []byte) error {

// 	file, err := os.Create(name)
// 	if err != nil {
// 		return fmt.Errorf("failed to create file: %w", err)
// 	}

// 	_, err = file.Write(data)
// 	if err != nil {
// 		return fmt.Errorf("failed to write data to file: %w", err)
// 	}

// 	return nil
// }

// func GetAllNames() ([]string, error) {

// 	var (
// 		fileName  = "names.xls"
// 		sheetName = "random"
// 	)

// 	f, err := excelize.OpenFile(fileName)

// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read file %s: %v", fileName, err)
// 	}

// 	rows, err := f.GetRows(sheetName)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get %s sheet: %v", sheetName, err)
// 	}

// 	names := make([]string, len(rows)-1)

// 	for i := 1; i < len(rows); i++ {
// 		names[i-1] = rows[i][0]
// 	}

// 	return names, nil
// }

// type SchoolMaker struct {
// 	mu            sync.RWMutex
// 	studentsNames []string
// }

// func NewSchoolMaker(names []string) SchoolMaker {

// 	return SchoolMaker{
// 		mu:            sync.RWMutex{},
// 		studentsNames: names,
// 	}
// }

// func (s *SchoolMaker) CreateSchool(schoolName string) types.School {

// 	var (
// 		// classChan = make(chan types.Class, 3)
// 		className string
// 		classes   = make([]types.Class, TotalClasses)
// 	)

// 	// receive all classes
// 	for i := 1; i <= TotalClasses; i++ {

// 		classes[i-1] = s.createClass(className)
// 	}

// 	return types.School{
// 		Name:    schoolName,
// 		Classes: classes,
// 	}
// }

// func (s *SchoolMaker) createClass(name string) types.Class {

// 	// select a random student count between min and max student count
// 	studentCount := random.GetIntBetween(MinStudents, MaxStudents)

// 	students := s.getStudents(studentCount)

// 	return types.Class{
// 		Name:          name,
// 		Students:      students,
// 		TotalStudents: uint(len(students)),
// 	}
// }

// func (s *SchoolMaker) getStudents(count int) []types.Student {

// 	// get random students
// 	return random.GetStudentsByNames(s.getNames(count))
// }

// func (s *SchoolMaker) getNames(count int) []string {

// 	names := make([]string, count)

// 	var (
// 		startIdx = 0
// 		endIdx   = len(s.studentsNames) - 1
// 	)

// 	for i := range names {
// 		// select a random name from the student names and add it
// 		names[i] = s.studentsNames[random.GetIntBetween(startIdx, endIdx)]
// 	}

// 	return names
// }

// func main() {

// 	var (
// 		wg         sync.WaitGroup
// 		schoolName string
// 	)

// 	// create the output directory
// 	if err := os.MkdirAll(OutDir, 0700); err != nil {
// 		log.Fatal("failed to create output dir: ", err)
// 	}

// 	names, err := GetAllNames()

// 	if err != nil {
// 		log.Fatal("failed to get all names: ", err)
// 	}

// 	start := time.Now()
// 	schoolMaker := NewSchoolMaker(names)

// 	fmt.Print("Enter the school name: ")
// 	fmt.Scanf("%s", &schoolName)

// 	school := schoolMaker.CreateSchool(schoolName)

// 	fmt.Println("school created in: ", time.Since(start))
// 	_ = school
// 	_ = wg
// 	// write the school data on json and xml  file
// 	// wg.Add(2)

// 	// go func() {
// 	// 	defer wg.Done()

// 	// 	data, err := ConvertToJSON(school)
// 	// 	if err != nil {
// 	// 		log.Println("failed to convert school data to json: ", err)
// 	// 	}

// 	// 	fileName := OutDir + "/" + schoolName + ".json"
// 	// 	if err := WriteToFile(fileName, data); err != nil {
// 	// 		log.Println("failed to write data to file as json: ", err)
// 	// 	}
// 	// }()

// 	// go func() {
// 	// 	defer wg.Done()

// 	// 	data, err := ConvertToXML(school)
// 	// 	if err != nil {
// 	// 		log.Println("failed to convert school data to xml: ", err)
// 	// 	}

// 	// 	fileName := OutDir + "/" + schoolName + ".xml"
// 	// 	if err := WriteToFile(fileName, data); err != nil {
// 	// 		log.Println("failed to write data to file as xml: ", err)
// 	// 	}

// 	// }()

// 	// wg.Wait()

// 	fmt.Println("total time taken method: ", time.Since(start))
// }
