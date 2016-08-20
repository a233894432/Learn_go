/*
Project:Learn_go
Time:2016/8/20-16:18
author:Dio
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
)

func main() {
	inputFile, _ := os.Open("goprogram.go")
	outputFile, _ := os.OpenFile("goprogramT.go", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			return
		}
		outputString := string([]byte(inputString)[2:5]) + "\r\n"
		n, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(n)
	}
	fmt.Println("Conversion done")
}
