package supporting

import (
	"bufio"
	"fmt"
	"os"
)

func input(printString string) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(printString)

	scanner.Scan()
	return scanner.Text()
}
