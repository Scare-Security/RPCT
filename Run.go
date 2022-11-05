package main

import (
	"fmt"
	RPC_Engine "main/Engine"
)

func init() {

}

func main() {
	fmt.Println("\033[38;5;123m", RPC_Engine.Banner)
	RPC_Engine.Verify_File("HTTP_Info.rpct")
}
