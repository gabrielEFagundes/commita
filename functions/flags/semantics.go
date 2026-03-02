package functions

import (
	"fmt"

	functions_auxiliers "github.com/gabrielefagundes/commita/functions/auxiliers"
)

func CommitFeat(argM string) {
	fmt.Print(functions_auxiliers.Append("✨ feat: ", argM))
}
