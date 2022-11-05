package RPC

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Write_Frame(filename string, frame string) {
	f, x := os.Create(filename)
	if x != nil {
		log.Fatal(x)
	}
	defer f.Close()
	_, x2 := f.WriteString(frame)
	if x2 != nil {
		log.Fatal(x2)
	}
	fr, x3 := os.Open(filename)
	if x3 != nil {
		log.Fatal(x3)
	}
	defer fr.Close()
	scanner := bufio.NewScanner(fr)
	fmt.Println("\033[38;5;20m[\033[38;5;123mEngine\033[38;5;20m]: \033[38;5;123mInformation - \033[38;5;57mIncoming data frame sequence for translated script")
	fmt.Print("\n")
	fmt.Println("\033[38;5;214m        ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	var line int = 0
	for scanner.Scan() {
		line++
		fmt.Println("[", line, "]", "\033[38;5;214m\t┃\033[38;5;57m \t ", scanner.Text())
	}
	fmt.Println("[ef]", "\033[38;5;214m   ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
}
