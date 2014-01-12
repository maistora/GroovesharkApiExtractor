package pageUtil

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
	"../structs"
	"../groupRegexp"
)

func ExtractMethodsDiv(page string) string {
	groupName := "div"
	regex := `(?s).*<div class="main_content_panes">(?P<` + groupName + `>.*)</div>.*`

	groupRegex := groupRegexp.NewGroupRegexp(regex, page)

	methodDiv := groupRegex.GetGroup(groupName)
	return methodDiv
}

func ExtractMethodURLsAndTitles(page string) *[]structs.MethodExtraction {
	regex := regexp.MustCompile(`(<a href="(.*)".*</a>)|(<h3>(.*)</h3>)`)
	methodsAndSections := removeOverviewSection(regex.FindAllString(page, -1))
	fmt.Println(methodsAndSections)

	// methodUrlGroup := "methodUrl"
	// sectionGroup := "section"
	// regex := `(<a href="(?P<` + methodUrlGroup + `>.*)".*</a>)|(<h3>(?P<` + sectionGroup + `>.*)</h3>)`

	// groupRegex := groupRegexp.NewGroupRegexp(regex, page)

	// methodUrl := groupRegex.GetGroup(methodUrlGroup)
	// section := groupRegex.GetGroup(sectionGroup)

	// fmt.Println(methodUrl)
	// fmt.Println(section)

	return nil
}

func removeOverviewSection(methodsAndSections []string) []string {
	return methodsAndSections[1:]
}

func GetPageAsString(url string) string {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	respBodyStr := getRespBodyAsString(response)
	return respBodyStr
}

func getRespBodyAsString(response *http.Response) string {
	if response.StatusCode != 200 { 
		return ""
	}
	bodyBytes, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		panic(err2)
	}
	bodyString := string(bodyBytes) 
	return bodyString
}

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

const PARAM_TABLE_COLS = 3

func GetFuncParams(page string) *[]structs.FuncParam {
	regex := regexp.MustCompile(`<td>.*</td>`)
	allParams := regex.FindAllString(page, -1)
	params := make([]structs.FuncParam, 0, 0)
	param := new(structs.FuncParam)
	for i, paramStr := range allParams {
		if newParamIsReached(i) {
			param = new(structs.FuncParam)
		}
		switch i % PARAM_TABLE_COLS {
		case 0:
			param.Name = extractName(paramStr)
		case 1:
			param.ParamType = extractParamType(paramStr)
		case 2:
			param.Required = extractRequired(paramStr)
		}
		if paramIsFilled(i) { 
			params = append(params, *param)
		}
	}
	return &params
}

func newParamIsReached(count int) bool {
	return count % PARAM_TABLE_COLS == 0
}

func paramIsFilled(count int) bool {
	return count % PARAM_TABLE_COLS == 2
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
