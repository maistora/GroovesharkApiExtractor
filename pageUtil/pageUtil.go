package pageUtil

import (
	"fmt"
	"regexp"
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
	regex := regexp.MustCompile(`<td>.*</td>`)
	allParams := regex.FindAllString(page, -1)
	params := make([]structs.FuncParam, 0, 0)
	param := new(structs.FuncParam)
	for i, paramStr := range allParams {
		if i % COL_NUM == 0 {
			param = new(structs.FuncParam)
		}

		switch i % COL_NUM {
		case 0:
			param.Name = cleanTags(paramStr, 4, 1)
		case 1:
			param.ParamType = cleanTags(paramStr, 8, 2)
		case 2:
			param.Required = cleanTags(paramStr, 12, 2)
		}

		// TODO fix this - it does not persist params after appending the value to the list
		fmt.Println(*param)
		if i % COL_NUM == 0 {
			params = append(params, *param)
		}
	}

	// fmt.Println((&params[0]).ParamType)
	return &params
}

func cleanTags(text string, startTagsLen, tagNum int) string {
	return text[startTagsLen:len(text) - (startTagsLen + tagNum)]	
}