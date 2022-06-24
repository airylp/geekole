package main

import (
	"fmt"
	"os"
)

//geekole.com

var path = "textfile.txt"

func main() {
	createFile()
	writeFile()
	fmt.Println("Text File Created Successfully")
}
func createFile() {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		checkError(err)
		defer file.Close()
	}
}
func writeFile() {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()
	_, err = file.WriteString("Hello \n")
	checkError(err)
	_, err = file.WriteString("World! \n")
	checkError(err)
	err = file.Sync()
	checkError(err)
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
