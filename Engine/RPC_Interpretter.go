package RPC

import (
	"fmt"
	"os"
	"strings"
)

var RPC_Notes = []string{
	"!<<",
	"!>>",
	"!!!-> rpcn",
	"!!!<-",
}

func Indexer(x string, y string) {
	i := strings.Index(x, y)
	fmt.Println("Index: ", i)
	if i > -1 {
		chars := x[:i]
		arefun := x[i+1:]
		fmt.Println(chars)
		fmt.Println(arefun)
	} else {
		fmt.Println("Index not found")
		fmt.Println(x)
	}
}

func Interpreter() {
	for p := 0; p < len(RPCC.CodeBlock); p++ {
		//
		//
		//
		// Find mods by macro block as standard library
		libs := LoadingVariablesBySTD.FindAllStringSubmatch(strings.Trim(RPCC.CodeBlock[p], " "), 1)
		for i := range libs {
			if libs != nil {
				TemplateData.Import_names_STDLIB = append(TemplateData.Import_names_STDLIB, libs[i][1])
			}
		}
		//
		//
		// Find loaded libraries local
		stdlibs := LoadingVariablesFrom.FindAllStringSubmatch(strings.Trim(RPCC.CodeBlock[p], " "), 1)
		for l := range stdlibs {
			if stdlibs != nil {
				TemplateData.Import_VariablesFrom = append(TemplateData.Import_VariablesFrom, stdlibs[l][1])
			}
		}
		//
		// Find package by macro block
		//
		Package := LoadingPackagesAs.FindStringSubmatch(RPCC.CodeBlock[p])
		if len(Package) > 1 {
			TemplateData.Package_name = Package[1]
		}
		//
		//
		//
		// Find and execute genertion
		//

		Argument := RPC_Generate_Arguments.FindStringSubmatch(strings.Trim(RPCC.CodeBlock[p], " "))
		if len(Argument) > 1 {
			method := RPC_Generate_LIA.FindStringSubmatch(strings.Trim(RPCC.CodeBlock[p], " "))
			if len(method) > 1 {
				Generate_Under_For_Loop = true
			}
			argumentcolor = Argument[1]
		}

		// locate for loop
		loop := RPC_Generate_FOR.FindStringSubmatch(strings.Trim(RPCC.CodeBlock[p], " "))
		if len(loop) > 1 {
			data := loop[0]
			value_range := strings.Trim(data[18:], "{")
			Argument1 := strings.Trim(data[4:5], ",")
			Argument2 := strings.Trim(data[5:8], ",")
			if argumentcolor != "" {
				if Generate_Under_For_Loop {
					functioncallwithvals := "Generate_LIA_Tag(%s)"
					functioncallwithvals = fmt.Sprintf(functioncallwithvals, argumentcolor)
					RPCBOD.BodyFunctions = append(RPCBOD.BodyFunctions, Create_Method("function", LIA_Template))
					RPCBOD.Body = append(RPCBOD.Body, Create_Method("for range", Argument1, Argument2, value_range, functioncallwithvals))
				} else {
					RPCBOD.Body = append(RPCBOD.Body, Create_Method("for range", Argument1, Argument2, value_range, argumentcolor))
				}
			}
		}
		// Locate for loop length based
		looplen := RPCgenerateForWithVals.FindStringSubmatch(strings.Trim(RPCC.CodeBlock[p], " "))
		if len(looplen) > 1 {
			data := looplen[0]
			// expecting two single character values
			forchar1 := strings.Trim(data[11:13], " ")  // what the second char equals
			foreqchar := strings.Trim(data[16:18], " ") // what the first char equals
			forlenchar := strings.Trim(data[24:FindSubStrIndex(data, "|", 24)], " ")
			if forchar1 == "" || len(forchar1) > 1 {
				fmt.Println("Char to long 1")
				os.Exit(0)
			}
			if argumentcolor != "" {
				if Generate_Under_For_Loop {
					functioncallwithvals := "Generate_LIA_Tag(%s)"
					functioncallwithvals = fmt.Sprintf(functioncallwithvals, argumentcolor)
					RPCBOD.BodyFunctions = append(RPCBOD.BodyFunctions, Create_Method("function", LIA_Template))
					RPCBOD.Body = append(RPCBOD.Body, Create_Method("for valg", forchar1, foreqchar, forchar1, forlenchar, functioncallwithvals))
				} else {
					RPCBOD.Body = append(RPCBOD.Body, Create_Method("for valg", forchar1, foreqchar, forchar1, forlenchar, argumentcolor))
				}
			}

		}
	}
}
