package RPC

import (
	"fmt"
	"strings"
)

var RPC_Notes = []string{
	"!<<",
	"!>>",
	"!!!-> rpcn",
	"!!!<-",
}

func Interpreter() {
	for p := 0; p < len(RPCC.CodeBlock); p++ {
		//
		//
		//
		// Find mods by macro block
		libs := LoadingVariablesFrom.FindStringSubmatch(RPCC.CodeBlock[p])
		if len(libs) > 1 {
			TemplateData.Import_name = libs[1]
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
		/*
			ft := Generation_Methods[strings.Trim(RPCC.CodeBlock[p], " ")]
			if ft != nil {
				ft.(func(...string))(Argument[1])
				fmt.Println(Strchar.TagData)
			}
		*/
	}
	RPCBOD.Generate_Template()
}
