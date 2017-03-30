package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const defaultPort int = 80

func main() {
	var err error
	port := defaultPort
	if len(os.Args) > 1 {
		rawPort := os.Args[1]
		port, err = strconv.Atoi(rawPort)
		if err != nil {
			log.Fatalf("could not parse port \"%s\": %s", rawPort, err)
		}
	}

	log.Printf("listening on port %d", port)

	http.HandleFunc("/", echoHandler)

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		fallthrough
	case "GET":
		handleGet(w, r)
	case "PUT":
		fallthrough
	case "POST":
		handlePost(w, r)
	default:
		fmt.Fprintf(w, "%s unimplemented", r.Method)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	method := fmt.Sprintf("Method:\n%s", r.Method)
	host := fmt.Sprintf("Host:\n%s", r.Host)
	header := fmt.Sprintf("Header:\n%v", r.Header)

	strs := []string{method, host, header}
	resp := strings.Join(strs, "\n\n")

	log.Print("\n" + resp)
	fmt.Fprint(w, resp)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	method := fmt.Sprintf("Method:\n%s", r.Method)
	host := fmt.Sprintf("Host:\n%s", r.Host)
	header := fmt.Sprintf("Header:\n%v", r.Header)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}
	body := fmt.Sprintf("Body:\n%s", string(b))

	strs := []string{method, host, header, body}
	resp := strings.Join(strs, "\n\n")

	log.Print("\n" + resp)
	fmt.Fprint(w, resp)
}
