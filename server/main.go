package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/gorilla/pat"

	pb "github.com/tiziano88/gelm/server/proto"
)

var marshaler = &jsonpb.Marshaler{
	Indent:   "  ",
	OrigName: false,
}

func main() {
	log.Print("starting server")

	m := pat.New()

	m.Get("/api/", Handler)
	// m.Get("/", RootHandler)
	m.Get("/", ProxyHandler)

	http.Handle("/", m)

	s := &http.Server{
		Addr: "0.0.0.0:1234",
	}
	s.ListenAndServe()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	m := &pb.Message{
		FieldWithLongName: "hello1",
		Enum:              pb.Enum_ENUM_VALUE_2,
		SubMessage: &pb.SubMessage{
			Id: 222,
		},
		BoolField: true,
		User: &pb.User{
			Name: "name",
			Address: &pb.Address{
				Line_1: "line1",
				City:   "London",
			},
		},
	}
	log.Printf("v: %v", m)
	marshaler.Marshal(w, m)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/out/index.html")
}

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	r.Host = "localhost:8000"
	r.Proto = "HTTP"
	r.RequestURI = ""
	r.URL.Scheme = "http"
	r.URL.Host = "localhost:8000"
	r.URL.Path = strings.Replace(r.URL.Path, ".html", ".elm", -1)
	log.Printf("request: %#v", r)
	log.Printf("url: %#v", *r.URL)

	cl := &http.Client{}
	resp, err := cl.Do(r)
	if err != nil {
		log.Printf("Error proxying: %v", err)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	for h := range resp.Header {
		w.Header().Add(h, resp.Header.Get(h))
	}
	w.Write(body)
}
