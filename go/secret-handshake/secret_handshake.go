package secret

import (
	"fmt"
	"strconv"
)

//Creates or tests the handshake
func Handshake(num uint) []string {
	binary := strconv.FormatInt(int64(num), 2)
	fmt.Println(binary)
	return make([]string, 1)
}
