package console

import (
	"bufio"
	"os"
)

func BlockUntilEnter() {
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}
