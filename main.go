package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
)

type DataFrame struct {
	patientID []uint16
	//demographic details
	age       []uint8 // 60 - 90 years
	gender    []bool  // tru = female, false = male
	ethnicity []uint8 // 0: caucasian, 1:afroamerican, 2:asian, 3: other
	education []uint8 // 0:none, 1:high school, 2: belchelor's 3: higher
	// lifestyle factors
	BMI                []float64 // 15 -40
	smoking            []bool
	alcoholConsumption []float64 // weekly alcohol consumption
	physicalActivity   []float64 // weekly physical activity 0-10
	dietQuality        []float64 // 0 - 10	(grade)
	sleepQuality       []float64 // 4 - 10
	//medical history
	familyAlzheimers []bool // alzhaimers in family history
	cardioDisease    []bool
	diabetes         []bool
	depression       []bool
	headInjury       []bool
	hypertension     []bool
	//clinical measurements
	systolicBP       []uint
	diastolicBP      []uint
	cholesterolTotal []float64
	cholesterolLDL   []float64
	cholesterolHDL   []float64
	triglycerides    []float64
	//cognitive and functional assessments
	MMSE             []float64 //mini-mental state examination score 0 - 30
	funcAssessment   []float64 // 0-10
	memoryComplaints []bool
	behavioralProb   []bool
	ADL              []float64 //activities of daily living score 0 - 10
	//symptoms
	confusion          []bool
	disorientation     []bool
	personalityChanges []bool
	difficultyTask     []bool
	forgetfulness      []bool
	//DIAGNOSIS
	diagnosis []bool
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

		df.patientID = append(df.patientID, parseUint16(record[0]))
		df.age = append(df.age, parseUint8(record[1]))
		df.gender = append(df.gender, parseBool(record[2]))
		df.ethnicity = append(df.ethnicity, parseUint8(record[3]))
		df.education = append(df.education, parseUint8(record[4]))
		df.BMI = append(df.BMI, parseFloat64(record[5]))
		df.smoking = append(df.smoking, parseBool(record[6]))
		df.alcoholConsumption = append(df.alcoholConsumption, parseFloat64(record[7]))
		df.physicalActivity = append(df.physicalActivity, parseFloat64(record[8]))
		df.dietQuality = append(df.dietQuality, parseFloat64(record[9]))
		df.sleepQuality = append(df.sleepQuality, parseFloat64(record[10]))
		df.familyAlzheimers = append(df.familyAlzheimers, parseBool(record[11]))
		df.cardioDisease = append(df.cardioDisease, parseBool(record[12]))
		df.diabetes = append(df.diabetes, parseBool(record[13]))
		df.depression = append(df.depression, parseBool(record[14]))
		df.headInjury = append(df.headInjury, parseBool(record[15]))
		df.hypertension = append(df.hypertension, parseBool(record[16]))
		df.systolicBP = append(df.systolicBP, parseUint(record[17]))
		df.diastolicBP = append(df.diastolicBP, parseUint(record[18]))
		df.cholesterolTotal = append(df.cholesterolTotal, parseFloat64(record[19]))
		df.cholesterolLDL = append(df.cholesterolLDL, parseFloat64(record[20]))
		df.cholesterolHDL = append(df.cholesterolHDL, parseFloat64(record[21]))
		df.triglycerides = append(df.triglycerides, parseFloat64(record[22]))
		df.MMSE = append(df.MMSE, parseFloat64(record[23]))
		df.funcAssessment = append(df.funcAssessment, parseFloat64(record[24]))
		df.memoryComplaints = append(df.memoryComplaints, parseBool(record[25]))
		df.behavioralProb = append(df.behavioralProb, parseBool(record[26]))
		df.ADL = append(df.ADL, parseFloat64(record[27]))
		df.confusion = append(df.confusion, parseBool(record[28]))
		df.disorientation = append(df.disorientation, parseBool(record[29]))
		df.personalityChanges = append(df.personalityChanges, parseBool(record[30]))
		df.difficultyTask = append(df.difficultyTask, parseBool(record[31]))
		df.forgetfulness = append(df.forgetfulness, parseBool(record[32]))
		df.diagnosis = append(df.diagnosis, parseBool(record[33]))
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

func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func sigmoidDerivative(x float64) float64 {
	return x * (1.0 - x)
}

type NeuralNetwork struct {
	inputSize  int
	hiddenSize int
	outputSize int

	weightsInputHidden  [][]float64
	weightsHiddenOutput []float64
	hiddenLayerInput    []float64
	hiddenLayerOutput   []float64

	outputLayerInput float64
	output           float64
}

func newNeuralNetwork(inputSize, hiddenSize, outputSize int) NeuralNetwork {
	tempWeightsInputHidden := make([][]float64, inputSize)
	for i := range tempWeightsInputHidden {
		tempWeightsInputHidden[i] = make([]float64, hiddenSize)
	}
	tempWeightsHiddenOutput := make([]float64, hiddenSize)

	for i := range tempWeightsInputHidden {
		for j := range tempWeightsInputHidden[i] {
			tempWeightsInputHidden[i][j] = rand.Float64()
		}
	}

	for i := range tempWeightsHiddenOutput {
		tempWeightsHiddenOutput[i] = rand.Float64()
	}

	return NeuralNetwork{hiddenSize: hiddenSize,
		inputSize:           inputSize,
		outputSize:          outputSize,
		weightsInputHidden:  tempWeightsInputHidden,
		weightsHiddenOutput: tempWeightsHiddenOutput,
		outputLayerInput:    0,
		output:              0,
	}
}

func (nn *NeuralNetwork) forward(input []float64) float64 {
	nn.hiddenLayerInput = make([]float64, nn.hiddenSize)
	nn.hiddenLayerOutput = make([]float64, nn.hiddenSize)
	for i := 0; i < nn.hiddenSize; i++ {
		nn.hiddenLayerInput[i] = 0.0
		for j := 0; j < nn.inputSize; j++ {
			nn.hiddenLayerInput[i] += input[j] * nn.weightsInputHidden[j][i]
		}
		nn.hiddenLayerOutput[i] = sigmoid(nn.hiddenLayerInput[i])
	}
	nn.outputLayerInput = 0.0
	for i := 0; i < nn.hiddenSize; i++ {
		nn.outputLayerInput += nn.hiddenLayerOutput[i] * nn.weightsHiddenOutput[i]
	}

	nn.output = sigmoid(nn.outputLayerInput)

	return nn.output
}

func (nn NeuralNetwork) backward(input []float64, target, learningRate float64) {
	var outputError float64 = target - nn.output
	var outputDelta float64 = outputError * sigmoidDerivative(nn.output)

	hiddenError := make([]float64, nn.hiddenSize)
	hiddenDelta := make([]float64, nn.hiddenSize)

	for i := 0; i < nn.hiddenSize; i++ {
		hiddenError[i] = outputDelta * nn.weightsHiddenOutput[i]
		hiddenDelta[i] = hiddenError[i] * sigmoidDerivative(nn.hiddenLayerOutput[i])
	}
	for i := 0; i < nn.hiddenSize; i++ {
		nn.weightsHiddenOutput[i] += nn.hiddenLayerOutput[i] * outputDelta * learningRate
	}

	for i := 0; i < nn.hiddenSize; i++ {
		for j := 0; j < nn.hiddenSize; j++ {
			nn.weightsInputHidden[i][j] += input[i] * hiddenDelta[i] * learningRate
		}
	}
}

func boolToFloat(x bool) float64 {
	if x {
		return 1.0
	} else {
		return 0.0
	}
}

func (nn NeuralNetwork) loss(target float64) float64 {
	loss := math.Pow(target-nn.output, 2)
	return loss
}

func averageLoss(lossTable []float64) float64 {
	var result float64
	for i := range lossTable {
		result += lossTable[i]
	}
	return result / float64(len(lossTable))
}

func main() {
	inputSize := 31
	hiddenSize := 31
	outputSize := 1
	learningRate := 0.1
	epochs := 10
	nn := newNeuralNetwork(inputSize, hiddenSize, outputSize)

	var data DataFrame
	data.loadData("alzheimers_disease_data.csv")
	for j := 0; j < epochs; j++ {
		var lossTable []float64
		for i := range data.patientID {
			var d []float64
			d = append(d, float64(data.age[i]))
			d = append(d, boolToFloat(data.gender[i]))
			d = append(d, float64(data.ethnicity[i]))
			d = append(d, float64(data.education[i]))
			d = append(d, data.BMI[i])
			d = append(d, boolToFloat(data.smoking[i]))
			d = append(d, data.alcoholConsumption[i])
			d = append(d, data.physicalActivity[i])
			d = append(d, data.dietQuality[i])
			d = append(d, data.sleepQuality[i])
			d = append(d, boolToFloat(data.familyAlzheimers[i]))
			d = append(d, boolToFloat(data.cardioDisease[i]))
			d = append(d, boolToFloat(data.diabetes[i]))
			d = append(d, boolToFloat(data.headInjury[i]))
			d = append(d, boolToFloat(data.hypertension[i]))
			d = append(d, float64(data.systolicBP[i]))
			d = append(d, float64(data.diastolicBP[i]))
			d = append(d, data.cholesterolTotal[i])
			d = append(d, data.cholesterolLDL[i])
			d = append(d, data.cholesterolHDL[i])
			d = append(d, data.triglycerides[i])
			d = append(d, data.MMSE[i])
			d = append(d, data.funcAssessment[i])
			d = append(d, boolToFloat(data.memoryComplaints[i]))
			d = append(d, boolToFloat(data.behavioralProb[i]))
			d = append(d, data.ADL[i])
			d = append(d, boolToFloat(data.confusion[i]))
			d = append(d, boolToFloat(data.disorientation[i]))
			d = append(d, boolToFloat(data.personalityChanges[i]))
			d = append(d, boolToFloat(data.difficultyTask[i]))
			d = append(d, boolToFloat(data.forgetfulness[i]))
			nn.forward(d)
			nn.backward(d, boolToFloat(data.diagnosis[i]), learningRate)
			target := boolToFloat(data.diagnosis[i])
			loss := nn.loss(target)
			//fmt.Println(loss)
			nn.backward(d, target, learningRate)
			lossTable = append(lossTable, loss)
		}
		if j%1 == 0 {
			fmt.Printf("Epoch: %d Loss: %f \n", j, averageLoss(lossTable))
		}
	}
}
