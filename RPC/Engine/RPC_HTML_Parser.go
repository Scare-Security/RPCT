package RPC

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var RPCDATA []string

func Open_Store(filename string) {
	f, x := os.Open(filename)
	if x != nil {
		fmt.Println(x)
	} else {
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			RPCDATA = append(RPCDATA, scanner.Text())
		}
		parser()
	}
}

func parser() {
	for i := 0; i < len(RPCDATA); i++ {
		if !UnitSymbolFoundBegin {
			if Symbols[strings.Trim(RPCDATA[i], " ")] == "Init" {
				line = i
				if !Expect("{", strings.Trim(RPCDATA[i+1], " ")) {
					fmt.Println("CODE: next_line_false | INVALID_SEQUENCE (check1")
					os.Exit(0)
				} else {
					fmt.Println("\033[38;5;20m[\033[38;5;123mEngine\033[38;5;20m]: \033[38;5;123mInformation - \033[38;5;57mFound valid sequence ")

					if !Expect("(", strings.Trim(RPCDATA[i+2], " ")) {
						fmt.Println("CODE: next_line_false | INVALID_SEQUENCE (check2")
						os.Exit(0)
					} else {
						fmt.Println("\033[38;5;20m[\033[38;5;123mEngine\033[38;5;20m]: \033[38;5;123mInformation - \033[38;5;57mFound valid sequence ")
						if !Expect("[", strings.Trim(RPCDATA[i+3], " ")) {
							fmt.Println("CODE: next_line_false | INVALID_SEQUENCE (check3")
							os.Exit(0)
						} else {
							fmt.Println("\033[38;5;20m[\033[38;5;123mEngine\033[38;5;20m]: \033[38;5;123mInformation - \033[38;5;57mFound valid sequence ")
							if !Expect("RPC -> RPCL_Count | RPC.BRICK => {{{", strings.Trim(RPCDATA[i+4], " ")) {
								fmt.Println("CODE: next_line_false | INVALID_SEQUENCE")
								os.Exit(0)
							} else {
								fmt.Println("\033[38;5;20m[\033[38;5;123mEngine\033[38;5;20m]: \033[38;5;123mInformation - \033[38;5;57mBrick Sequence Detected T(RPC) ")

								if Expect("RPC:::Unit::SYMBOL | RPC.BRICK START", strings.Trim(RPCDATA[i+5], " ")) {
									fmt.Println("\033[38;5;20m[\033[38;5;123mEngine\033[38;5;20m]: \033[38;5;123mInformation - \033[38;5;57mUnit Symbol on table T<RPC -> RPCL_Count | RPC.BRICK => {{{}}}+++___> found")
									UnitSymbolFoundBegin = true
									// continue generation and interpret symbols until eos ( End Of Section )
								}
							}
						}
					}
				}
			}
		} else { // this means it was true and you can start interpreting data
			if SymbolRegexuse.FindAllStringSubmatch(strings.Trim(RPCDATA[i], " "), -1) != nil {
				break
			} else {
				RPCC.CodeBlock = append(RPCC.CodeBlock, strings.Trim(RPCDATA[i], " "))
			}
			RPCC.CodeBlock = ValueRemover(RPCC.CodeBlock)
		}
	}
	Interpreter()

}
