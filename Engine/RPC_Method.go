package RPC

import (
	"fmt"
)

func Create_Method(method_name string, arguments ...string) string {
	switch method_name {
	case "for range":
		template := `for %s, %s := range %s { fmt.Println(%s) }`
		template = fmt.Sprintf(template, arguments[0], arguments[1], arguments[2], arguments[3])
		return template
	case "for valg":
		template := `for %s := %s; %s < %s; %s++{ fmt.Println(%s)} `
		// arg 0 | Example = i = %s1
		// arg 1 | Example = 0 = %s2
		// arg 2 | Example = i = %s3
		// arg 3 | Example = RPC_PRE.Values
		// arg 4 | Example = i
		// arg 5 | Example = Generate_LIA_Tag(RPC_PRE.Values[i]) or => XQ : RPC::Prepare::LI:A =RPC_PRE.Values[i]=
		// translates to => for i := 0; i < len(RPC_PRE.Values); i++ {}
		template = fmt.Sprintf(template,
			arguments[0], // Example = i
			arguments[1], // Example = 0
			arguments[0], // Example = i
			arguments[3], // Example = len(RPC_PRE.Values);
			arguments[0], // Example = i
			arguments[4], // Example = function call
		)

		return template
	case "function":
		template := arguments[0]
		return template
	case "functiondef":
		template := `func %s() { %s }`
		template = fmt.Sprintf(template, arguments[0], arguments[1])
		return template
	}
	return ""
}
