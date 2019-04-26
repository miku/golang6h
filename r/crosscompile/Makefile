all: main main.arm main.exe

main: main.go
	go build -o $@ $<

main.arm: main.go
	GOOS=linux GOARCH=arm go build -o $@ $<

main.exe: main.go
	GOOS=windows GOARCH=amd64 go build -o $@ $<

