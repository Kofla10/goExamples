package files

import (
	"bufio"
	"exercises/operations"
	"fmt"
	"io/ioutil"
	"os"
)

var fileName = "./files/txt/file.txt"

func CreateFile() {
	var text string = operations.MultiplicationTable2()
	file, err := os.Create(fileName) //function create of package os for sava file .txt
	if err != nil {
		fmt.Println("here Error " + err.Error())
		return //make a return so thah it leaves the function
	}

	fmt.Fprintln(file, text) // write the content of variable text in the file created
	file.Close()
}

func AddTable() {

	var text string = operations.MultiplicationTable2()
	if !Append(fileName, text) {
		fmt.Println("Error in the concatenation of file")
	}

}

func Append(filename string, text string) bool {
	//creation of bafer
	//read about flag and os.OpenFile
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error in the reading file " + err.Error())
		return false
	}

	_, err = arch.WriteString(text)
	if err != nil {
		fmt.Println("Error during the WriteString " + err.Error())
	}
	return true
}

// method for reading file, shape old
func ReadingFile() {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error in the Reading File " + err.Error())
		return
	}
	//convert file to string for reading
	fmt.Println(string(file))
}

// method for reading file two, new shape
func ReadinFile2() {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error is the reading file with this shape " + err.Error())
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		register := scanner.Text()

		fmt.Println("=> " + register)
	}
	file.Close()

}
