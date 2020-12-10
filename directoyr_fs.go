package phpfuncs

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
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
// Mode : os.O_RDONLY | os.O_WRONLY | os.O_RDWR | os.O_APPEND | os.O_CREATE | os.O_EXCL | os.O_SYNC | os.O_TRUNC
func FOpen(file string, mode int) (*os.File, error) {
	f, err := os.OpenFile(file, mode|os.O_CREATE, 0644)
	if err != nil {
		return f, err
	}
	return f, err
}

// FRead - Binary-safe file read.
// Original : https://www.php.net/manual/en/function.fread.php
// fread() reads up to length bytes from the file pointer referenced by handle.
func FRead(f *os.File, sb int64) string {
	r := bufio.NewReader(f)
	b := make([]byte, sb)
	var cls string

	for {
		n, err := r.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		cls += string(b[0:n])
	}
	return cls
}

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

// FilePerms - Gets file permissions.
// Original : https://www.php.net/manual/en/function.fileperms.php
// Gets permissions for the given file.
func FilePerms(path string) os.FileMode {
	p, _ := os.Open(path)
	m, _ := p.Stat()
	p.Close()
	return m.Mode().Perm()
}

// FileSize - Gets file permissions.
// Original : https://www.php.net/manual/en/function.filesize.php
// Gets permissions for the given file.
func FileSize(path string) (int64, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

// FileType - Gets file type.
// Original : https://www.php.net/manual/en/function.filetype.php
// Returns the type of the given file.
func FileType(fs string) (string, error) {
	f, err := os.Open(fs)
	if err != nil {
		return "", err
	}
	defer f.Close()
	buffer := make([]byte, 512)
	fff, err := f.Read(buffer)
	if err != nil {
		fmt.Println(fff)
		return "", err
	}
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}

// FileGetContents - Reads entire file into a string.
// Original : https://www.php.net/manual/en/function.file-get-contents.php
// This function is similar to file(), except that file_get_contents() returns the file in a string, starting at the specified offset up to maxlen bytes. On failure, file_get_contents() will return FALSE.
// TODO : Context Implementation.
func FileGetContents(path string, includePath bool, context []string, offset int, maxlen int) string {
	var v string
	if IsURL(path) {
		includePath = false
	}

	if includePath == true {
		fileHere := FileExists(path)
		if fileHere {
			file, _ := FOpen(path, os.O_RDONLY)

			if offset >= 0 && maxlen != 0 {
				var err error
				r := bufio.NewReader(file)
				if offset > 0 {
					_, err := r.Discard(offset)
					if err != nil {
						log.Fatalln(err)
					}
				}
				buf := new(strings.Builder)
				_, err = io.CopyN(buf, r, int64(maxlen-offset))
				if err != nil {
					log.Fatal(err)
				}
				v = buf.String()

			} else {
				v = FRead(file, 512)
			}
			FClose(file)
		}
	} else {
		u, err := http.Get(path)
		if err != nil {
			log.Fatalln(err)
		}

		defer u.Body.Close()
		if offset >= 0 && maxlen > 0 {
			var err error
			r := bufio.NewReader(u.Body)
			if offset > 0 {
				_, err := r.Discard(offset)
				if err != nil {
					log.Fatalln(err)
				}
			}
			buf := new(strings.Builder)
			_, err = io.CopyN(buf, r, int64(maxlen-offset))
			if err != nil {
				log.Fatal(err)
			}
			v = buf.String()
		} else {
			body, err := ioutil.ReadAll(u.Body)
			if err != nil {
				log.Fatalln(err)
			}
			v = string(body)
		}
	}

	return v
}

// FilePutContents - Write data to a file
// Original : https://www.php.net/manual/en/function.file-put-contents.php
// This function is identical to calling fopen(), fwrite() and fclose() successively to write data to a file.
// TODO: Flags
func FilePutContents(path, data string) {
	f, err := FOpen(path, os.O_RDWR|os.O_CREATE)
	if err != nil {
		log.Fatalln(err)
	}
	FWrite(f, data)
	FClose(f)
}

// FPuts - Alias of FWrite
// Original : https://www.php.net/manual/en/function.fputs.php
// This function is an alias of: FWrite()..
func FPuts(f *os.File, data string) int {
	return FWrite(f, data)
}

// FWrite - Binary-safe file write
// Original : https://www.php.net/manual/en/function.fwrite.php
// fwrite() writes the contents of string to the file stream pointed to by handle.
func FWrite(f *os.File, data string) int {
	dataToByte := []byte(data)
	fw, err := f.Write(dataToByte)
	if err != nil {
		log.Fatal(err)
	}
	return fw
}

// Glob - Find pathnames matching a pattern.
// Original : https://www.php.net/manual/en/function.glob.php
// The glob() function searches for all the pathnames matching pattern according to the rules used by the libc glob() function, which is similar to the rules used by common shells.
func Glob(path string) (matches []string, err error) {
	return filepath.Glob(path)
}

// IsDir - Tells whether the filename is a directory.
// Original : https://www.php.net/manual/en/function.is-dir.php
// Tells whether the given filename is a directory.
func IsDir(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.IsDir()
}

// IsExecutable - Tells whether the filename is executable
// Original : https://www.php.net/manual/en/function.is-executable.php
// Tells whether the filename is executable.
func IsExecutable(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		log.Fatal(err)
	}
	return fi.Mode()&0100 != 0
}

// IsFile - Tells whether the filename is a regular file.
// Original : https://www.php.net/manual/en/function.is-file.php
// Tells whether the given file is a regular file.
func IsFile(path string) bool {
	file, err := os.Stat(path)
	return err == nil && file.Mode().IsRegular()
}

