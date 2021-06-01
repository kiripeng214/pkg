package xrpc

import (
	"kiripeng214/pkg/xrpc/codec"
)

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber int
	CodeType    codec.Type
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodeType:    codec.GobType,
}

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

//https://github.com/geektutu/7days-golang/blob/master/gee-rpc/day1-codec/server.go
//func (server Server) ServeConn(conn io.ReadWriteCloser)  {
//	defer conn.Close()
//	var opt Option
//	if err := json.NewDecoder()
//}
