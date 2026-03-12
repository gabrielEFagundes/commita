package utils

import "strings"

func ValidateEmail(mx string, u string) (string, string) {
	if strings.Contains(mx, "@") {
		return mx, u
	}

	return u, mx
	// enhance this because it makes no sense
}
