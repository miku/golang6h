# golang6h

Golang in six hours (360 minutes). Language and tooling.

# Outline

* Diving in (90)
* Tooling revisited (30)
* Interfaces and type embedding (30)
* Concurreny (90)
* Standard library (60)
* Third-party libraries and tools (30)
* Wrap-Up (30)

About 24 example programs, 15 minutes per program, allow for 5 minutes of
experimentation.

## Diving in (90)

* Intro to language and first run
* A bit of history, general development approach

> r/helloworld.go

* How to write Go, editors, vscode, plugins
* play.golang.org
* Types and variable declarations

> r/vars.go

* Container: arrays and slices

> r/slices.go

* Container: map

> r/map.go

* Pointers

> r/pointers.go

* Compound: structs
* Go has value types (what is a value type)

> r/compound.go

* Struct tags

> r/tags.go

* functions are first class values
* you can declare function types
* closures

> r/hof.go

## Tooling revisited (30)

The go tool does a lot of things:

* go run
* go build
* go get
* go test (also: coverage, benchmark, race detector)

Always format your code:

* go fmt, goimports

Run basic checks.

* go vet, golint

Dependency management with modules:

* go mod

Building small docker containers, with alpine and CGO_ENABLED=0.

## Interfaces and type embedding (30)

Go uses structural typing.

> r/iface.go

You can embed types.

> r/embed.go

## Concurrency (90)

Go follows a CSP style concurrency, no callback, but channels, which are typed conduits.

> r/channels.go

A dummy example with sleeps.

> r/dummysleep.go

Coordinating processes with select.

> r/timeout.go

Supporting constructs, like mutexes.

> r/mutex.go

The mighty WaitGroup.

> r/waitgroup.go

## Standard library (60)

We have a beautiful standard library. People recommend to read it in full.

Writing a command line HTTP client, like curl.

> r/fetch1.go

Writing a parallel version of fetch (use your new concurrency knowledge).

> r/fetch2.go

Better error handling.

> r/fetch3.go

Writing a small HTTP server.

> r/server1.go

Beautiful IO with the io and bufio package. Maybe read an image and apply a
filter then write it out.

> r/imagefilter.go

Connect io with HTTP (e.g. like ch01 from gopl).

> r/lissajous.go

## Third party tools, general development approaches (30)

Where to find packages

* godoc.org

How to write your own packages.

* command line tools, project layout

Testing with Go, table driven tests.

> r/testing.go

XML, JSON and other serialization options.

> r/xml.go

## Wrap up, discussion, Q&A (30)

Fade out.
