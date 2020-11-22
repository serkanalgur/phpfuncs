package phpfuncs

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

// Exec - Start a command on system
func Exec(of string) {
	var out bytes.Buffer
	cmd := exec.Command(of)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// ShellExec - Start a command on system
func ShellExec(of string) {
	var out bytes.Buffer
	cmd := exec.Command(of)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q\n", out.String())
}
