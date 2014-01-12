package main

import (
	"fmt"
	"strings"
	"./myIoUtil"
	"./pageUtil"
	"./structs"
)

var FILE_NAME = "groovesharkApi.go"

func main() {
	page := pageUtil.GetPageAsString("http://developers.grooveshark.com/docs/public_api/v3")
	mainDiv := pageUtil.ExtractMethodsDiv(page)
	ulrs := getAllUrlsToFunctions(mainDiv)
	createApiGoFile(ulrs)
}

func getAllUrlsToFunctions(mainDiv string) []string {
	ulrs := make([]string, 0, 0)
	pageUtil.ExtractMethodURLsAndTitles(mainDiv)
	// ulrs = append(ulrs, "http://developers.grooveshark.com/docs/public_api/v3/method.php?method=addUserLibrarySongs")
	// ulrs = append(ulrs, "http://developers.grooveshark.com/docs/public_api/v3/method.php?method=getUserLibrarySongs")
	return ulrs 
}

func createApiGoFile(ulrs []string) {
	initFileText := getInitFileText()
    myIoUtil.AppendTextToFile(initFileText, FILE_NAME)
	for _, ulr := range ulrs {
		funcPage := pageUtil.GetPageAsString(ulr)
		funcProps := extractFuncProperties(funcPage)
		funcAsText := populateFuncTemplate(funcProps)
		myIoUtil.AppendTextToFile(funcAsText, FILE_NAME)
	}
}

func extractFuncProperties(funcPage string) *structs.FuncProperties {
	funcProps := new(structs.FuncProperties)
	funcProps.Name = pageUtil.GetFuncName(funcPage)
	funcProps.Doc = pageUtil.GetFuncDoc(funcPage)
	funcProps.Params = pageUtil.GetFuncParams(funcPage)
	return funcProps
}

func populateFuncTemplate(props *structs.FuncProperties) string {
	doc := strings.Replace(props.Doc, "\n", "\n//", -1)
	funcName := props.Name
	params := buildFuncParams(*props.Params)
	return fmt.Sprintf("\n//%v\nfunc %v(%v) {\n\t//TODO impelemnt\n}\n", doc, funcName, params)
}

func buildFuncParams(params []structs.FuncParam) string {
	if params == nil {
		return ""
	}
	result := ""
	for _, param := range params {
		result += param.Name + " " + param.ParamType + ", "
	}
	result = result[:len(result) - 2] // for the last comma
	return result
}

func getInitFileText() string {
	return "// This file was auto-generated from the\n// Grooveshark API Extractor\npackage main\n\n"
}