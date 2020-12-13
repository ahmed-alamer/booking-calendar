package rpc

type Processor interface {
	Execute(request Request) Response
}
