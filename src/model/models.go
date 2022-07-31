package model

type Program struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Nodes       string `json:"nodes"`
	Drawflow    string `json:"drawflow"`
	Code        string `json:"code"`
}

type ProgramInfo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CodePython struct {
	Code string `json:"code"`
}
