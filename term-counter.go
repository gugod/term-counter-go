
package main

import (
	"strings"
	"sync"
	"net/http"
	"encoding/json"
)

var (
	termcount map[string]int
	termcountMu sync.Mutex
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var wg sync.WaitGroup
		tc := make( map[string]int )

		r.ParseForm();
		arr := strings.Split( r.Form.Get("q"), " " )
		wg.Add( len(arr) )
		for i := 0; i < len(arr); i++ {
			go func(term string) {
				termcountMu.Lock()
				defer termcountMu.Unlock()
				termcount[ term ]++
				tc[term] = termcount[ term ]
				wg.Done()
			}(arr[i]);
		}
		wg.Wait()

		b, err := json.Marshal(tc)
		if err == nil {
			w.Write(b)
		} 
	}

	if r.Method == "GET" {
		termcountMu.Lock()
		defer termcountMu.Unlock()

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
