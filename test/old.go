// package main

// import (
// 	"context"
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
// 	MinStudents  = 100  // minimum students in a class
// 	MaxStudents  = 100  // maximum students in a class
// 	TotalClasses = 1000 // total classes in a school
// )

// var (
// 	OutDir           = "./out"
// 	mu               sync.RWMutex
// 	namesCollection  []string
// 	nameChan         = make(chan string)
// 	nameCountReqChan = make(chan int)
// 	totalWorkers     = 5
// )

// func main() {

// 	start := time.Now()
// 	// setup the name collection
// 	setupTheNameCollection()

// 	if err := os.MkdirAll(OutDir, 0700); err != nil {
// 		log.Fatal("failed to create the output folder: ", err)
// 	}

// 	var schoolName string
// 	fmt.Print("Enter the school name: ")
// 	fmt.Scanf("%s", &schoolName)

// 	// create a context for workers to check the work is completed or not
// 	ctx, cancel := context.WithCancel(context.Background())

// 	// run workers
// 	RunWorkers(ctx)

// 	// start creating school
// 	school := CreateSchool(schoolName)
// 	fmt.Println("school created in: ", time.Since(start))
// 	// call cancel to announce the workers that work is completed
// 	cancel()

// 	var (
// 		wg sync.WaitGroup
// 	)

// 	wg.Add(2)

// 	go func() {
// 		defer wg.Done()

// 		data, err := ConvertToJSON(school)
// 		if err != nil {
// 			log.Println("failed to convert school data to json: ", err)
// 		}

// 		fileName := OutDir + "/" + schoolName + ".json"
// 		if err := WriteToFile(fileName, data); err != nil {
// 			log.Println("failed to write data to file as json: ", err)
// 		}
// 	}()

// 	go func() {
// 		defer wg.Done()

// 		data, err := ConvertToXML(school)
// 		if err != nil {
// 			log.Println("failed to convert school data to xml: ", err)
// 		}

// 		fileName := OutDir + "/" + schoolName + ".xml"
// 		if err := WriteToFile(fileName, data); err != nil {
// 			log.Println("failed to write data to file as xml: ", err)
// 		}

// 	}()

// 	wg.Wait()

// 	fmt.Println("total time for first method: ", time.Since(start))
// }

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

// func RunWorkers(ctx context.Context) {

// 	for i := 1; i <= totalWorkers; i++ {
// 		go ReadNames(ctx)
// 	}
// }

// func ReadNames(ctx context.Context) {

// 	// lock the mutex as read lock until the read operations complete
// 	mu.RLock()
// 	defer mu.RUnlock()

// 	var randIdx, requestCount int
// 	for {
// 		select {
// 		case <-ctx.Done(): // check the ctx done for completion
// 			fmt.Println("worker returning...")
// 			return
// 		case requestCount = <-nameCountReqChan: // listen on name count channel for request names

// 			fmt.Printf("got %d name requests sending names....\n", requestCount)

// 			for i := 1; i <= requestCount; i++ { // send n number of names to the nameChan

// 				// select a random index in between the slice
// 				randIdx = random.GetIntBetween(0, len(namesCollection)-1)
// 				nameChan <- namesCollection[randIdx]
// 			}
// 		}
// 	}
// }

// func CreateSchool(name string) types.School {

// 	classChan := make(chan types.Class, 4)
// 	// create classes
// 	for i := 1; i <= TotalClasses; i++ {
// 		className := fmt.Sprintf("class-%d", i)
// 		go MakeClass(className, classChan)
// 	}

// 	classes := make([]types.Class, TotalClasses)

// 	// save the classes to the slice
// 	for i := range classes {
// 		classes[i] = <-classChan
// 	}

// 	return types.School{
// 		Name:    name,
// 		Classes: classes,
// 	}
// }

// func MakeClass(name string, classChan chan<- types.Class) {

// 	students := MakeStudents()
// 	classChan <- types.Class{
// 		Name:          name,
// 		Students:      students,
// 		TotalStudents: uint(len(students)),
// 	}
// }

// func MakeStudents() []types.Student {

// 	studentsCount := random.GetIntBetween(MinStudents, MaxStudents)

// 	studentNames := make([]string, studentsCount)

// 	nameCountReqChan <- studentsCount // send that I need n number of names
// 	// fill the slice by workers sending names
// 	for i := range studentNames {
// 		studentNames[i] = <-nameChan
// 	}

// 	return random.GetStudentsByNames(studentNames)
// }

// func setupTheNameCollection() {
// 	var (
// 		fileName  = "names.xls"
// 		sheetName = "random"
// 	)

// 	f, err := excelize.OpenFile(fileName)

// 	if err != nil {
// 		log.Fatalf("failed to read file %s: %v", fileName, err)
// 	}

// 	rows, err := f.GetRows(sheetName)
// 	if err != nil {
// 		log.Fatalf("failed to get %s sheet: %v", sheetName, err)
// 	}

// 	// lock the mutex because writing to the slice
// 	mu.Lock()
// 	defer mu.Unlock()
// 	for i := 1; i < len(rows); i++ {
// 		namesCollection = append(namesCollection, rows[i][0])
// 	}

// 	return
// }
