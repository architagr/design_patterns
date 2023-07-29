package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type MyServer struct {
}

func (server *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "called My Server handler\n")
}

type MyLoggerServer struct {
	existingHandler http.Handler
	logWriter       io.Writer
}

func (server *MyLoggerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(server.logWriter, "-----------------------------------------\n")
	fmt.Fprintf(server.logWriter, "request URI %s \n", r.RequestURI)
	fmt.Fprintf(server.logWriter, "-----------------------------------------\n")
	fmt.Fprintf(w, "called logger Server handler\n")

	server.existingHandler.ServeHTTP(w, r)
}

type MyAuthServer struct {
	existingHandler    http.Handler
	username, password string
}

func (server *MyAuthServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if ok {
		if user == server.username && pass == server.password {
			fmt.Fprintf(w, "Auth successfull\n")
			server.existingHandler.ServeHTTP(w, r)
		} else {
			fmt.Fprintf(w, "username and password incorrect\n")
		}
	} else {
		fmt.Fprintf(w, "error in getting basic auth data\n")
	}
}

func main() {

	// initilize base server
	var myFinalServer http.Handler = &MyServer{}
	// adding logger on the base server
	myFinalServer = &MyLoggerServer{
		existingHandler: myFinalServer,
		logWriter:       os.Stdout,
	}
	// adding auth server on the existing server
	myFinalServer = &MyAuthServer{
		existingHandler: myFinalServer,
		username:        "admin",
		password:        "admin",
	}

	http.Handle("/", myFinalServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
