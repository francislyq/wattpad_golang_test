package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Write Results into File
func WriteOutputFile(outputFile *os.File, output string) error {
	_, err := outputFile.WriteString(output)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// Read Input File By Name
func ReadInputFileByName(fileName string) (string, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(content), nil
}

// Read Risk Phrases File By Name
func ReadRiskPhrasesFileByName(fileName string) ([]string, error) {
	riskPhrases, err := ReadInputFileByName(fileName)
	if err != nil {
		log.Fatal(err)
	}

	riskPhrasesArray := strings.Split(riskPhrases, "\n")

	return riskPhrasesArray, nil
}

// Create a New File
func CreateNewFile(fileName string) (*os.File, error) {
	outputFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	//defer outputFile.Close()

	return outputFile, nil
}

// Read Input Files By Directory
func ReadInputFileByDir(dirName string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	return files, nil
}
