package main

import (
	"log"
	"github.com/chenliu1993/go-console/pkg"
)


func main {
	unixProcTty, err := pkg.GetUnixProcess()
	if err != nil {
		log.Fatal(err)
	}
	sock, err := unixProcTty.SetUnixProcessIO()
	if err != nil {
		log.Fatal(err)
	}
	err := pkg.HandleSocket(sock)
	if err != nil {
		log.Fatal(err)
	}
	return
}
