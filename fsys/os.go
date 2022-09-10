package fsys

import (
	"os"
	"time"
)

// OSfs is a FS implementation that uses functions provided by the os package.
type OSfs struct{}

func NewOSfs() FS {
	return &OSfs{}
}

func (OSfs) Name() string { return "OSfs" }

func (OSfs) Create(name string) (File, error) {
	f, e := os.Create(name)
	if f == nil {
		// f is nil of type *os.File
		// we return nil of type nil
		return nil, e
	}
	return f, e
}

func (OSfs) Mkdir(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}

func (OSfs) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (OSfs) Open(name string) (File, error) {
	f, e := os.Open(name)
	if f == nil {
		// f is nil of type *os.File
		// we return nil of type nil
		return nil, e
	}
	return f, e
}

func (OSfs) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	f, e := os.OpenFile(name, flag, perm)
	if f == nil {
		// f is nil of type *os.File
		// we return nil of type nil
		return nil, e
	}
	return f, e
}

func (OSfs) Remove(name string) error {
	return os.Remove(name)
}

func (OSfs) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func (OSfs) Rename(oldname, newname string) error {
	return os.Rename(oldname, newname)
}

func (OSfs) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (OSfs) Chmod(name string, mode os.FileMode) error {
	return os.Chmod(name, mode)
}

func (OSfs) Chown(name string, uid, gid int) error {
	return os.Chown(name, uid, gid)
}

func (OSfs) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return os.Chtimes(name, atime, mtime)
}

func (OSfs) Lstat(name string) (os.FileInfo, error) {
	return os.Lstat(name)
}

func (OSfs) Symlink(oldname, newname string) error {
	return os.Symlink(oldname, newname)
}

func (OSfs) Readlink(name string) (string, error) {
	return os.Readlink(name)
}
