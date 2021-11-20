package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if os.Args[1] == "--check" && len(os.Args) == 3 {
		// Get the checksum file and convert the file into string
		sha := os.Args[2]
		shaFileContent, err := ioutil.ReadFile(sha)
		if err != nil {
			panic(err)
		}
		shaFileContentInStr := strings.Fields(string(shaFileContent))

		// Get the file name from the checksum and convert
		// the data of the real file into string
		originalFileName := shaFileContentInStr[1]
		originalFileContent, err := ioutil.ReadFile(originalFileName)
		if err != nil {
			panic(err)
		}
		originalFileContentInStr := string(originalFileContent)
		// fmt.Println(originalFileContentInStr)

		// Check the of the string generated from the hash and the original
		// file content is same or not
		if hex.EncodeToString([]byte(shaFileContentInStr[0])) != originalFileContentInStr {
			fmt.Errorf("want %v; got %v", originalFileContentInStr, shaFileContentInStr[0])
		} else {
			fmt.Printf("%v : OK", originalFileName)
		}
	}
	if len(os.Args) != 2 {
		fmt.Errorf("Error: number of argument is %v", len(os.Args))
		os.Exit(0)
	}

	fileContent, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	h := sha256.Sum256(fileContent)
	fmt.Printf("%x  %v", h, os.Args[1])
}
