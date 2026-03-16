package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Confirm(m string) bool {
	bufio := bufio.NewReader(os.Stdin)

	fmt.Printf("%s (y/n) ", m)
	confirm, _ := bufio.ReadString('\n')

	confirm = strings.TrimSpace(strings.ToLower(confirm))
	return confirm == "y" || confirm == "yes"
}
