package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// Get the file content of the sha checksum file and get the target file name
// and compute the hash and check it with the hash in the checksum file
func main() {
	if len(os.Args) == 3 && os.Args[1] == "--check" {
		shaFileContent := getFileContent(os.Args[2])
		shaFileContentInStr := strings.Fields(string(shaFileContent))

		originalFileName := shaFileContentInStr[1]
		originalFileContent := getFileContent(shaFileContentInStr[1])

		hex := getSHA256(originalFileContent)

		if shaFileContentInStr[0] != hex {
			fmt.Printf("%v: FAILED\nsha256sum: WARNING: 1 computed checksum did NOT match", originalFileName)
		} else {
			fmt.Printf("%v : OK", originalFileName)
		}
	} else if len(os.Args) == 2 {
		fileContent := getFileContent(os.Args[1])
		var sha string = getSHA256(fileContent)
		fmt.Printf("%v  %v", sha, os.Args[1])
	} else {
		fmt.Errorf("Number of argument: %v", len(os.Args))
		os.Exit(0)
	}
}

func getFileContent(filename string) []byte {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return fileContent
}

func getSHA256(data []byte) string {
	bytes := sha256.Sum256(data)
	hex := hex.EncodeToString(bytes[:])
	return hex
}
