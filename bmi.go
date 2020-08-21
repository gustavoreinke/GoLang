// This code is to discover your BMI (Body Mass Index) - Male Case Only

package main

import (
	"fmt"
)

func main() {

	var height, weight, bmi float64
	var name string

	fmt.Println("What is your name?")
	fmt.Scan(&name)

	fmt.Println("What is your weight (kilograms)?")
	fmt.Scan(&weight)

	fmt.Println("What is your height (meters)?")
	fmt.Scan(&height)

	bmi = weight / (height * height)
	bmiString := fmt.Sprintf("%.2f", bmi)

	fmt.Println("Your BMI is ", bmiString)

	if bmi < 18.5 {
		fmt.Println(name, ",", "Careful. you are underweight.")
	} else if bmi < 25 {
		fmt.Println(name, ",", "Congratulations! You're on your ideal weight.")
	} else if bmi < 30 {
		fmt.Println(name, ",", "Careful. You are overrweight.")
	} else if bmi < 35 {
		fmt.Println(name, ",", "Careful. You are with obesity grade 1.")
	} else if bmi < 40 {
		fmt.Println(name, ",", "Careful. you are with obesity grade 2.")
	} else {
		fmt.Println(name, ",", "WARNING!! You are with obesity grade 3.")
	}
}
