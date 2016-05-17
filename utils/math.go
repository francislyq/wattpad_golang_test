package utils

// Calculate offense score
func CalOffensiveScore(numOfHighRisk int, numOfLowRisk int) int {
	return (numOfLowRisk + numOfHighRisk*2)
}
