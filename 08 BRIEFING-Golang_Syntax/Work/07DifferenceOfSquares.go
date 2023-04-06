package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func SquareOfSums(N int) int {
	sum := 0
	for i := 1; i <= N; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(N int) int {
	sum := 0
	for i := 1; i <= N; i++ {
		sum += i * i
	}
	return sum
}
func Difference(N int) int {
	return SquareOfSums(N) - SumOfSquares(N)
}

func main() {
	var TextLine string
	fmt.Print("How many natural numbers: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	TextLine = scanner.Text()

	Number, err := strconv.Atoi(TextLine)
	if err != nil {
		fmt.Println("Error during conversion")
	} else {
		fmt.Print("Difference Of Squares (", Number, " natural numbers): ", Difference(Number),"\n")
	}
}
