package validate

import "strings"

// Package validate provides simple syntactic validators for common
// user-input formats.

// ValidateEmail accepts ASCII dot-atom local parts and domains whose labels
// contain letters, digits, or interior hyphens. It rejects Unicode, quoted
// local parts, consecutive dots, trailing dots, and addresses over 254 bytes.
func ValidateEmail(s string) bool {
	if len(s) == 0 || len(s) > 254 || strings.Count(s, "@") != 1 {
		return false
	}

	parts := strings.SplitN(s, "@", 2)
	local, domain := parts[0], parts[1]
	if len(local) == 0 || len(local) > 64 || len(domain) == 0 {
		return false
	}
	if local[0] == '.' || local[len(local)-1] == '.' || strings.Contains(local, "..") {
		return false
	}
	for _, char := range local {
		if !isEmailLocalChar(char) {
			return false
		}
	}

	labels := strings.Split(domain, ".")
	for _, label := range labels {
		if len(label) == 0 || len(label) > 63 || !isASCIIAlphaNumeric(rune(label[0])) || !isASCIIAlphaNumeric(rune(label[len(label)-1])) {
			return false
		}
		for _, char := range label {
			if !isASCIIAlphaNumeric(char) && char != '-' {
				return false
			}
		}
	}
	return true
}

// ValidatePhone accepts E.164-style input: a leading plus followed by 8 to
// 15 digits, with the first digit non-zero. Spaces and separators are rejected.
func ValidatePhone(s string) bool {
	if len(s) < 9 || len(s) > 16 || s[0] != '+' || s[1] < '1' || s[1] > '9' {
		return false
	}
	for i := 2; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func isEmailLocalChar(char rune) bool {
	return isASCIIAlphaNumeric(char) || strings.ContainsRune(".!#$%&'*+-/=?^_`{|}~", char)
}

func isASCIIAlphaNumeric(char rune) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9'
}
