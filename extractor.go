package main

import (
	"fmt"
	"bytes"
	"./myIoUtil"
	"./pageUtil"
	"./structs"
)

var FILE_NAME = "groovesharkApi.go"

func main() {
	page := pageUtil.GetPageAsString("http://developers.grooveshark.com/docs/public_api/v3")
	mainDiv := pageUtil.ExtractMethodsDiv(page)
	ulrsOrSections := pageUtil.ExtractMethodURLsAndTitles(mainDiv)
	createApiGoFile(*ulrsOrSections)
}

func createApiGoFile(urlsOrSections []structs.MethodUrlExtraction) {
	initFileText := getInitFileText()
    myIoUtil.AppendTextToFile(initFileText, FILE_NAME)
    var buffer bytes.Buffer
	for _, urlOrSect := range urlsOrSections {
		fmt.Print(".")
		if urlOrSect.IsURL {
			buffer.WriteString(getFunctionFromUrl(urlOrSect.Text))
		} else {
			buffer.WriteString(getSection(urlOrSect.Text))
		}
	}
	fmt.Println("")
	myIoUtil.AppendTextToFile(buffer.String(), FILE_NAME)
}

func getSection(section string) string {
	return fmt.Sprintf("\n// ================= %v =================\n", section)
}

func getFunctionFromUrl(url string) string {
	funcPage := pageUtil.GetPageAsString(url)
	funcProps := extractFuncProperties(funcPage)
	funcAsText := populateFuncTemplate(funcProps)
	return funcAsText
}

func extractFuncProperties(funcPage string) *structs.FuncProperties {
	funcProps := new(structs.FuncProperties)
	funcProps.Name = pageUtil.GetFuncName(funcPage)
	funcProps.Doc = pageUtil.GetFuncDoc(funcPage)
	funcProps.Params = pageUtil.GetFuncParams(funcPage)
	return funcProps
}

func populateFuncTemplate(props *structs.FuncProperties) string {
	funcName := props.Name
	params := buildFuncParams(*props.Params)
	return fmt.Sprintf("\n%vfunc %v(%v) {\n\t//TODO impelemnt\n}\n", props.Doc, funcName, params)
}

func buildFuncParams(params []structs.FuncParam) string {
	if params == nil || len(params) == 0 {
		return ""
	}
	result := ""
	for i, param := range params {
		if i < (len(params) - 1) && params[i + 1].ParamType == param.ParamType {
			result += param.Name + ", "
		} else {
			result += param.Name + " " + param.ParamType + ", "
		}
	}
	result = result[:len(result) - 2] // for the last comma
	return result
}

func getInitFileText() string {
	return "// This file was auto-generated from the\n// Grooveshark API Extractor\npackage main\n\n"
}