package RPC

import (
	"fmt"
)

// Now we generate all methods based on the data in the template file

func (RPC *RPC_Data) Generate_Template(importarray, modulesarray []string) {
	Template = fmt.Sprintf(Template, TemplateData.Package_name)
	Template += "\n\nimport (\n"
	for q := 0; q < len(importarray); q++ {
		Template += "\t" + importarray[q]
		Template += "\n"
	}
	for a := 0; a < len(modulesarray); a++ {
		Template += "\tRPC_PRE " + modulesarray[a]
		Template += "\n"
	}
	Template += ")\n\n"
	for k := 0; k < len(RPC.BodyFunctions); k++ {
		Template += RPC.BodyFunctions[k]
	}
	Template += "func main() {"
	for i := 0; i < len(RPC.Body); i++ {
		Template += RPC.Body[i]
	}
	Template += "}"
	Write_Frame("body.go", Template)
}
