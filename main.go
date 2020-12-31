package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

func main() {
	portNumber := flag.Int("port", 8080, "Port number to listen on.")

	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Protocol: %s\n", r.Proto)
		fmt.Printf("Host: %s\n", r.Host)
		fmt.Printf("Method: %s\n", r.Method)
		fmt.Printf("Path: %s\n", r.URL.Path)

		headerNames := make([]string, 0, len(r.Header))
		for headerName := range r.Header {
			headerNames = append(headerNames, headerName)
		}
		sort.Strings(headerNames)
		fmt.Printf("Headers: (%d)\n", len(r.Header))
		for _, headerName := range headerNames {
			for _, value := range r.Header[headerName] {
				fmt.Printf("* %s: %s\n", headerName, value)
			}
		}

		contents, _ := ioutil.ReadAll(r.Body)

		fmt.Printf("Body:\n")
		fmt.Printf("%s", string(contents))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 OK"))
	})

	fmt.Printf("Listening on port %d.\n", *portNumber)
	http.ListenAndServe(fmt.Sprintf(":%d", *portNumber), nil)
}
