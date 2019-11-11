package main

import (
	"fmt"
	"github.com/chenliu1993/go-console/pkg"
)

func main() {
	unixProcs, _ := pkg.UnixProcesses()
	for _, unixProc := range unixProcs {
		fmt.Printf("unixProc is %v\n", unixProc)
	}
	return
}
