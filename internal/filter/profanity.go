package filter

import "regexp"

func FilterProfanity(text string) string {
	// Regex pattern to match standalone words (case-insensitive)
	profanityRegex := regexp.MustCompile(`\b(?i)(fuck|shit|bitch)\b`)
	return profanityRegex.ReplaceAllString(text, "****")
}
