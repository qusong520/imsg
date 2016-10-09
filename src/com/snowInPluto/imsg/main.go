package main

import "fmt"
import "flag"

const specialText = "#########"

type Cmd struct {
	decrypt   bool
	input     string
	output    string
	splitText string
	arg       string
}

func main() {
	cmd := Cmd{}
	flag.BoolVar(&cmd.decrypt, "d", false, "decrypt the file and get the message")
	flag.StringVar(&cmd.input, "i", "", "the input file path")
	flag.StringVar(&cmd.output, "o", "", "the out file path")
	flag.StringVar(&cmd.splitText, "s", specialText, "the split text between origin file text and message")
	flag.Parse()

	if flag.NArg() > 0 {
		cmd.arg = flag.Arg(0)
	}

	if cmd.decrypt {
		if cmd.input != "" {
			fmt.Println(decrypt(cmd.input, cmd.splitText))
		} else if cmd.arg != "" {
			fmt.Println(decrypt(cmd.arg, cmd.splitText))
		} else {
			fmt.Println("File path is empty")
		}
	} else {
		if cmd.output == "" {
			fmt.Println("Output file path is empty")
		}
		if cmd.arg == "" {
			fmt.Println("Message is empty")
		}

		encrypt(cmd.input, cmd.output, cmd.arg, cmd.splitText)
	}
}
