package rpc

type Api interface {
	Execute(request Request) Response
}
