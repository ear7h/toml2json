package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/BurntSushi/toml"
)

func exitErr(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	var err error
	in := os.Stdin
	out := os.Stdout

	switch len(os.Args) {

	case 3:
		if os.Args[2] != "-" {
			in, err = os.OpenFile(os.Args[1], os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				exitErr(err)
			}
		}
		fallthrough

	case 2:
		if os.Args[1] ==  "--help" || os.Args[1] == "-h" {
			fmt.Fprintf(os.Stderr, "usage: %s [in [out]]\n", os.Args[0])
			os.Exit(0)
		}

		if os.Args[1] != "-" {
			in, err = os.Open(os.Args[1])
			if err != nil {
				exitErr(err)
			}
		}
		fallthrough
	case 1:
		err := run(in, out)
		if err != nil {
			exitErr(err)
		}
	default:
		fmt.Fprintf(os.Stderr, "usage: %s [in [out]]\n", os.Args[0])
		os.Exit(1)
	}
}

func run(in io.Reader, out io.Writer) error {
	var v interface{}
	_, err := toml.DecodeReader(in, &v)
	if err != nil {
		return err
	}

	return json.NewEncoder(out).Encode(v)
}

