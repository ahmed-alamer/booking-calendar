package rpc

type Method string
type Params string
type ErrorCode int

const (
	ParseError               = -32700
	InvalidRequest           = -32600
	MethodNotFound           = -32601
	InvalidParams            = -32602
	InternalError            = -32603
	ProviderNotFound         = -32001
	RequestMissingProviderId = -32002
)

type Request struct {
	Id     string `json:"id,omitempty"`
	Method Method `json:"method"`
	Params Params `json:"params"` // This will be parsed by the method
}

type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Data    string    `json:"data,omitempty"`
}

type Response struct {
	Id     string      `json:"id"`
	Result interface{} `json:"result"`
	Error  Error       `json:"error,omitempty"`
}
