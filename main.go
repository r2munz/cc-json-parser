package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		_, dataString := processStdin()
		jsonErr, _ := jsonParser(dataString)
		fmt.Printf("%s", jsonErr)
	} else {
		// Process each file provided as an argument
		for _, filename := range flag.Args() {
			_, dataString := processFile(filename)
			jsonErr, _ := jsonParser(dataString)
			fmt.Printf("%s\t%s", filename, jsonErr)
			continue
		}
	}
}

func jsonParser(dataString string) (jsonErr error, dataMap map[string]interface{}) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(dataString), &data)
	if err != nil {
		return errors.New("code 1 - could not unmarshal json \n"), data
	} else {
		return errors.New("code 0 - json parsed \n"), data
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
