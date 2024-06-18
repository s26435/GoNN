package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Data struct {
	patientID          uint16
	age                uint8
	gender             bool
	ethnicity          uint8
	education          uint8
	BMI                float64
	smoking            bool
	alcoholConsumption float64
	physicalActivity   float64
	dietQuality        float64
	sleepQuality       float64
	familyAlzheimers   bool
	cardioDisease      bool
	diabetes           bool
	depression         bool
	headInjury         bool
	hypertension       bool
	systolicBP         uint
	diastolicBP        uint
	cholesterolTotal   float64
	cholesterolLDL     float64
	cholesterolHDL     float64
	triglycerides      float64
	MMSE               float64
	funcAssessment     float64
	memoryComplaints   bool
	behavioralProb     bool
	ADL                float64
	confusion          bool
	disorientation     bool
	personalityChanges bool
	difficultyTask     bool
	forgetfulness      bool
	diagnosis          bool
}

type DataFrame struct {
	frame []Data
}

func (data Data) Print() {
	fmt.Printf("Patient ID: %d\n", data.patientID)
	fmt.Printf("Age: %d\n", data.age)
	fmt.Printf("Gender: %t\n", data.gender)
	fmt.Printf("Ethnicity: %d\n", data.ethnicity)
	fmt.Printf("Education: %d\n", data.education)
	fmt.Printf("BMI: %.2f\n", data.BMI)
	fmt.Printf("Smoking: %t\n", data.smoking)
	fmt.Printf("Alcohol Consumption: %.2f\n", data.alcoholConsumption)
	fmt.Printf("Physical Activity: %.2f\n", data.physicalActivity)
	fmt.Printf("Diet Quality: %.2f\n", data.dietQuality)
	fmt.Printf("Sleep Quality: %.2f\n", data.sleepQuality)
	fmt.Printf("Family Alzheimer's: %t\n", data.familyAlzheimers)
	fmt.Printf("Cardio Disease: %t\n", data.cardioDisease)
	fmt.Printf("Diabetes: %t\n", data.diabetes)
	fmt.Printf("Depression: %t\n", data.depression)
	fmt.Printf("Head Injury: %t\n", data.headInjury)
	fmt.Printf("Hypertension: %t\n", data.hypertension)
	fmt.Printf("Systolic BP: %d\n", data.systolicBP)
	fmt.Printf("Diastolic BP: %d\n", data.diastolicBP)
	fmt.Printf("Total Cholesterol: %.2f\n", data.cholesterolTotal)
	fmt.Printf("LDL Cholesterol: %.2f\n", data.cholesterolLDL)
	fmt.Printf("HDL Cholesterol: %.2f\n", data.cholesterolHDL)
	fmt.Printf("Triglycerides: %.2f\n", data.triglycerides)
	fmt.Printf("MMSE: %.2f\n", data.MMSE)
	fmt.Printf("Functional Assessment: %.2f\n", data.funcAssessment)
	fmt.Printf("Memory Complaints: %t\n", data.memoryComplaints)
	fmt.Printf("Behavioral Problems: %t\n", data.behavioralProb)
	fmt.Printf("ADL: %.2f\n", data.ADL)
	fmt.Printf("Confusion: %t\n", data.confusion)
	fmt.Printf("Disorientation: %t\n", data.disorientation)
	fmt.Printf("Personality Changes: %t\n", data.personalityChanges)
	fmt.Printf("Difficulty with Tasks: %t\n", data.difficultyTask)
	fmt.Printf("Forgetfulness: %t\n", data.forgetfulness)
	fmt.Printf("Diagnosis: %t\n", data.diagnosis)
}

func (df *DataFrame) loadData(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	reader := csv.NewReader(bufio.NewReader(file))

	// Skip the header row
	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		if len(record) != 35 {
			log.Fatalf("Expected 35 fields, but got %d", len(record))
		}

		data := Data{
			patientID:          parseUint16(record[0]),
			age:                parseUint8(record[1]),
			gender:             parseBool(record[2]),
			ethnicity:          parseUint8(record[3]),
			education:          parseUint8(record[4]),
			BMI:                parseFloat64(record[5]),
			smoking:            parseBool(record[6]),
			alcoholConsumption: parseFloat64(record[7]),
			physicalActivity:   parseFloat64(record[8]),
			dietQuality:        parseFloat64(record[9]),
			sleepQuality:       parseFloat64(record[10]),
			familyAlzheimers:   parseBool(record[11]),
			cardioDisease:      parseBool(record[12]),
			diabetes:           parseBool(record[13]),
			depression:         parseBool(record[14]),
			headInjury:         parseBool(record[15]),
			hypertension:       parseBool(record[16]),
			systolicBP:         parseUint(record[17]),
			diastolicBP:        parseUint(record[18]),
			cholesterolTotal:   parseFloat64(record[19]),
			cholesterolLDL:     parseFloat64(record[20]),
			cholesterolHDL:     parseFloat64(record[21]),
			triglycerides:      parseFloat64(record[22]),
			MMSE:               parseFloat64(record[23]),
			funcAssessment:     parseFloat64(record[24]),
			memoryComplaints:   parseBool(record[25]),
			behavioralProb:     parseBool(record[26]),
			ADL:                parseFloat64(record[27]),
			confusion:          parseBool(record[28]),
			disorientation:     parseBool(record[29]),
			personalityChanges: parseBool(record[30]),
			difficultyTask:     parseBool(record[31]),
			forgetfulness:      parseBool(record[32]),
			diagnosis:          parseBool(record[33]),
		}

		df.frame = append(df.frame, data)
	}
}

func parseUint16(s string) uint16 {
	value, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		log.Fatalf("Error parsing uint8: %s", err)
	}
	return uint16(value)
}

func parseUint8(s string) uint8 {
	v, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		log.Fatalf("Error parsing uint8: %s", err)
	}
	return uint8(v)
}

func parseUint(s string) uint {
	v, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		log.Fatalf("Error parsing uint: %s", err)
	}
	return uint(v)
}

func parseFloat64(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalf("Error parsing float64: %s", err)
	}
	return v
}

func parseBool(s string) bool {
	v, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatalf("Error parsing bool: %s", err)
	}
	return v
}

func (df DataFrame) Print() {
	for _, data := range df.frame {
		data.Print()
		fmt.Println("-------------")
	}
}

func main() {
	dataFrame := DataFrame{}
	dataFrame.loadData("alzheimers_disease_data.csv")
	dataFrame.Print()
}
