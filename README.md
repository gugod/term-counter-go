term-counter-go
===============

Run:

    go run term-counter.go

Send query string with terms...

    curl -XPOST http://localhost:3000/'?q=foo+bar+baz'

Retrieve term counts.

    curl http://localhost:3000/

