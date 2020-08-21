package util

import (
	"bufio"
	"os"
	"strings"
)

// LecturaTeclado reconoce la entrada del teclado
func LecturaTeclado() string {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.Replace(str, "\n", "", -1)
	return strings.TrimSpace(str)
}
