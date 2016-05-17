package handlers

import (
	"log"
	"os"
	"strconv"
	"strings"

	"wattpad/models"
	"wattpad/utils"
)

// get offensive score
func getOffensiveScore(highRiskPhrasesArray []string, lowRiskPhrasesArray []string, input string) (int, error) {
	var numOfHighRisk, numOfLowRisk int = 0, 0

	for i := 0; i < len(highRiskPhrasesArray); i++ {
		if strings.Contains(strings.ToLower(input), strings.ToLower(highRiskPhrasesArray[i])) {
			numOfHighRisk++
		}
	}
	for j := 0; j < len(lowRiskPhrasesArray); j++ {
		if strings.Contains(strings.ToLower(input), strings.ToLower(lowRiskPhrasesArray[j])) {
			numOfLowRisk++
		}
	}

	score := utils.CalOffensiveScore(numOfHighRisk, numOfLowRisk)

	return score, nil
}

// Check offensive of input file
func OffensiveCheck(results []models.Result, outputFile *os.File, highRiskPhrases []string, lowRiskPhrases []string) error {
	for i := 0; i < len(results); i++ {
		//read input file
		inputFile, err := utils.ReadInputFileByName("input_files/" + results[i].FileName)
		if err != nil {
			log.Fatal(err)
		}

		// check offensive phrases & calculate score
		score, err := getOffensiveScore(highRiskPhrases, lowRiskPhrases, inputFile)
		if err != nil {
			log.Fatal(err)
		}

		// put into results object
		results[i].OffensiveScore = score

		// write into output file
		err = utils.WriteOutputFile(outputFile, results[i].FileName+":"+strconv.Itoa(results[i].OffensiveScore)+"\n")
		outputFile.Sync()
	}

	return nil
}
