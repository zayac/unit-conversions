package main

import "fmt"

/*

Sample input/output:

The program converts temperature from Celsius to Fahrenheit.
Enter the temperature in Celsius:
100
The temperature in Fahrenheit: 212

*/
func main() {
	var celsius float64
	fmt.Println("The program converts temperature from Celsius to Fahrenheit.")
	fmt.Println("Enter the temperature in Celsius:")
	fmt.Scan(&celsius)
	fahrenheit := 1.8 * celsius + 32
	fmt.Println("The temperature in Fahrenheit:", fahrenheit)
}