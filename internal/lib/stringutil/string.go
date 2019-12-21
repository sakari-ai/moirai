package stringutil

import (
	"bytes"
	"net/url"
	"regexp"
	"strings"
	"text/template"
)

const (
	PatternUrl = "^https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)$"
)

func TrimLeadingHead(sa []string) []string {
	firstNonEmptyStringIndex := 0
	for i := 0; i < len(sa); i++ {
		if sa[i] == "" {
			firstNonEmptyStringIndex++
		} else {
			break
		}
	}
	return sa[firstNonEmptyStringIndex:len(sa)]
}

func UrlValidateString(pattern string, s string) (bool, error) {
	matched, err := regexp.MatchString(pattern, s)
	return matched, err
}

func Obfuscate(value string, size int) string {
	if len(value) <= size {
		return value
	}
	var result string
	for index, char := range value {
		if index > size {
			result += "*"
		} else {
			result += string(char)
		}
	}
	return result
}

func DifferDomainAndPath(left, right string) bool {
	leftQuery, _ := url.Parse(strings.TrimSpace(left))
	rightQuery, _ := url.Parse(strings.TrimSpace(right))

	if leftQuery.Host == rightQuery.Host && leftQuery.Path == rightQuery.Path && leftQuery.Scheme == rightQuery.Scheme {
		return false
	}
	return true
}

func MutualString(source, destination string, length int) string {
	if length > len(source) || length > len(destination) {
		return ""
	}
	for i := 0; i <= len(source)-length; i++ {
		detectString := source[i : i+length]
		if strings.Contains(destination, detectString) {
			return detectString
		}
	}
	return ""
}

func Format(fmt string, args map[string]interface{}) string {
	var msg bytes.Buffer
	tmpl, err := template.New("").Parse(fmt)

	if err != nil {
		return fmt
	}
	tmpl.Execute(&msg, args)
	return msg.String()
}
