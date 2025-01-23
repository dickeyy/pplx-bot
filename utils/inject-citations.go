package utils

import (
	"fmt"
	"regexp"
	"strconv"
)

func InjectCitations(text string, citations []string) string {
	// Compile the regex pattern for finding citations
	re := regexp.MustCompile(`\[(\d+)\]`)

	// Replace each citation with the formatted link
	result := re.ReplaceAllStringFunc(text, func(match string) string {
		// Extract the number from between brackets
		num := re.FindStringSubmatch(match)[1]
		index, _ := strconv.Atoi(num)

		// Create the new format: [[x](citation)]
		// Note: index-1 because citations are 1-based but slice is 0-based
		return fmt.Sprintf("[[%s](%s)]", num, citations[index-1])
	})

	return result
}
