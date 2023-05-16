package operations

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Multiplicacion(number1 string, number2 string) (int, int, int, string) {
	fmt.Println("test one")

	value1, _ := strconv.Atoi(number1)
	value2, _ := strconv.Atoi(number2)
	if x := value1 * value2; x >= 200 {
		return value1, value2, x, "The value is greater to 200"
	} else {
		return value1, value2, x, "The values is little to 200"
	}

}

func MultiplicationTable() {

	scanner := bufio.NewScanner(os.Stdin)
	var number int
	var err error
	for {
		fmt.Println("Insert number for multiplication table")
		if scanner.Scan() {
			number, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}
	for i := 1; i <= 10; i++ {
		result := i * number
		fmt.Printf("%d * %d = %d\n", i, number, result)
	}

}
