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

# Working with JSON and XML

* struct tags
* struct generators

# Working with SQL

* Example from [dvmweb](https://github.com/miku/dvmweb)

