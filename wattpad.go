// wattpad project wattpad.go
package main

import (
	"fmt"
	"log"

	"wattpad/handlers"
	"wattpad/models"
	"wattpad/utils"
)

func main() {

	fmt.Println("************ Coding Challenge ************")

	// initial results object
	r := models.Result{}

	// read high risk phrases
	highRiskPhrases, err := utils.ReadRiskPhrasesFileByName("risk_phrases_files/high_risk_phrases.txt")
	if err != nil {
		log.Fatal(err)
	}

	// read low risk phrases
	lowRiskPhrases, err := utils.ReadRiskPhrasesFileByName("risk_phrases_files/low_risk_phrases.txt")
	if err != nil {
		log.Fatal(err)
	}

	// create a empty file
	outputFile, err := utils.CreateNewFile("output_files/output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// read all input files
	files, err := utils.ReadInputFileByDir("input_files/")
	if err != nil {
		log.Fatal(err)
	}

	// initial filename & offensive score
	results, err := r.GetResults(files)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)

	// Check offensive phrases
	err = handlers.OffensiveCheck(results, outputFile, highRiskPhrases, lowRiskPhrases)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("************ Done! ************")
	}
}
