package files

import (
	"exercises/operations"
	"fmt"
	"os"
)

func CreateFile() {
	var text string = operations.MultiplicationTable2()
	file, err := os.Create("./files/txt/file.txt") //function create of package os for sava file .txt
	if err != nil {
		fmt.Println("here Error " + err.Error())
		return //make a return so thah it leaves the function
	}

	fmt.Fprintln(file, text) // write the content of variable text in the file created
	file.Close()
}

func AddTable() {

	// var text string = operations.MultiplicationTable2()
}
