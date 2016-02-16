package main

import (
	"log"
	"net/http"
	"time"

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
	log.Print("starting server 33x")

	m := pat.New()

	m.Get("/api/", Handler)
	m.Get("/", RootHandler)

	http.Handle("/", m)

	s := &http.Server{
		Addr:         "0.0.0.0:1234",
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	s.SetKeepAlivesEnabled(false)
	s.ListenAndServe()
	s.SetKeepAlivesEnabled(false)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	m := &pb.Message{
		Id:                123,
		FieldWithLongName: "hello1",
	}
	log.Printf("v: %v", m)
	marshaler.Marshal(w, m)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/out/index.html")
}
