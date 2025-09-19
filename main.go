package main

import (
	"fmt"
)

func main() {
	for {
		printMenu()

		var choice int
		fmt.Print("Enter choice number: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			num := readPositiveInt("Enter number: ")
			fmt.Printf("Number: %d, Sum: %d\n", num, naturalNumber(num))

		case 2:
			num := readPositiveInt("Enter number: ")
			multiplicationTable(num)

		case 3:
			num := readPositiveInt("Enter number: ")
			fmt.Printf("Number: %d, Factorial: %d\n", num, findFactorial(num))

		case 4:
			num := readPositiveInt("Enter number: ")
			fmt.Printf("Number: %d, Reversed number: %d\n", num, reverseNumber(num))

		case 5:
			num := readPositiveInt("Enter number of rows: ")
			printStar(num)

		case 6:
			num := readPositiveInt("Enter number of rows: ")
			printMatrixStar(num)

		case 7:
			num := readPositiveInt("Enter number: ")
			fmt.Printf("Number: %d, First divisor: %d\n", num, getFirstDivisor(num))

		case 8:
			num := readPositiveInt("Enter number: ")
			skipNumberThree(num)

		case 10:
			num := readPositiveInt("Enter number: ")
			printFibonacci(num)

		case 11:
			fmt.Println("Exiting program. Goodbye")
			return

		default:
			fmt.Println("❌ Invalid choice! Please enter a number from the menu.")
		}
	}
}

// ------------------ Menu & Input Helpers ------------------

func printMenu() {
	fmt.Println("\n=== Go Practice Menu ===")
	fmt.Println("1. Sum of N Natural Numbers")
	fmt.Println("2. Multiplication Table")
	fmt.Println("3. Factorial of a Number")
	fmt.Println("4. Reverse Digits of a Number")
	fmt.Println("5. Right Triangle of Stars")
	fmt.Println("6. Square Matrix of Stars")
	fmt.Println("7. Find First Divisor of a Number")
	fmt.Println("8. Skip Multiples of 3")
	fmt.Println("10. Fibonacci Series")
	fmt.Println("11. Exit")
}

func readPositiveInt(prompt string) int {
	var num int
	for {
		fmt.Print(prompt)
		fmt.Scan(&num)
		if num > 0 {
			return num
		}
		fmt.Println("Please enter a number greater than 0")
	}
}

// ------------------ Exercises ------------------

func naturalNumber(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}

func multiplicationTable(num int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", num, i, num*i)
	}
}

func findFactorial(num int) int {
	fact := 1
	for i := 2; i <= num; i++ {
		fact *= i
	}
	return fact
}

func reverseNumber(num int) int {
	reverse := 0
	for num > 0 {
		digit := num % 10
		reverse = reverse*10 + digit
		num /= 10
	}
	return reverse
}

func printStar(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func printMatrixStar(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

func getFirstDivisor(num int) int {
	if num > 1 {
		for i := 2; i <= num; i++ {
			if num%i == 0 {
				return i
			}
		}
	}
	return 1
}

func skipNumberThree(num int) {
	fmt.Println("Numbers from 1 to", num, "(skipping multiples of 3):")
	for i := 1; i <= num; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func printFibonacci(num int) {
	a, b := 0, 1
	fmt.Print("Fibonacci Series: ", a, " ", b, " ")
	for i := 3; i <= num; i++ {
		next := a + b
		fmt.Print(next, " ")
		a, b = b, next
	}
	fmt.Println()
}
