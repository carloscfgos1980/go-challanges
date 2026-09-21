# Challenge 18: Temperature Converter
Problem Statement
Write a program that converts temperatures between Celsius and Fahrenheit. You'll implement two functions:

CelsiusToFahrenheit - converts a temperature from Celsius to Fahrenheit.
FahrenheitToCelsius - converts a temperature from Fahrenheit to Celsius.
Function Signatures
func CelsiusToFahrenheit(celsius float64) float64
func FahrenheitToCelsius(fahrenheit float64) float64
Input Format
A float64 temperature value in either Celsius or Fahrenheit.
Output Format
A float64 temperature value converted to the other unit.
Conversion Formulas
Celsius to Fahrenheit: F = C × 9/5 + 32
Fahrenheit to Celsius: C = (F - 32) × 5/9
Sample Input and Output
Sample Input 1
CelsiusToFahrenheit(0)
Sample Output 1
32.0
Sample Input 2
FahrenheitToCelsius(32)
Sample Output 2
0.0
Sample Input 3
CelsiusToFahrenheit(100)
Sample Output 3
212.0
Requirements
Round the result to 2 decimal places
Handle negative temperatures correctly
The functions should work with any valid temperature value