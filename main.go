package main

import (
	"fmt"
	"io"
	"os"

	cli "github.com/jawher/mow.cli"
	shell "github.com/noffle/easy-ipfs-shell/shell"
)

func main() {
	cmd := cli.App("ipcat", "Retrieve and save IPFS objects.")
	cmd.Spec = "IPFS_PATH"

	hash := cmd.String(cli.StringArg{
		Name:  "IPFS_PATH",
		Value: "",
		Desc:  "the IPFS object path",
	})

	cmd.Action = func() {
		if err := cat(*hash); err != nil {
			fmt.Fprintf(os.Stderr, "ipcat failed: %s\n", err)
			os.Exit(2)
		}
	}
	cmd.Run(os.Args)
}

func cat(path string) error {
	shell, err := shell.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	reader, err := shell.Cat("QmVVjWrps58cFS1hSvCdAxmS4wggKfRGbDzJway6QCxR4U")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, reader)

	return nil
}
