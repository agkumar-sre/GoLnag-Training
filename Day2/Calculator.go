package main

import "fmt"

func divide(a float64, b float64)(float64, string){
	if b == 0{
		return 0, "cannot divide by zero"
	}
	return a/b,""
}

func main(){
	var num1, num2 float64
	var operator string

	fmt.Println("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Println("Enter operatr(+,-,*,/):")
	fmt.Scan(&operator)

	fmt.Println("Enter second number: ")
	fmt.Scan(&num2)

	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		value, err := divide(num1,num2)
		if err != ""{
			fmt.Println("Error: ", err)
		} 
		result = value
	default:
		fmt.Println("Invalid operator")
		return
	}

	fmt.Println("Result: ", result)
}