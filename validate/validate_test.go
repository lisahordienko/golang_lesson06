// Homework — Task 2: extend emailCases and/or phoneCases below to at
// least 8 cases each (your mentor may ask for just one of the two
// functions), then implement validate.go until every subtest passes.
//
// Like todo_test.go, this file uses t.Run per case and t.Errorf (not
// t.Fatalf) for the actual assertions, so one wrong case never hides
// the others. Run `go test -v ./validate/...` and read every FAIL line.
package validate

import (
	"strings"
	"testing"
)

const minCases = 8

// emailCases is the table of test cases for ValidateEmail.
//
// TODO: add at least 5 more cases here — for example: whitespace inside
// the address, a missing domain, a trailing dot, consecutive dots, a
// very long local part, or a unicode character.
var emailCases = []struct {
	name  string
	input string
	want  bool
}{
	{"valid simple", "student@softserve.academy", true},
	{"missing at sign", "student-softserve.academy", false},
	{"empty string", "", false},
	{"missing local part", "@example.com", false},
	{"missing domain", "student@", false},
	{"whitespace", "student @example.com", false},
	{"trailing domain dot", "student@example.com.", false},
	{"consecutive local dots", "first..last@example.com", false},
	{"unicode address", "élève@example.com", false},
	{"domain label starts with hyphen", "student@-example.com", false},
	{"single-label domain", "student@localhost", true},
	{"multiple at signs", "a@@example.com", false},
	{"maximum local part length", strings.Repeat("a", 64) + "@example.com", true},
	{"local part exceeds maximum", strings.Repeat("a", 65) + "@example.com", false},
	{"maximum domain label length", "student@" + strings.Repeat("a", 63), true},
	{"domain label exceeds maximum", "student@" + strings.Repeat("a", 64), false},
}

func TestValidateEmail(t *testing.T) {
	if len(emailCases) < minCases {
		t.Fatalf(
			"emailCases has %d case(s), need at least %d — add more edge cases to validate_test.go before this test can pass",
			len(emailCases), minCases,
		)
	}

	for _, tc := range emailCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := ValidateEmail(tc.input)
			if got != tc.want {
				t.Errorf("ValidateEmail(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

// phoneCases is the table of test cases for ValidatePhone.
// This is only required if your mentor asked you to validate phone
// numbers instead of (or in addition to) email addresses.
//
// TODO: add at least 5 more cases here — for example: missing digits,
// letters mixed in, an unexpected country code format, or extra
// separators like spaces, dots or parentheses.
var phoneCases = []struct {
	name  string
	input string
	want  bool
}{
	{"valid with plus", "+380501234567", true},
	{"contains letters", "050-abc-4567", false},
	{"empty string", "", false},
	{"missing plus", "380501234567", false},
	{"too few digits", "+1234567", false},
	{"minimum 8 digits", "+12345678", true},
	{"maximum 15 digits", "+123456789012345", true},
	{"more than 15 digits", "+1234567890123456", false},
	{"leading zero country code", "+012345678", false},
	{"embedded plus", "+1+2345678", false},
	{"spaces are not allowed", "+380 50 123 4567", false},
	{"hyphens are not allowed", "+380-50-123-4567", false},
	{"parentheses are not allowed", "+1(202)5550123", false},
	{"digits only after plus", "+38050.1234567", false},
}

func TestValidatePhone(t *testing.T) {
	if len(phoneCases) < minCases {
		t.Fatalf(
			"phoneCases has %d case(s), need at least %d — add more edge cases to validate_test.go before this test can pass",
			len(phoneCases), minCases,
		)
	}

	for _, tc := range phoneCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := ValidatePhone(tc.input)
			if got != tc.want {
				t.Errorf("ValidatePhone(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
