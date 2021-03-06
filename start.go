package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
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
		handleNoBody(w, r)
	case "PATCH":
		fallthrough
	case "PUT":
		fallthrough
	case "POST":
		handleWithBody(w, r)
	default:
		fmt.Fprintf(w, "%s unimplemented", r.Method)
	}
}

func handleNoBody(w http.ResponseWriter, r *http.Request) {
	commonFields := commonParse(w, r)
	resp := strings.Join(commonFields, "\n\n")

	log.Print("\n" + resp)
	fmt.Fprint(w, resp)
}

func handleWithBody(w http.ResponseWriter, r *http.Request) {
	commonFields := commonParse(w, r)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}
	body := fmt.Sprintf("Body:\n%s", string(b))

	strs := append(commonFields, body)
	resp := strings.Join(strs, "\n\n")

	log.Print("\n" + resp)
	fmt.Fprint(w, resp)
}

func commonParse(w http.ResponseWriter, r *http.Request) []string {
	method := fmt.Sprintf("Method:\n%s", r.Method)
	host := fmt.Sprintf("Host:\n%s", r.Host)
	path := fmt.Sprintf("Path:\n%s", r.URL.Path)

	var headerKeys []string
	header := "Header:"
	for k, _ := range r.Header {
		headerKeys = append(headerKeys, k)
	}
	sort.Strings(headerKeys)
	for _, k := range headerKeys {
		header += fmt.Sprintf("\n%s=%s", k, strings.Join(r.Header[k], ","))
	}

	return []string{method, host, path, header}
}
