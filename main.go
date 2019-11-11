package main

import (
	"log"
	"github.com/chenliu1993/go-console/pkg"
)


func main() {
	unixProcTty, err := pkg.GetUnixProcess(68760)
	if err != nil {
		log.Fatal(err)
	}
	sock := "test.sock"
	err = pkg.HandleSocket(sock)
        if err != nil {
                log.Fatal(err)
        }
	err = unixProcTty.SetUnixProcessIO()
	if err != nil {
		log.Fatal(err)
	}
	return
}
