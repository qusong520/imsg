package imsg

import "fmt"
import "io/ioutil"

func encrypt(inputFilePath string, outputFilePath string, msg string, splitText string) {
	fileByte, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("input file read error")
	}

	outByte := []byte(splitText + msg)
	ioutil.WriteFile(outputFilePath, append(fileByte, outByte...), 0644)
}
