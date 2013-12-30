package main

import (
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
	"./myIoUtil"
)

var FILE_NAME = "groovesharkApi.go"

type FuncProperties struct {
	name string
	doc string
	params *[]FuncParam
}

type FuncParam struct {
	name string
	paramType string
	required bool
}

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

func extractFuncProperties(funcPage string) *FuncProperties{
	funcProps := new(FuncProperties)
	return funcProps
}

func populateFuncTemplate(props *FuncProperties) string {
	// TODO format the doc as it can be on multiple rows
	// TODO format params as they can be many. Check if they are all same type
	// and if yes then put the type at the end. Can try with groups of types.
	doc := strings.Replace(props.doc, "\n", "\n//", -1)
	funcName := props.name
	params := buildFuncParams(props.params)
	return fmt.Sprintf("\n//%v\nfunc %v(%v) {\n}\n", doc, funcName, params)
}

func buildFuncParams(params *[]FuncParam) string {
	// TODO add implementation
	return "testParams int"	
}

func getPlainGoFile() string {
	return "package main\n\n"
}