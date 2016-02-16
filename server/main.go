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
	OrigName:     false,
}

func main() {
	m := pat.New()

	m.Get("/", Handler)

	http.Handle("/", m)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	m := &pb.Message{
		FieldWithLongName: "hello",
	}
	marshaler.Marshal(w, m)
}
