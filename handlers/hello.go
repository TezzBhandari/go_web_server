package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	log *log.Logger
}

func NewHello(log *log.Logger) *Hello {
	return &Hello{
		log: log,
	}
}

func (hello *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	hello.log.Println("Handling Hello Request")

	// reads the request body
	hello.log.Printf("%#v", r.Body)
	d, err := io.ReadAll(r.Body)
	if err != nil {
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Bad Request"))

		http.Error(rw, "Bad Rquest", http.StatusBadRequest)
		return
	}

	// rw.Write([]byte("Hello"))
	fmt.Fprintf(rw, "Hello %s", d)
}
