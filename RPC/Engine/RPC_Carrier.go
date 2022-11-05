package RPC

import "regexp"

/*

RPC stands for Radical Processing Language, It is a language concept with a golang engine which allows you to create certain elements within other languages

RPC also has its own subset but for right now RPC will be built using the Go programming language and will be used with a custom app built around engine in frizz

for dynamic HTML generation, please see https://github.com/org/Scare-Security/RPC for more information



RPC GENERAL


init RPC templating

    {
        (
            [
                RPC -> RPCL6_Count | RPC.BRICK => {{{
                    RPC:::Unit::SYMBOL | RPC.BRICK START

                      !!!-> rpcn
                                This is where the HTML code will be renedered by the engine
                      !!!<-

                    RPC::Unit::SYMBOL | RPC.BRICK END
                }}}
            ]
        )
    }

end RPC templating

you will embed RPC's templating language in HTML and run that HTML template through the engine, the engine will go through the entire file and use a regular expression to

find a given statement then prepare or predict one. for example if you have the block

init RPC templating

    {
        (
            [

it will predict that the next line is RPC counting symbol, which RPC -> RPCL must have after L the verified line of the file, if this is not accurate the engine will

spit out the error or assume that you have fixed it
*/

type Tagged struct {
	TagData []string
}

// START SYMBOL PARSING

var (
	line                 int
	lines                int  = 0
	SymbolRegex               = "(?i)RPC::Unit::SYMBOL__RPC.BRICK::END(.*)"
	SymbolRegexuse            = regexp.MustCompile(SymbolRegex)
	UnitSymbolFoundBegin bool = false
)

// This is a very weird way to do this but it works for now, so....
var times int = 0

var Symbols = map[string]string{
	"init RPC templating":                 "Init",     // Prepare templating
	"end RPC templating":                  "End Init", // End Template generation section
	"RPC::Unit::SYMBOL | RPC.BRICK START": "STG",      // Start Template Generation
	"!!!-> rpcn":                          "IGNORE",   // RPC Note [BEGIN]
	"!!!<-":                               "IGNORE",   // RPC Note [END]
}

var Engine_Error_Messages = map[string]string{
	"File Missing":        "EC1",
	"File Not Correct":    "EC2",
	"File Invalid":        "EC3",
	"File Not acceptable": "EC4",
	"Initation did not have a valid end of call":   "EC5",
	"Initation did not have a valid start of call": "EC6",
}

var File_Extensions_Possible = map[string]string{
	".rpct":   "RPC template file HTML",
	".rpc":    "RPC source code file",
	".rpc++":  "RPC class file",
	".rpch++": "RPC C++, C, Lua source code file",
}

type RPC_Data struct {
	Package_name  string
	Import_name   string
	Body          []string
	BodyFunctions []string
}

var Template = `
	package %s // package name @{}@

	import (
		RPC_PRE %s // Importing from the !{}! macro
	)

`

var Banner = `
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃     ___       ___       ___    ┃
┃    /\  \     /\  \     /\  \   ┃
┃   /::\  \   /::\  \   /::\  \  ┃
┃  /::\:\__\ /::\:\__\ /:/\:\__\ ┃
┃  \;:::/  / \/\::/  / \:\ \/__/ ┃
┃   |:\/__/     \/__/   \:\__\   ┃
┃    \|__|               \/__/   ┃
┃   Radical Processing Language  ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
`

// we want to tell the engine to ignore all code notes
var RPCC RPCInterpretation
var notelist bool
var Args Arguments
var Strchar Tagged
var TemplateData RPC_Data

var (
	LoadingVariablesFrom   = regexp.MustCompile(`!{(.*)}!`)
	LoadingPackagesAs      = regexp.MustCompile(`@{(.*)}@`)
	RPC_Generate_Arguments = regexp.MustCompile(`=(.*)=`)
	RPC_Generate_FOR       = regexp.MustCompile(`(i?)for(.*)`)
	RPC_Generate_LIA       = regexp.MustCompile(`(i?)RPC::Prepare::LI:A(.*)`)
	RPCBOD                 RPC_Data
)

var Conditions = map[string]string{
	"RPCIF<<":   "RPC if condition",
	"RPCEIF<<":  "RPC else if condition",
	"RPCELSE<<": "RPC else condition",
}

var RPC_Symbols = map[string]bool{
	"RPC::Generate::LI::A": true,
}

var LIA_Template = `
func Generate_LIA_Tag(values ...string) {
	var lias string
	for index, val := range values {
		lias += fmt.Sprintf("<li><a>%v</li></a>", val[index])
	}
	return string(lias)
}
`

type Arguments struct {
	Strargs string
}

var Generation_Methods = map[string]interface{}{ // Based on a found method call a function
	"RPC::Generate::LI::A": Strchar.Generate_LIA_Tag,
}

type RPCInterpretation struct {
	CodeBlock []string
}

var (
	EXTERN_FILE_MACRO       = "(i?)@EXTERNF >"
	EXTERN_FILE_MACRO_USE   = regexp.MustCompile(EXTERN_FILE_MACRO)
	argumentcolor           string
	Generate_Under_For_Loop bool
)
