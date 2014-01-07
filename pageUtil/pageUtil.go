package pageUtil

import (
	"fmt"
	"regexp"
	"strings"
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

const COL_NUM = 3

func GetFuncParams(page string) *[]structs.FuncParam {
	// TODO serious refactor
	regex := regexp.MustCompile(`<td>.*</td>`)
	allParams := regex.FindAllString(page, -1)
	params := make([]structs.FuncParam, 1, 1)
	param := new(structs.FuncParam)
	requiredPrefix := "<td><strong>"
	typePrefix := "<td><em>"
	namePrefix := "<td>"
	for i, paramStr := range allParams {
		if i % COL_NUM == 0 {
			param = new(structs.FuncParam)
		}
		if strings.Contains(paramStr, requiredPrefix) {
			param.Required = paramStr[len(requiredPrefix):len(paramStr) - (len(requiredPrefix) + 2)] // 2 closing tags
		} else if strings.Contains(paramStr, typePrefix) {
			param.ParamType = paramStr[len(typePrefix):len(paramStr) - (len(typePrefix) + 2)] // 2 closing tags
		} else if strings.Contains(paramStr, namePrefix) {
			param.Name = paramStr[len(namePrefix):len(paramStr) - (len(namePrefix) + 1)] // 1 closing tags
		}
		if i % COL_NUM == 0 {
			params = append(params, *param)
		}
	}

	fmt.Println(params)
	return &params
}