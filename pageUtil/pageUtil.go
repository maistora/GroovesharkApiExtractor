package pageUtil

import (
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
			param.Name = extractName(paramStr)
		case 1:
			param.ParamType = extractParamType(paramStr)
		case 2:
			param.Required = extractRequired(paramStr)
		}

		// append only the fully filled param
		if i % COL_NUM == 2 { 
			params = append(params, *param)
		}
	}
	return &params
}

func cleanTags(text string, startTagsLen, tagNum int) string {
	return text[startTagsLen:len(text) - (startTagsLen + tagNum)]	
}

func extractName(text string) string {
	return cleanTags(text, 4, 1)
}

func extractParamType(text string) string {
	return cleanTags(text, 8, 2)
}

func extractRequired(text string) string {
	return cleanTags(text, 12, 2)
}
