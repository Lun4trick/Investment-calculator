package main

import "math"
import "fmt"

func main() {
	const inflationRate = 2.5;
	var investmentAmount float64;
	var expectedReturn float64;
	var years float64;

	fmt.Print("Enter the investment amount: ");
	fmt.Scan(&investmentAmount);

	fmt.Print("Enter the expected return: ");
	fmt.Scan(&expectedReturn);

	fmt.Print("Enter the years: ");
	fmt.Scan(&years);

	futureValue := investmentAmount * math.Pow(1 + expectedReturn / 100, years);
	futureRalValue := futureValue / math.Pow(1 + inflationRate / 100, years);

	fmt.Println("Future value: ", futureValue);
	fmt.Println("Future real value: ", futureRalValue);
}