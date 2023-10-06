package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("p", 8000, "port")
	flag.Parse()
	fmt.Printf("Starting http server on port %d\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			bodyBytes = []byte(fmt.Sprintf("!{error reading body: %s}", err.Error()))
		}
		fmt.Printf("%s %s %s %v %s\n", r.Method, r.URL, r.Proto, r.Header, bodyBytes)
	}))
	if err != nil {
		log.Fatal(err)
	}
}
