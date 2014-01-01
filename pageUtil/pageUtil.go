package pageUtil

import (
	"../structs"
)

func GetFuncName(page string) string {
	return ""
}

func GetFuncDoc(page string) string {
	return ""
}

func GetFuncParams(page string) *[]structs.FuncParam {
	param1 := new(structs.FuncParam)
	param2 := new(structs.FuncParam)
	params := []structs.FuncParam {*param1, *param2}
	return &params
}