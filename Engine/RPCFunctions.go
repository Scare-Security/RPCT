package RPC

import (
	"fmt"
	"strings"
)

func FindSubStrIndex(x, y string, initialindex int) int {
	i := strings.Index(x, y)
	if i > -1 {
		// extra indexing : debug
		//		c1 := x[:i]
		//		c2 := x[i+1:]
	} else {
		fmt.Println("Index not found in string > ", x)
	}
	return i
}
