package validators

import (
	"unicode"
)

// ValidatePasswordStrength checks if the password is strong.
func ValidatePasswordStrength(password string) bool {
	var hasMinLen, hasUpper, hasLower, hasNumber bool

	// Check if the password has at least 8 characters
	if len(password) >= 8 {
		hasMinLen = true
	}

	// Iterate over each character in the password to validate its strength
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}

	// Ensure the password has all required components
	return hasMinLen && hasUpper && hasLower && hasNumber
}
