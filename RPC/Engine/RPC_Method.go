package RPC

import "fmt"

func Create_Method(method_name string, arguments ...string) string {
	switch method_name {
	case "for range":
		template := `for %s, %s := range %s { fmt.Println(%s) }`
		template = fmt.Sprintf(template, arguments[0], arguments[1], arguments[2], arguments[3])
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
