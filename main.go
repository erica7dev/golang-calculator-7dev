package main

import (
	"fmt"
	"math"
)

func main (){
	calculate()
}
 
func calculate(){
	var operator string
	var n1, n2 int

	fmt.Println("Welcome to the calculator!")
	fmt.Println("Enter the operator:")
	fmt.Scanln(&operator)
	fmt.Println("Enter the first number:")
	fmt.Scanln(&n1)
	fmt.Println("Enter the second number:")
	fmt.Scanln(&n2)
	output := 0

	switch operator{
		case "+":
			output = n1 + n2
			break
		case "-":
			output = n1 - n2
		case "/":
			output = n1 / n2
		case "*":
			output = n1 * n2
		case "^":
			output = n1 ^ n2
		case "%":
			output = n1 % n2
		case "sqrt":
			output = int(math.Sqrt(float64(n1)))
		default:
			fmt.Println("Invalid operator")
	}
	fmt.Printf("%d %s %d = %d", n1, operator, n2, output)
}

/*func main(){
	fmt.Println(sum(1, 2))
	fmt.Println(sub(1, 2))
	fmt.Println(mul(1, 2))
	fmt.Println(div(1, 2))
	fmt.Println(expo(1, 2))
	fmt.Println(square_root(4))
}

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func expo(a, b int) int {
	return a ^ b
}

func square_root(a int) int {
	return int(math.Sqrt(float64(a)))
}
*/