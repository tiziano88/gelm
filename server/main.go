package main

import (
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"github.com/gorilla/pat"

	pb "github.com/tiziano88/gelm/server/proto"
)

var marshaler = &jsonpb.Marshaler{
	EnumsAsInts:  false,
	EmitDefaults: true,
	Indent:       "  ",
	OrigName:     true,
}

func main() {
	m := pat.New()

	m.Get("/", Handler)

	http.Handle("/", m)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	m := &pb.Message{}
	marshaler.Marshal(w, m)
}
