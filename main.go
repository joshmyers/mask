package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/thamaji/ioutils"
	"golang.org/x/crypto/ssh/terminal"
)

const Version = "v1.0.0"

func main() {
	var help, version bool
	var secret string

	flag.BoolVar(&help, "h", false, "show help")
	flag.BoolVar(&version, "v", false, "show version")
	flag.StringVar(&secret, "s", "", "set secret word")

	flag.Usage = func() {
		fmt.Println()
		fmt.Println("Usage: " + os.Args[0] + " [OPTIONS] COMMAND [ARG...]")
		fmt.Println()
		fmt.Println("Run command with masked stdout")
		fmt.Println()
		fmt.Println("Options:")
		flag.CommandLine.PrintDefaults()
		fmt.Println()
	}

	flag.Parse()

	args := flag.Args()

	if help {
		flag.Usage()
		return
	}

	if version {
		fmt.Println(Version)
		return
	}

	if secret == "" {
		if !terminal.IsTerminal(int(os.Stdin.Fd())) {
			fmt.Fprintln(os.Stderr, "stdin is not a terminal")
			return
		}

		if len(args) == 0 {
			flag.Usage()
			return
		}

		fmt.Fprint(os.Stderr, "Enter secret word: ")
		bytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return
		}
		fmt.Fprintln(os.Stderr)

		secret = string(bytes)
	}

	writer := ioutils.NewMaskWriter(os.Stdout, []byte(secret), '*')

	if len(args) == 0 {
		if !terminal.IsTerminal(int(os.Stdin.Fd())) {
			io.Copy(writer, os.Stdin)
			return
		}

		flag.Usage()
		return
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	stdout, _ := cmd.StdoutPipe()
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	io.Copy(writer, stdout)

	if err := cmd.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
