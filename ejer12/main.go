package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func readFile() {
	file, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
	} else {
		fmt.Println(string(file))
	}
}

func readFile2() {
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			register := scanner.Text()
			fmt.Printf("Line >" + register + "\n")
		}
	}
	file.Close()
}

func recordFile() {
	file, err := os.Create("./newFile.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	fmt.Fprintln(file, "This a new Line")
	file.Close()
}

func recordFile2() {
	fileName := "./newFile.txt"
	if !Append(fileName, "This a new second Line") {
		fmt.Println("Error in second file")
	}
}

func Append(fileName string, text string) bool {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file", err)
		return false
	}

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error writing file", err)
		return false
	}

	return true
}

func main() {
	readFile()
	readFile2()

	recordFile()
	recordFile2()
}
