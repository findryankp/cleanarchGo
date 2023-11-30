package generates

import (
	"strings"
	"unicode"
)

func ToTitle(featuresName string) string {
	return strings.Map(unicode.ToTitle, featuresName)
}

func ToLower(featuresName string) string {
	return strings.ToLower(featuresName)
}
