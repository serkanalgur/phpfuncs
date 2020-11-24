package phpfuncs

import (
	"os"
	"path/filepath"
)

// Basename - Returns trailing name component of path.
// Original : https://www.php.net/manual/en/function.basename.php
// Given a string containing the path to a file or directory, this function will return the trailing name component.
func Basename(path string) string {
	return filepath.Base(path)
}

// Chgrp - Changes file group.
// Original : https://www.php.net/manual/en/function.chgrp.php
// Attempts to change the group of the file filename to group.
// Only the superuser may change the group of a file arbitrarily; other users may change the group of a file to any group of which that user is a member.
func Chgrp(name string, uid, gid int) error {
	return os.Chown(name, uid, gid)
}

// Chmod - Changes file mode
// Original : https://www.php.net/manual/en/function.chmod.php
//Attempts to change the mode of the specified file to that given in mode.
func Chmod(name string, mode os.FileMode) error {
	return os.Chmod(name, mode)
}

// Chown - Changes file owner.
// Original : https://www.php.net/manual/en/function.chown.php
// Attempts to change the owner of the file filename to user user. Only the superuser may change the owner of a file.
func Chown(name string, uid int, gid int) error {
	return os.Chown(name, uid, gid)
}
