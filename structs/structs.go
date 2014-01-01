package structs

type FuncProperties struct {
	Name string
	Doc string
	Params *[]FuncParam
}

type FuncParam struct {
	Name string
	ParamType string
	Required bool
}
