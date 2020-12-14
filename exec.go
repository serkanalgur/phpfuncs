package phpfuncs

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var rPatt *regexp.Regexp


func init() {
	rPatt = regexp.MustCompile(`[^\w@%+=:,./-]`)
}

// Exec - Start a command on system
//
// Original : https://www.php.net/manual/tr/function.exec.php
//
// exec() executes the given command.
func Exec(of string) string {
	var out bytes.Buffer
	cmd := exec.Command(of)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}

// ShellExec - Execute command via shell and return the complete output as a string
//
// Original : https://www.php.net/manual/en/function.shell-exec.php
//
// This function is identical to the backtick operator.
func ShellExec(of string) string {
	var out bytes.Buffer
	cmd := exec.Command(of)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}

// Exit - Output a message and terminate the current script
//
// Original : https://www.php.net/manual/en/function.exit.php
//
// Terminates execution of the script. Shutdown functions and object destructors will always be executed even if exit is called.
func Exit(of int) {
	os.Exit(of)
}

// Die - Equivalent to exit
//
// Original : https://www.php.net/manual/en/function.die.php
//
// This language construct is equivalent to exit().
func Die(of int) {
	os.Exit(of)
}

// Escapeshellarg - Escape a string to be used as a shell argument
//
// Original: https://www.php.net/manual/en/function.escapeshellarg.php
//
func Escapeshellarg(s string) string {
	if len(s) == 0 {return "''"}

	if rPatt.MatchString(s) {
		return "'" + strings.ReplaceAll(s, "'", "'\"'\"'") + "'"
	}

	return s
}

// Escapeshellcmd - Escape shell metacharacters
//
// Original: https://www.php.net/manual/en/function.escapeshellcmd.php
//
func Escapeshellcmd(s string) string {
	cmds := Explode(s, " ") // I know thats maybe not proper way...

	z := make([]string, len(cmds))

		for i ,s := range cmds {
			z[i] = Escapeshellarg(s)
		}

		return Join(" ",z)
}