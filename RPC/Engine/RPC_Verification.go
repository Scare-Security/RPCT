package RPC

import (
	"fmt"
	"os"
	"path/filepath"
)

// The engine will only be looking for RPC template names and files
func Verify_File(filename string) {
	if _, x := os.Stat(filename); x != nil {
		// file does not exist, throw engine error
		fmt.Println(x)
	} else {
		// File exists

		if File_Extensions_Possible[filepath.Ext(filename)] != "RPC template file HTML" {
			// Throw incorrect file
		} else {
			// verify and move on
			switch File_Extensions_Possible[filepath.Ext(filename)] {
			case "RPC template file HTML":
				//fmt.Println("[+] Starting HTML templating and generation, Found RPCT file ( Radical Processing Templating Language )")
				// Start HTML generation
				Open_Store(filename)
			}
		}
	}
}
