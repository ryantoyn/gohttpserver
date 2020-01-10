package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	fmt.Printf("%s: http.ListenAndServe on 8080\n", time.Now().UTC().Format(time.RFC3339))
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", nil)
}

func myHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s: received from %s\n", time.Now().UTC().Format(time.RFC3339), req.RemoteAddr)
	fmt.Fprintf(w, "Path: /%s\n", req.URL.Path[1:])
	fmt.Fprintf(w, "RemoteAddr: %s\n", req.RemoteAddr)
	fmt.Fprintf(w, "User-Agent: %s\n", req.Header.Get("User-Agent"))

	// This will only be defined when site is accessed via non-anonymous proxy
	// and takes precedence over RemoteAddr
	// Header.Get is case-insensitive
	fmt.Fprintf(w, "X-Forwarded-For: %s\n", req.Header.Get("X-Forwarded-For"))
	fmt.Fprintf(w, "X-Forwarded-Host: %s\n", req.Header.Get("X-Forwarded-Host"))

	resp, err := http.Get("http://ifconfig.me/")
	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			fmt.Fprintf(w, "OutboundIP: %s\n", string(body))
		}
	}
	defer resp.Body.Close()

	fmt.Fprintf(w, "Date: %s\n", time.Now().UTC().Format(time.RFC3339))
}
