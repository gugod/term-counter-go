
package main

import (
	"strings"
	"net/http"
	"encoding/json"
)

var termcount map[string]int;

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		tc := make( map[string]int )

		r.ParseForm();
		arr := strings.Split( r.Form.Get("q"), " " )
		for i := 0; i < len(arr); i++ {
			termcount[ arr[i] ]++
			tc[arr[i]] = termcount[ arr[i] ]
		}

		b, err := json.Marshal(tc)
		if err == nil {
			w.Write(b)
		} 
	}

	if r.Method == "GET" {
		b, err := json.Marshal(termcount)
		if err == nil {
			w.Write(b)
		} 
	}
}

func main() {
	termcount = make(map[string]int)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
