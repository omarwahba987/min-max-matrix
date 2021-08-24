package main

import "fmt"

type NumValue struct {
	Value  int
	Index  int
	Row    int
	Column int
}

func GetMinMax(arr []int, columns int) (min, max NumValue) {
	min.Value, max.Value = arr[0], arr[0]
	min.Index, max.Index = 0, 0
	if len(arr) == 1 {
		min.Row, min.Column = 1, 1
		max.Row, max.Column = 1, 1
		return min, max
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] > max.Value {
			max.Value = arr[i]
			max.Index = i
		}
		if arr[i] < min.Value {
			min.Value = arr[i]
			min.Index = i
		}
	}
	min.Row = (min.Index / columns) + 1
	min.Column = (min.Index % columns) + 1
	max.Row = (max.Index / columns) + 1
	max.Column = (max.Index % columns) + 1
	return min, max
}
func main() {
	var rows, columns int
	var matrixValues []int

rowsInput:
	fmt.Println("enter rows number")
	_, err := fmt.Scan(&rows)
	if err != nil {
		fmt.Println("enter valid number")
		goto rowsInput
	}
columnsInput:
	fmt.Println("enter columns number")
	_, err = fmt.Scan(&columns)
	if err != nil {
		fmt.Println("enter valid number")
		goto columnsInput
	}
	fmt.Println("enter matrix values")
	for i := 0; i < rows*columns; i++ {
		var input int
	inputValues:
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("continue entering valid numbers please")
			goto inputValues

		}
		matrixValues = append(matrixValues, input)
	}
	if len(matrixValues) < 1 {
		return
	}
	min, max := GetMinMax(matrixValues, columns)
	fmt.Printf("minimum value is %v (index : %v row : %v column : %v) \n",
		min.Value, min.Index, min.Row, min.Column)
	fmt.Printf("maximun value is %v (index : %v row : %v column : %v)",
		max.Value, max.Index, max.Row, max.Column)
}
