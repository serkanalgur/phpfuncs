package phpfuncs

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

//ReAd constants
const (
	read   int = syscall.O_RDONLY // open the file read-only.
	write  int = syscall.O_WRONLY // open the file write-only.
	readp  int = syscall.O_RDWR   // open the file read-write.
	append int = syscall.O_APPEND // append data to the file when writing.
	create int = syscall.O_CREAT  // create a new file if none exists.
	excl   int = syscall.O_EXCL   // used with O_CREATE, file must not exist
	sync   int = syscall.O_SYNC   // open for synchronous I/O.
	trunc  int = syscall.O_TRUNC  // if possible, truncate file when opened.
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

// Copy - Copies file
// Original : https://www.php.net/manual/en/function.copy.php
// Makes a copy of the file source to dest.
func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// Delete - Deletes a file.
// Original : https://www.php.net/manual/en/function.delete.php
// Deletes filename. Similar to the Unix C unlink() function. An E_WARNING level error will be generated on failure.
func Delete(name string) error {
	return os.Remove(name)
}

// DirName - Returns a parent directory's path
// Original : https://www.php.net/manual/en/function.dirname.php
// Given a string containing the path of a file or directory, this function will return the parent directory's path that is levels up from the current directory.
func DirName(path string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(path)
}

//DiskStatus struct
type DiskStatus struct {
	Free string `json:"Free"`
}

// DiskFreeSpace - Returns available space on filesystem or disk partition
// Original : https://www.php.net/manual/en/function.disk-free-space.php
// Given a string containing a directory, this function will return the number of bytes available on the corresponding filesystem or disk partition.
// DEVELOPER NOTE : PROBABLY WORKING ON ONLY LINUX AND MAC. TO-DO : WINDOWS
func DiskFreeSpace(path string) (disk DiskStatus) {
	stat := syscall.Statfs_t{}
	err := syscall.Statfs(path, &stat)
	if err != nil {
		return
	}
	disk.Free = ByteCountIEC(stat.Bfree * uint64(stat.Bsize))
	return
}

// FClose - Closes an open file pointer
// Original : https://www.php.net/manual/en/function.fclose.php
// The file pointed to by handle is closed.
func FClose(file *os.File) error {
	return file.Close()
}

// FOpen - Opens file
// Original : https://www.php.net/manual/en/function.fopen.php
// fopen() binds a named resource, specified by filename, to a stream.
// NOT COMPLETED
// func FOpen(file string, mode int) (os.file, error) {
// 	f, err := os.OpenFile(file, mode, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := f.Close(); err != nil {
// 		log.Fatal(err)
// 	}

// 	defer file.Close()
// }

// FileExists - Checks whether a file or directory exists.
// Original : https://www.php.net/manual/en/function.file-exists.php
// Checks whether a file or directory exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// FileMime - Gets file modification time
// Original : https://www.php.net/manual/en/function.filemtime.php
// This function returns the time when the data blocks of a file were being written to, that is, the time when the content of the file was changed.
func FileMime(file string) time.Time {
	fi, err := os.Stat(file)
	if err != nil {
		return time.Time{}
	}
	return fi.ModTime()
}

// ByteCountIEC - Bytecount & Humanize Bytes
func ByteCountIEC(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}
