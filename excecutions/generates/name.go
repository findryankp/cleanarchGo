package generates

import "strings"

func ToTitle(featuresName string) string {
	return strings.Title(strings.ToLower(featuresName))
}

func ToLower(featuresName string) string {
	return strings.ToLower(featuresName)
}
