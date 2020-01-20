package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
	cssh "golang.org/x/crypto/ssh"
)

func publicKeyFile(file string) (ssh.AuthMethod, error) {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(key), nil
}

func executeCmd(command, hostname string, port string, config *ssh.ClientConfig) string {
	conn, _ := ssh.Dial("tcp", fmt.Sprintf("%s:%s", hostname, port), config)
	session, _ := conn.NewSession()
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run(command)

	return fmt.Sprintf("%s -> %s", hostname, stdoutBuf.String())
}
func main() {
	cmd := os.Args[1]
	hosts := os.Args[2]
	// timeout := time.After(10 * time.Second)
	timeout := time.After(10 * time.Second)
	results := make(chan string, 10)
	// timeout := time.After(10 * time.Second)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "22"
	}
	publicKey, _ := publicKeyFile("/home/maxime/github/validator_maxime/data/id_shared")

	config := &ssh.ClientConfig{
		User: "sles",
		Auth: []ssh.AuthMethod{
			publicKey,
		},
		HostKeyCallback: cssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	go func(hostname string, port string) {
		results <- executeCmd(cmd, hostname, port, config)
	}(hosts, port)

	select {
	case res := <-results:
		fmt.Print(res)
	case <-timeout:
		fmt.Println("Timed out!")
		return
	}

	fmt.Print(results)
}
