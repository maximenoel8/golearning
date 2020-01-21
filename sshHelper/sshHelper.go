package sshhelper

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
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

func NewClient(hostname string) *ssh.Client {
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
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}
	conn, _ := ssh.Dial("tcp", fmt.Sprintf("%s:%s", hostname, port), config)
	return conn
}

func ExecuteCmd(command string, hostname string) string {

	conn := NewClient(hostname)

	// conn, _ := ssh.Dial("tcp", fmt.Sprintf("%s:%s", hostname, port), config)
	session, _ := conn.NewSession()
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run(command)

	return stdoutBuf.String()
}

// func executeCmd(command, hostname string, port string, config *ssh.ClientConfig) string {

// 	port := os.Getenv("PORT")
// 	if len(port) == 0 {
// 		port = "22"
// 	}
// 	publicKey, _ := publicKeyFile("/home/maxime/github/validator_maxime/data/id_shared")

// 	config := &ssh.ClientConfig{
// 		User: "sles",
// 		Auth: []ssh.AuthMethod{
// 			publicKey,
// 		},
// 		HostKeyCallback: cssh.InsecureIgnoreHostKey(),
// 		Timeout:         5 * time.Second,
// 	}

// 	conn, _ := ssh.Dial("tcp", fmt.Sprintf("%s:%s", hostname, port), config)
// 	session, _ := conn.NewSession()
// 	defer session.Close()

// 	var stdoutBuf bytes.Buffer
// 	session.Stdout = &stdoutBuf
// 	session.Run(command)

// 	return fmt.Sprintf("%s -> %s", hostname, stdoutBuf.String())
// }
