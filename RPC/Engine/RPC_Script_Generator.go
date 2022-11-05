package RPC

import "fmt"

// Now we generate all methods based on the data in the template file

func (RPC *RPC_Data) Generate_Template() {
	Template = fmt.Sprintf(Template, TemplateData.Package_name, TemplateData.Import_name)
	for k := 0; k < len(RPC.BodyFunctions); k++ {
		Template += RPC.BodyFunctions[k]
	}
	Template += "func main() {"
	for i := 0; i < len(RPC.Body); i++ {
		Template += RPC.Body[i]
	}
	Template += "}"
	Write_Frame("body.txt", Template)
}
