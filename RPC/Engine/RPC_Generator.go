package RPC

import "fmt"

func (Tags *Tagged) Generate_LIA_Tag(values ...[]string) {
	var lias string
	for index, val := range values {
		lias += fmt.Sprintf("<li><a>%v</li></a>", val[index])
	}
	Tags.TagData = append(Tags.TagData, string(lias))
}
