// Package terminal provides utility function to read from terminals.
package terminal

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"syscall"

	"github.com/frankbraun/codechain/util/bzero"
	"golang.org/x/crypto/ssh/terminal"
)

// ReadPassphrase reads a single line from fd without local echo and returns
// it (without trailing newline). When confirm is true it reads a second line
// and makes sure both passphrases match.
func ReadPassphrase(fd int, confirm bool) ([]byte, error) {
	var (
		pass   []byte
		pass2  []byte
		reader *bufio.Reader
		err    error
	)
	isTerminal := terminal.IsTerminal(fd)
	fmt.Printf("passphrase: ")
	if isTerminal {
		pass, err = terminal.ReadPassword(fd)
		fmt.Println("")
	} else {
		reader = bufio.NewReader(os.NewFile(uintptr(fd), "terminal"))
		pass, err = reader.ReadBytes('\n')
	}
	if err != nil {
		if err == io.EOF {
			return nil, errors.New("unable to read passphrase")
		}
		return nil, err
	}
	if len(pass) == 0 {
		return nil, errors.New("please provide a passphrase")
	}
	pass = bytes.TrimRight(pass, "\n")
	if confirm {
		fmt.Printf("confirm passphrase: ")
		if isTerminal {
			pass2, err = terminal.ReadPassword(syscall.Stdin)
			fmt.Println("")
		} else {
			pass2, err = reader.ReadBytes('\n')
		}
		if err != nil {
			return nil, err
		}
		defer bzero.Bytes(pass2)
		pass2 = bytes.TrimRight(pass2, "\n")
		if !bytes.Equal(pass, pass2) {
			return nil, errors.New("passphrases don't match")
		}
	}
	return pass, nil
}

// ReadLine reads a single line from r it and returns it (without trailing
// newline).
func ReadLine(r io.Reader) ([]byte, error) {
	str, err := bufio.NewReader(r).ReadBytes('\n')
	if err != nil {
		if err == io.EOF {
			return nil, errors.New("unable to read line")
		}
		return nil, err
	}
	return bytes.TrimSpace(str), nil
}