package imsg

import "fmt"
import "io/ioutil"
import "strings"

func decrypt(filePath string, splitText string) string {
	fileByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Read file error")
	}

	ss := string(fileByte)
	specIndex := strings.LastIndex(ss, splitText)
	if specIndex >= 0 {
		return ss[specIndex+len(splitText):]
	} else {
		return ""
	}
}
