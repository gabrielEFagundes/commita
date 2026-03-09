package utils

import "fmt"

func LeftAlign(msg string, w int) string {
	if len(msg) > w {
		return msg[:w]
	}

	return fmt.Sprintf("%-*s", w, msg)
}
