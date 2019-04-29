package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Dump out request details & header.
	fmt.Printf("\n==============================\n")
	fmt.Printf("%s %s %s\n", r.Proto, r.Method, r.URL.Path)
	fmt.Printf("\n")
	r.Header.Write(os.Stdout)
	fmt.Printf("\n")

	// Handle gzipped data
	var err error
	var body io.Reader = r.Body
	encoding := r.Header.Get("Content-Encoding")
	if encoding == "gzip" {
		dec, err := gzip.NewReader(r.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: gzip failure (%s).", err)
			return
		}
		defer dec.Close()
		body = dec
	}

	// Dump out the request body.
	byteCnt, err := io.Copy(os.Stdout, body)
	fmt.Printf("\n(%d bytes in body)\n", byteCnt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: io.Copy() failed (%s).", err)
	}
}

func usage() {

	fmt.Fprintf(os.Stderr, `Usage: %s [OPTIONS]
Dummy web server for debugging. Dumps requests to stdout.

Options:
`, os.Args[0])

	flag.PrintDefaults()
}

func main() {
	var addr string
	flag.Usage = usage

	flag.StringVar(&addr, "a", ":8080", "address to listen on")
	flag.Parse()

	http.ListenAndServe(addr, http.HandlerFunc(handler))
}
