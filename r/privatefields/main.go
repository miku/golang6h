package main

import "fmt"
import "github.com/miku/golang6h/r/private"

func main() {
	semi := private.Semi{}
	fmt.Printf("%+v\n", semi)
	semi.P = 1
	fmt.Printf("%+v\n", semi)
	// semi.priv undefined (cannot refer to unexported
	// field or method priv)
	semi.priv = 10
}
