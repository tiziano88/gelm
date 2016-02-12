package server

import (
	"github.com/golang/protobuf/jsonpb"
)

var marshaler = &jsonpb.Marshaler{
	EnumAsInts:   false,
	EmitDefaults: true,
	Indent:       "  ",
	OrigName:     true,
}

func main() {
}
