
package main

import (
	"fmt"
	"net/http"
	"strings"
)

var termcount map[string]int;

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm();
		arr := strings.Split( r.Form.Get("q"), " " )
		for i := 0; i < len(arr); i++ {
			termcount[ arr[i] ]++;
			fmt.Fprintf(w, "%s = %d\n", arr[i], termcount[arr[i]])
		}
	}
}

func main() {
	termcount = make(map[string]int)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
