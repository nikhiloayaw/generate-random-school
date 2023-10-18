package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"school/random"
	"school/school"
	"school/types"
	"sync"
	"time"

	"github.com/xuri/excelize/v2"
)

const (
	MinStudents = 30 // minimum students in a class
	MaxStudents = 40 // maximum students in a class
	MaxClasses  = 10 // maximum classes in a school
)

var (
	startWorkChan = make(chan struct{}, 2)
	OutDir        = "./out"

	NamesSheetFileName = "./inputs/names.xls"
	NamesSheetName     = "names"

	StatesSheetFileName = "./inputs/states.xlsx"
	StatesSheetName     = "states"
)

func ConvertToJSON(data any) ([]byte, error) {

	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data to json")
	}

	return b, nil
}

func ConvertToXML(data any) ([]byte, error) {

	b, err := xml.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data to xml")
	}

	return b, nil
}

func ConvertToExcel(school types.School) error {

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheetName := "school"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return fmt.Errorf("failed to create sheet: %w", err)
	}

	headers := []string{
		"School", "Class",
		"Name", "Age", "RollNumber", "Gender", "HaveDisability",

		"House No", "City", "State",

		"Subject1", "Score", "Grade", "Class Category", "Passed",
		"Subject2", "Score", "Grade", "Class Category", "Passed",
		"Subject3", "Score", "Grade", "Class Category", "Passed",
		"Subject4", "Score", "Grade", "Class Category", "Passed",
		"Subject5", "Score", "Grade", "Class Category", "Passed",
		"Subject6", "Score", "Grade", "Class Category", "Passed",
	}

	cellNames := make([]string, len(headers)) // to store cell names
	for i, header := range headers {

		// save cell names
		cellName, err := excelize.ColumnNumberToName(i + 1)
		if err != nil {
			return fmt.Errorf("failed to get column name")
		}

		cellNames[i] = cellName
		// set header values
		f.SetCellValue(sheetName, cellName+"1", header)
	}

	row := 2

	for _, class := range school.Classes {
		for _, student := range class.Students {

			// school and class
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[0], row), school.Name)
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[1], row), class.Name)

			// student detail
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[2], row), student.Name)
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[3], row), student.Age)
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[4], row), student.RollNumber)
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[5], row), student.Gender)
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[6], row), student.HaveDisability)

			// set addresses
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[7], row), student.Address.HouseNumber)
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[8], row), student.Address.City)
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[9], row), student.Address.State)
			// set score details
			for i, subject := range student.Scores {

				// each subject's details diff is 5
				increment := (i * 5)
				f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[10+increment], row), subject.Name)
				f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[11+increment], row), subject.Score)
				f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[12+increment], row), subject.Grade)
				f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[13+increment], row), subject.ClassCategory)
				f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[14+increment], row), subject.Passed)
			}

			row++
		}
	}

	f.SetActiveSheet(index)

	fileName := OutDir + "/" + school.Name + ".xlsx"

	if f.SaveAs(fileName); err != nil {
		return fmt.Errorf("failed to save excel file: %w", err)
	}

	return nil
}

func WriteToFile(name string, data []byte) error {

	file, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %w", err)
	}

	return nil
}

func GetAllNamesFromSheet(fileName, sheetName string) ([]string, error) {

	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %v", fileName, err)
	}

	defer f.Close()

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s sheet: %v", sheetName, err)
	}

	names := make([]string, len(rows)-1)

	for i := 1; i < len(rows); i++ {
		names[i-1] = rows[i][0]
	}

	return names, nil
}

func getAllStatesAndCitiesFromSheet(fileName, sheetName string) ([]types.State, error) {

	f, err := excelize.OpenFile(fileName)

	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %v", fileName, err)
	}
	defer f.Close()

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s sheet: %v", sheetName, err)
	}

	states := make([]types.State, 3)

	var city, state string

	states[0].Name = "Kerala"
	states[1].Name = "Tamil Nadu"
	states[2].Name = "Karnataka"

	for i := 1; i < len(rows); i++ {

		city = rows[i][0]  //get city
		state = rows[i][1] // get state

		switch state {
		case "Kerala":
			states[0].Cities = append(states[0].Cities, city)
		case "Tamil Nadu":
			states[1].Cities = append(states[1].Cities, city)
		case "Karnataka":
			states[2].Cities = append(states[2].Cities, city)
		}

		// check map of districts have districts exist or not
		// if _, ok := districts[district]; !ok {
		// 	districts[district] = []string{city}
		// } else {
		// 	districts[district] = append(districts[district], city)
		// }
	}

	return states, nil
}

func main() {

	var (
		wg         sync.WaitGroup
		schoolName string
	)

	// create the output directory
	if err := os.MkdirAll(OutDir, 0700); err != nil {
		log.Fatal("failed to create output dir: ", err)
	}

	// get all states details from sheet
	states, err := getAllStatesAndCitiesFromSheet(StatesSheetFileName, StatesSheetName)
	if err != nil {
		log.Fatal("failed to get all address from sheet: ", err)
	}

	// get all names from sheet
	names, err := GetAllNamesFromSheet(NamesSheetFileName, NamesSheetName)

	if err != nil {
		log.Fatal("failed to get all names: ", err)
	}

	// create random generator
	randomGen := random.NewRandomGenerator(states, names)

	// using random generator and school details create shool maker
	schoolMaker := school.NewSchoolMaker(MinStudents, MaxStudents, MaxClasses, randomGen)

	// get the school name from user
	fmt.Print("Enter the school name: ")
	fmt.Scanf("%s", &schoolName)

	start := time.Now()
	school := schoolMaker.CreateSchool(schoolName)

	// write the school data on json and excel  file
	wg.Add(2)

	go func() {
		defer wg.Done()

		data, err := ConvertToJSON(school)
		if err != nil {
			log.Println("failed to convert school data to json: ", err)
		}

		fileName := OutDir + "/" + schoolName + ".json"
		if err := WriteToFile(fileName, data); err != nil {
			log.Println("failed to write data to file as json: ", err)
		}
	}()

	go func() {
		defer wg.Done()

		err := ConvertToExcel(school)
		if err != nil {
			log.Println("failed to convert school data to excel: ", err)
		}
	}()

	wg.Wait()

	fmt.Println("total time taken method: ", time.Since(start))
}
