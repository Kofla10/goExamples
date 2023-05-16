package exercices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func SquareArea() int {
	scanner := bufio.NewScanner(os.Stdin)
	var number1 int = 0
	var number2 int = 0

	fmt.Printf("Insert one number")
	if scanner.Scan() {
		number1, _ = strconv.Atoi(scanner.Text())
	}

	fmt.Printf("Insert one number")
	if scanner.Scan() {
		number2, _ = strconv.Atoi(scanner.Text())
	}

	return number1 * number2
}

func NumberOdd() string {
	var number int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Insert a number")
	if scanner.Scan() {
		number, _ = strconv.Atoi(scanner.Text())
	}

	if number%1 == 0 {
		fmt.Println(number % 2)
		return "true"
	} else {
		fmt.Println(number % 2)
		return "false"
	}

}
