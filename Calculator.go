package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var Calculator CalculationData = InputCalculation()
	fmt.Println("first :", Calculator.firstNum)
	fmt.Println("second :", Calculator.secondNum)
	fmt.Println("operation :", Calculator.operation)
	fmt.Println("RS :", Calculator.IsRomanSystem)
}

type CalculationData struct {
	firstNum      int
	secondNum     int
	operation     string
	IsRomanSystem bool
}

func InputCalculation() CalculationData {
	var cal CalculationData
	var str string

	calscanner := bufio.NewScanner(os.Stdin)
	calscanner.Scan()
	str = calscanner.Text()

	cal.operation = OperationSearch(str)

	strNumbers := SeparationStrNumbers(str, cal.operation)

	if StringIsNumber(strNumbers[0]) && StringIsNumber(strNumbers[1]) {
		cal.firstNum, _ = strconv.Atoi(strNumbers[0])
		cal.secondNum, _ = strconv.Atoi(strNumbers[1])
		cal.IsRomanSystem = false

	} else if StringIsRoman(strNumbers[0]) && StringIsRoman(strNumbers[1]) {
		cal.firstNum = RomeNumbers[strNumbers[0]]
		cal.secondNum = RomeNumbers[strNumbers[1]]
		cal.IsRomanSystem = true
	} else {
		panic("Система счисления не определена!	")
	}

	return cal
}

func OperationSearch(str string) string {

	var flag byte = 0

	if strings.Count(str, "+") == 1 {
		flag ^= 1
	}
	if strings.Count(str, "-") == 1 {
		flag ^= 2
	}
	if strings.Count(str, "*") == 1 {
		flag ^= 4
	}
	if strings.Count(str, "/") == 1 {
		flag ^= 8
	}

	switch flag {
	case 1:
		return "+"
	case 2:
		return "-"
	case 4:
		return "*"
	case 8:
		return "/"
	default:
		panic("Формат математической операции не удовлетворяет заданию!")
	}
}

func SeparationStrNumbers(str string, operator string) []string {
	return strings.Fields(strings.Replace(str, operator, " ", 1))
}

func StringIsNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func StringIsRoman(str string) bool {
	return RomeNumbers[str] > 0
}

var RomeNumbers = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}
