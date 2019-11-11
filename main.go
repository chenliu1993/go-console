package main

import (
	"github.com/chenliu1993/go-console/pkg"
)

func main() {
	unixProcs, err := pkg.UnixProcesses()
	return
}
