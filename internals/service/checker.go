package service

import (
	"net/mail"
	"strings"
	"unicode"
)

func CheckUsername(s string) bool {
	if len(s) < 1 {
		return false
	}

	for _, ch := range s {
		if ch < 33 || ch > 126 {
			return false
		}
	}
	return true
}

func CheckPassword(s string) bool {
	var (
		minLength = false
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)

	if len(s) > 7 {
		minLength = true
	}

	for _, ch := range s {
		if unicode.IsUpper(ch) {
			hasUpper = true
		} else if unicode.IsLower(ch) {
			hasLower = true
		} else if unicode.IsDigit(ch) {
			hasNumber = true
		}
	}

	if !minLength || !hasLower || !hasUpper || !hasNumber {
		return false
	}

	return true
}

func CheckEmail(s string) bool {
	_, err := mail.ParseAddress(s)
	if err != nil {
		return false
	}
	return true
}

func CheckTitle(s string) (string, bool) {
	s = strings.TrimSpace(s)

	if s == "" {
		return "", false
	}

	if len(s) > 40 {
		return "", false
	}

	return s, true
}

func CheckPostContent(s string) (string, bool) {
	s = strings.TrimSpace(s)

	if s == "" {
		return "", false
	}

	if len(s) > 1000 {
		return "", false
	}

	return s, true
}

func Contains(arr []string, s string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}

func CheckCategory(arr []string) ([]string, bool) {
	category := []string{"IT", "Cars", "Football", "CyberSport", "Other"}

	for _, v := range arr {
		if !Contains(category, v) {
			return nil, false
		}
	}
	arr = UniqueCat(arr)
	return arr, true
}

func UniqueCat(arr []string) []string {
	uniqueMap := make(map[string]bool)
	for _, s := range arr {
		uniqueMap[s] = true
	}

	uniqueSlice := make([]string, 0, len(uniqueMap))

	// Iterate over the keys of the map and append them to the slice
	for s := range uniqueMap {
		uniqueSlice = append(uniqueSlice, s)
	}
	return uniqueSlice
}

func CheckComment(s string) (string, bool) {
	s = strings.TrimSpace(s)

	if s == "" {
		return "", false
	}

	if len(s) > 280 {
		return "", false
	}

	return s, true
}
