package pageUtil

import (
	"net/http"
	"io/ioutil"
	"regexp"
	"fmt"
	"strings"
	"../structs"
	"../groupRegexp"
)

const GS_HOME = "http://developers.grooveshark.com"

func ExtractMethodsDiv(page string) string {
	groupName := "div"
	regex := `(?s)(?U).*<div class="main_content_panes">(?P<` + groupName + `>.*)</div>.*`

	groupRegex := groupRegexp.NewGroupRegexp(regex, page)

	methodDiv := groupRegex.GetGroup(groupName)
	return methodDiv
}

func ExtractMethodURLsAndTitles(page string) *[]structs.MethodUrlExtraction {
	extracts := make([]structs.MethodUrlExtraction, 0, 0)
	regex := regexp.MustCompile(`(<a.*class="method".*</a>)|(<h3>(.*)</h3>)`)
	methodsAndSections := removeOverviewSection(regex.FindAllString(page, -1))
	for _, selection := range methodsAndSections {
		extract := getMethodOrSectionExtraction(selection)
		extracts = append(extracts, *extract)
	}
	return &extracts
}

func getMethodOrSectionExtraction(selection string) *structs.MethodUrlExtraction {
	methodUrlGroup := "methodUrl"
	sectionGroup := "section"
	regex := `((?U)<a\shref="(?P<` + methodUrlGroup + `>.*)".*</a>)|(<h3>(?P<` + sectionGroup + `>.*)</h3>)`

	groupRegex := groupRegexp.NewGroupRegexp(regex, selection)

	methodUrl := groupRegex.GetGroup(methodUrlGroup)
	section := groupRegex.GetGroup(sectionGroup)

	return buildExtract(methodUrl, section)
}

func buildExtract(methodUrl, section string) *structs.MethodUrlExtraction {
	extract := new(structs.MethodUrlExtraction)
	if methodUrl != "" {
		extract.Text = GS_HOME + methodUrl
		extract.IsURL = true
	} else {
		extract.Text = section
		extract.IsURL = false
	}

	return extract
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
	funcDoc := getDoc(page)
	funcNote := getNote(page)

	doc := ""
	if len(funcDoc) != 0 {
		doc += fmt.Sprintf("// %v\n", funcDoc)
	}
	if len(funcNote) != 0 {
		doc += fmt.Sprintf("// Note: %v\n", funcNote)
	}

	return doc
}

func getDoc(page string) string {
	doc := `doc`
	regexDoc := `(?s)(?U).*<p>(?P<` + doc + `>.*)</p>.*`
	docRegex := groupRegexp.NewGroupRegexp(regexDoc, page)

	return strings.Trim(docRegex.GetGroup(doc), "")
}

func getNote(page string) string {
	note := `note`
	regexNote := `(?s)(?U).*<p><em>(?P<` + note + `>.*)</em></p>.*`
	noteRegex := groupRegexp.NewGroupRegexp(regexNote, page)

	return strings.Trim(noteRegex.GetGroup(note), " ")
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
