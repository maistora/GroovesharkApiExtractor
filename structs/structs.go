package structs

type FuncProperties struct {
	Name string
	Doc string
	Params *[]FuncParam
}

type FuncParam struct {
	Name string
	ParamType string
	Required string
}

type MethodUrlExtraction struct {
	Text string
	IsURL bool
}