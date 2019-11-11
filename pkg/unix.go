package pkg

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"
)

type UnixProcessTty struct {
	pid     int
	ppid    int
	exec    string
	console *os.File
}

func (p *UnixProcessTty) SetUnixProcessIO() (string, error) {
	sockpath := filepath.Join("./", p.exec, ".sock")
	conn, err := net.Dial("unix", sockpath)
	if err != nil {
		return "", err
	}
	uc, ok := conn.(*net.UnixConn)
	if !ok {
		return "", fmt.Errorf("casting to UnixConn failed")
	}
	socket, err := uc.File()
	if err != nil {
		return "", err
	}
	p.console = socket
	return sockpath, nil
}

func GetUnixProcess(pid int) (*UnixProcessTty, error) {
	dir := fmt.Sprintf("/proc/%d", pid)
	_, err := os.Stat(dir)
	if err != nil {
		// if os.IsNotExist(err) {
		// 	return nil
		// }
		return nil, err
	}
	stat := filepath.Join(dir, "stat")
	data, err := ioutil.ReadFile(stat)
	if err != nil {
		return nil, err
	}
	content := string(data)
	binStart := strings.IndexRune(content, '(') + 1
	binEnd := strings.IndexRune(content[binStart:], ')')
	exec := content[binStart : binStart+binEnd]
	return &UnixProcessTty{
		pid:  pid,
		ppid: os.Getppid(),
		exec: exec,
	}, nil
}
