package common

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func NormalizeString(s string) string {
	result := strings.TrimSpace(s)
	result = cases.Lower(language.Und).String(result)
	return result
}

func TitleString(s string) string {
	return cases.Title(language.Und).String(s)
}
