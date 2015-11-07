package main

import (
	"fmt"
	"io"
	"os"

	cli "github.com/jawher/mow.cli"
	fallback "github.com/noffle/fallback-ipfs-shell"
)

func main() {
	cmd := cli.App("ipcat", "Retrieve IPFS object data and output it to stdout.")
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
	shell, err := fallback.NewShell()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	reader, err := shell.Cat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, reader)

	return nil
}
