package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Calculo() {
	scanner := bufio.NewScanner(os.Stdin)
	var num1, num2, num3 int
	var err error

	fmt.Println("Insert firts value")
	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Error with the number one %s", err.Error())
			return
		}
	}

	fmt.Println("Insert second value")
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Error with the number one %s", err.Error())
			return
		}
	}

	// fmt.Println(num1 + num2 + num3)
	fmt.Printf("\n\n**********Menú**********\n1. Suma.\n2. Resta.\n3. Multiplicación.\n\n********END Menú********\n\n")

	if scanner.Scan() {
		num3, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Error with the number one %s", err.Error())
			return
		}
	}

	sum := func(number1 int, number2 int) int {
		suma := number1 + number2
		fmt.Printf("El resultado de la suma es: %d + %d = %d\n", num1, num2, suma)
		return suma
	}

	resta := func(number1 int, number2 int) int {
		resta1 := number1 - number2
		fmt.Printf("El resultado de la resta es: %d - %d = %d\n", num1, num2, resta1)
		return resta1
	}

	multiplicacion := func(number1 int, number2 int) int {
		multiplicacion1 := number1 * number2
		fmt.Printf("El resultado de la multiplicacion es: %d * %d = %d\n", num1, num2, multiplicacion1)
		return multiplicacion1
	}

	switch num3 {
	case 1:
		sum(num1, num2)
	case 2:
		resta(num1, num2)
	case 3:
		multiplicacion(num1, num2)
	default:
		fmt.Println("error opcion no valida")
	}

}