// IsLink - Tells whether the filename is a symbolic link.
// Original : https://www.php.net/manual/en/function.is-link.php
// Tells whether the given file is a symbolic link.
func IsLink(path string) bool {
	_, err := os.Readlink(path)
	return err == nil
}

// IsReadable - Tells whether a file exists and is readable.
// Original : https://www.php.net/manual/en/function.is-readable.php
// Tells whether a file exists and is readable.
func IsReadable(path string) bool {
	file, err := os.OpenFile(path, os.O_WRONLY, 0666)
	file.Close()
	return err == nil
}

// IsWritable - Tells whether the filename is writable.
// Original : https://www.php.net/manual/en/function.is-writable.php
// Returns TRUE if the filename exists and is writable. The filename argument may be a directory name allowing you to check if a directory is writable.
func IsWritable(path string) bool {
	file, err := os.OpenFile(path, os.O_WRONLY, 0)
	file.Close()
	return err == nil
}

// IsWriteable - Tells whether the filename is writable.
// Original : https://www.php.net/manual/en/function.is-writeable.php
// Returns TRUE if the filename exists and is writable. The filename argument may be a directory name allowing you to check if a directory is writable.
func IsWriteable(path string) bool {
	return IsWritable(path)
}

// Link - Create a hard link
// Original : https://www.php.net/manual/en/function.link.php
// link() creates a hard link.
func Link(target, link string) {
	err := os.Link(target, link)
	if err != nil {
		log.Fatal(err)
	}
}

// SymLink - Creates a symbolic link
// Original : https://www.php.net/manual/en/function.symlink.php
// symlink() creates a symbolic link to the existing target with the specified name link.
func SymLink(target, link string) {
	err := os.Symlink(target, link)
	if err != nil {
		log.Fatal(err)
	}
}

// MkDir - Makes directory.
// Original : https://www.php.net/manual/en/function.mkdir.php
// Attempts to create the directory specified by pathname.
func MkDir(path string, mode os.FileMode) error {
	return os.Mkdir(path, mode)
}

// ReadLink - Returns the target of a symbolic link.
// Original : https://www.php.net/manual/en/function.readlink.php
// readlink() does the same as the readlink C function.
func ReadLink(path string) (string, error) {
	li, err := os.Readlink(path)
	if err != nil {
		return "", err
	}
	return li, err
}

// RealPath - Returns canonicalized absolute pathname.
// Original : https://www.php.net/manual/en/function.realpath.php
// realpath() expands all symbolic links and resolves references to /./, /../ and extra / characters in the input path and returns the canonicalized absolute pathname.
func RealPath(path string) (string, error) {
	return filepath.Abs(path)
}

// Rename - Renames a file or directory.
// Original : https://www.php.net/manual/en/function.rename.php
// Attempts to rename oldname to newname, moving it between directories if necessary. If renaming a file and newname exists, it will be overwritten. If renaming a directory and newname exists, this function will emit a warning.
func Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

// RmDir — Removes directory.
// Original : https://www.php.net/manual/en/function.rmdir.php
// Attempts to remove the directory named by dirname. The directory must be empty, and the relevant permissions must permit this. A E_WARNING level error will be generated on failure.
func RmDir(path string) error {
	return os.RemoveAll(path)
}

// Stat - Gives information about a file.
// Original : https://www.php.net/manual/en/function.stat.php
// Gathers the statistics of the file named by filename. If filename is a symbolic link, statistics are from the file itself, not the symlink. Prior to PHP 7.4.0, on Windows NTS builds the size, atime, mtime and ctime statistics have been from the symlink, in this case.
func Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

// Touch - Sets access and modification time of file
// Original : https://www.php.net/manual/en/function.touch.php
// Attempts to set the access and modification times of the file named in the filename parameter to the value given in time. Note that the access time is always modified, regardless of the number of parameters.
// If the file does not exist, it will be created.
func Touch(path string, t int64, at int64) bool {

	_, err := FOpen(path, os.O_RDWR)
	if err != nil {
		log.Fatal(err)
		return false
	}

	atim := time.Unix(at, 0)
	ttim := time.Unix(t, 0)

	if err := os.Chtimes(path, ttim, atim); err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// Tempnam - Create file with unique file name
// Original : https://www.php.net/manual/en/function.tempnam.php
// Creates a file with a unique filename, with access permission set to 0600, in the specified directory. If the directory does not exist or is not writable, tempnam() may generate a file in the system's temporary directory, and return the full path to that file, including its name.
func Tempnam(dir, prefix string) string {
	raName := StringWithCharset(30, charset)
	p := filepath.FromSlash(dir + prefix + raName)
	_, err := os.OpenFile(p, os.O_CREATE, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	return p
}

// Tempfile - Creates a temporary file
// Original : https://www.php.net/manual/en/function.tmpfile.php
// Creates a temporary file with a unique name in read-write (w+) mode and returns a file handle.
//
// The file is automatically removed when closed (for example, by calling fclose(), or when there are no remaining references to the file handle returned by tmpfile()), or when the script ends.
func Tempfile() (f *os.File) {
	raName := StringWithCharset(30, charset)
	tmpfile, err := ioutil.TempFile("", raName)
	if err != nil {
		log.Fatal(err)
	}
	return tmpfile
}

// Unlink - Deletes a file.
// Original : https://www.php.net/manual/en/function.unlink.php
// Deletes filename. Similar to the Unix C unlink() function. An E_WARNING level error will be generated on failure.
func Unlink(name string) error {
	return Delete(name)
}

// ByteCountIEC - Bytecount & Humanize Bytes
// Complete calculator for DiskFreeSize
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
