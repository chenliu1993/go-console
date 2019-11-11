package pkg

import (
	"fmt"
	"os"
	"io"
	"net"
	"github.com/opencontainers/runc/libcontainer/utils"
	"github.com/containerd/console"
)

func HandleSocket(path string) error {
	ln, err := net.Listen("unix", path)
	if err != nil {
		return err
	}
	defer ln.Close()
	conn, err := ln.Accept()
	if err !=nil {
		return err
	}
	defer conn.Close()
	ln.Close()

	unixconn, ok := conn.(*net.UnixConn)
	if !ok {
		return fmt.Errorf("failed to cast to unixconn")
	}
	socket, err := unixconn.File()
	if err != nil {
		return err
	}
	defer socket.Close()

	// Get the master file descriptor from runC.
	master, err := utils.RecvFd(socket)
	if err != nil {
		return err
	}
	c, err := console.ConsoleFromFile(master)
	if err != nil {
		return err
	}
	console.ClearONLCR(c.Fd())
	// Copy from our stdio to the master fd.
	quitChan := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, c)
		quitChan <- struct{}{}
	}()
	go func() {
		io.Copy(c, os.Stdin)
		quitChan <- struct{}{}
	}()

	// Only close the master fd once we've stopped copying.
	<-quitChan
	c.Close()
	return nil
}
