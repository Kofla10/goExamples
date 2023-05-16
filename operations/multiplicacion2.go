package operations

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MultiplicationTable2() string {
	scanner := bufio.NewScanner(os.Stdin)
	var number int
	var err error
	var text string

	for {
		fmt.Println("Insert a number")
		if scanner.Scan() {
			number, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Printf("\nInsert a correct number:\n\n")
				continue
			} else {
				break
			}
		}
	}

	// text += fmt.Sprintf("\nThe Start Table Multiplication is %d\n", number)

	text += fmt.Sprintf("The Stard Table Multiplication is %d\n", number)
	for i := 1; i <= 15; i++ {
		value := i * number

		text += fmt.Sprintf("%d * %d = %d\n", i, number, value)
	}
	text += fmt.Sprintf("\n*****************************\n")

	return text

}
