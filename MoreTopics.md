# Cross compilation

```
$ GOOS=linux GOARCH=arm go build ...
```

# Testing

```
$ go test
```

* Table driven tests

# Web programming

An example server.

```
http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))
```

# The IO interfaces

* io.Reader
* io.Writer

Small interfaces, can be combined.

# Working with JSON and XML

* struct tags
* struct generators

# Working with SQL

* Example from [dvmweb](https://github.com/miku/dvmweb)

# Go and Docker example

* CGO_ENABLED hint

# More on tooling

* [Go tooling examples](https://www.alexedwards.net/blog/an-overview-of-go-tooling)

