package rpc

type ParamDoc struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
}

type MethodDoc struct {
	Method      Method     `json:"method"`
	Params      []ParamDoc `json:"params,omitempty"`
	Description string     `json:"description,omitempty"`
}

type ApiDoc struct {
	Methods     []MethodDoc `json:"methods"`
	Description string      `json:"description,omitempty"`
	Version     string      `json:"version"`
}
