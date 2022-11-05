package RPC

import (
	"fmt"
	"strings"
)

func Expect(symbol, value string) bool {
	if strings.Trim(value, "") != symbol {
		fmt.Println("\033[38;5;121m| RPC warning ^^^^^ \033[38;5;128m\n\t\t|Unexpected symbol| \033[38;5;210m => " + value + " <= \n\t\t\t\033[38;5;128m|Exptecting| \033[38;5;210m => " + symbol + " <= ")
		return false
	} else {
		return true
	}
}
