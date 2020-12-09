package phpfuncs

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Exec - Start a command on system
// Original : https://www.php.net/manual/tr/function.exec.php
// exec() executes the given command.
func Exec(of string) {
	var out bytes.Buffer
	cmd := exec.Command(of)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// ShellExec - Execute command via shell and return the complete output as a string
// Original : https://www.php.net/manual/en/function.shell-exec.php
// This function is identical to the backtick operator.
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

// Exit - Output a message and terminate the current script
// Original : https://www.php.net/manual/en/function.exit.php
// Terminates execution of the script. Shutdown functions and object destructors will always be executed even if exit is called.
func Exit(of int) {
	os.Exit(of)
}

// Die - Equivalent to exit
// Original : https://www.php.net/manual/en/function.die.php
// This language construct is equivalent to exit().
func Die(of int) {
	os.Exit(of)
}