package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		_, json := processStdin()
		err, _ := isValidJson(json)
		fmt.Println(err)
	} else {
		// Process each file provided as an argument
		for _, filename := range flag.Args() {
			_, json := processFile(filename)
			err, _ := isValidJson(json)
			fmt.Printf("%s\t%s\n", filename, err)
		}
	}
}

func isValidJson(data string) (err error, isJSON bool) {
	if len(data) == 0 {
		// empty data
		return errors.New("code 1"), true
	} else if data[len(data)-1] == '}' && data[0] == '{' {
		// valid JSON
		return errors.New("code 0"), false
	} else {
		// invalid JSON
		return errors.New("code 1"), false
	}
}

func processFile(filename string) (err error, data string) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error parsing data: %v\n", err)
	}
	dataString := string(bytes)
	return err, dataString
}

func processStdin() (err error, data string) {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {

		fmt.Printf("Error parsing data: %v\n", err)
	}
	dataString := string(bytes)
	return err, dataString
}
