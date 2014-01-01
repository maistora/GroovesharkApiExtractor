package main

import (
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
	"./myIoUtil"
	"./pageUtil"
	"./structs"
)

var FILE_NAME = "groovesharkApi.go"

func main() {
	page := getPageAsString("http://google.bg")
	mainDiv := extractDivWithFunctions(page, "class/id")
	links := getAllLinksToFunctions(mainDiv)
	createApiGoFile(links)
}

func getPageAsString(url string) string {
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

func extractDivWithFunctions(page, classifier string) string {
	mainDiv := " "
	return mainDiv
}

func getAllLinksToFunctions(mainDiv string) []string {
	links := make([]string, 0, 0)
	links = append(links, "http://google.bg")
	links = append(links, "http://google.bg")
	return links 
}

func createApiGoFile(links []string) {
	plainGoFileStr := getPlainGoFile()
    myIoUtil.AppendTextToFile(plainGoFileStr, FILE_NAME)
	for _, link := range links {
		fmt.Println(link)
		funcPage := getPageAsString(link)
		funcProps := extractFuncProperties(funcPage)
		funcString := populateFuncTemplate(funcProps)
		myIoUtil.AppendTextToFile(funcString, FILE_NAME)
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
	return fmt.Sprintf("\n//%v\nfunc %v(%v) {\n}\n", doc, funcName, params)
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

func getPlainGoFile() string {
	return "package main\n\n"
}