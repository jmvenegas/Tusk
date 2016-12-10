package parsing

import (
	"strings"
)

func WordBetweenStrings(original, firstString, secondString string) string {
	return original[strings.Index(original, firstString)+len(firstString) : strings.Index(original, secondString)]
}
