package pkg

import (
	"fmt"
	"strconv"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type UnixProcess struct {
	pid  int
	ppid int
	exec string
}

func GetUnixProcess(pid int) (*UnixProcess, error) {
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
	return &UnixProcess{
		pid:  pid,
		ppid: os.Getppid(),
		exec: exec,
	}, nil
}


func UnixProcesses() ([]*UnixProcess, error) {
	d, err := os.Open("/proc")
	if err != nil {
		return nil, err
	}
	defer d.Close()
	var unixProcesses []*UnixProcess
	for {
		files, err := d.Readdir(10)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		for _, file := range files {
			if !file.IsDir() {
				continue
			}
			name := file.Name()
			if name[0] < '0' || name[0] > '9' {
				continue
			}
			pid, err := strconv.ParseInt(name, 10, 0)
			if err != nil {
				continue
			}
			unixProcess, err := GetUnixProcess(int(pid))
			if err != nil {
				return nil, err
			}
			unixProcesses = append(unixProcesses, unixProcess)
		}
	}
	return unixProcesses, nil
}
