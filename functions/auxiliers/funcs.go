package functions_auxiliers

import "fmt"

func Append(prefix string, msg string) string {
	return fmt.Sprintf("%s: %s", prefix, msg)
}
