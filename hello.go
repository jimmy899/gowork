package main

import "net/http"
import "fmt"
import "crypto/tls"

type myHTTPHandler struct{}

func (h *myHTTPHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello\n")
}

func main() {

	cert, err := tls.LoadX509KeyPair("fullchain.pem", "privkey.pem")
	if err != nil {
		return
	}
	var nextProtos []string
	enableHTTP2 := false

	myMux := http.NewServeMux()

	myMux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "/")
	})

	if enableHTTP2 {
		nextProtos = []string{"h2"}
	} else {
		nextProtos = nil
	}

	tc := tls.Config{
		ServerName:   "localhost",
		Certificates: []tls.Certificate{cert},
		NextProtos:   nextProtos,
	}

	if ln, err := tls.Listen("tcp", ":443", &tc); err == nil {
		s := http.Server{
			Handler: myMux,
		}

		s.Serve(ln)
	} else {
		fmt.Printf("%v\n", err)
	}
}
