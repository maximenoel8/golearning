package main

import (
	"os"

	scp "github.com/hnakamur/go-scp"
	sshhelper "github.com/maximenoel8/golearning/sshHelper"
)

func main() {

	// cmd := os.Args[1]
	hosts := os.Args[1]

	// results := sshhelper.ExecuteCmd(cmd, hosts)

	// fmt.Print(results)
	sscp := scp.NewSCP(sshhelper.NewClient(hosts))
	// defer sscp.Close()
	sscp.ReceiveFile("/tmp/maxime", "./maxime")
	sscp.SendFile("./kubeconf", "/tmp/totiti")
}
