package pageUtil

import (
	"fmt"
	"../structs"
	"../groupRegexp"
)

func GetFuncName(page string) string {
	groupName := `funcName`
	regex := `.*<h3.*>(?P<` + groupName + `>.*)</h3>.*`

	groupRegex := groupRegexp.NewGroupRegexp(regex, page)

	funcName := groupRegex.GetGroup(groupName)
	return funcName
}

func GetFuncDoc(page string) string {
	doc := `doc`
	note := `note`

	regex := `(?s).*<p>(?P<` + doc + `>.*)</p>.*<p><em>(?P<` + note + `>.*)</em></p>.*`
	groupRegex := groupRegexp.NewGroupRegexp(regex, page)

	funcDoc := groupRegex.GetGroup(doc) + "\nNOTE: " + groupRegex.GetGroup(note)
	return funcDoc
}

func GetFuncParams(page string) *[]structs.FuncParam {
	param1 := new(structs.FuncParam)
	param1.Name = "param1"
	param1.ParamType = "Integer"
	param1.Required = true

	param2 := new(structs.FuncParam)
	param2.Name = "param2"
	param2.ParamType = "Boolean"
	param2.Required = true
	params := []structs.FuncParam {*param1, *param2}
	return &params
}