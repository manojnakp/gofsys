package fsys

import (
	"os"
	"testing"
	"time"
	// "github.com/spf13/afero"
)

func TestName(t *testing.T) {
	fs := NewOSfs()
	if name := fs.Name(); name != "OSfs" {
		t.Errorf("Wrong name of osfs")
	}
}

func TestCreate(t *testing.T) {
	fs := &OSfs{}
	file, err := fs.Create("tmp/normal/create")
	if err != nil {
		t.Errorf("Error creating a file: %v", err)
	}
	if file.Name() != "tmp/normal/create" {
		t.Errorf("Wrong file created")
	}
	_, err = fs.Create("tmp/doesnotexist/filename")
	if err == nil {
		t.Errorf("Expecting an error creating file: %v", err)
	}
}

func TestMkdir(t *testing.T) {
	fs := &OSfs{}
	if err := fs.Mkdir("tmp/mkdir", 0755); err != nil {
		t.Errorf("Error creating directory: %v", err)
	}
}

func TestMkdirAll(t *testing.T) {
	fs := &OSfs{}
	if err := fs.MkdirAll("tmp/mkdirall/parent/child", 0755); err != nil {
		t.Errorf("Error creating directories: %v", err)
	}
}

func TestOpen(t *testing.T) {
	fs := &OSfs{}
	if err := fs.Mkdir("tmp/normal/open", 0755); err != nil {
		t.Fatalf("Error creating temp directory: %v", err)
	}
	if _, err := fs.Create("tmp/normal/open/file"); err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	dir, err := fs.Open("tmp/normal/open/")
	if err != nil {
		t.Errorf("Error opening directory: %v", err)
	}
	defer dir.Close()
	file, err := fs.Open("tmp/normal/open/file")
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	defer file.Close()
	_, err = fs.Open("tmp/normal/doesnotexist/filename")
	if err == nil {
		t.Errorf("Error expected in creating file: %v", err)
	}
}

func TestOpenFile(t *testing.T) {
	fs := &OSfs{}
	if err := fs.Mkdir("tmp/normal/openfile", 0755); err != nil {
		t.Fatalf("Error creating temp directory: %v", err)
	}
	if _, err := fs.Create("tmp/normal/openfile/file"); err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	dir, err := fs.OpenFile("tmp/normal/openfile/", os.O_RDONLY, 0755)
	if err != nil {
		t.Errorf("Error opening directory: %v", err)
	}
	defer dir.Close()
	file, err := fs.OpenFile("tmp/normal/openfile/file", os.O_RDONLY, 0644)
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	defer file.Close()
	_, err = fs.OpenFile("tmp/normal/doesnotexist/filename", os.O_EXCL, 0644)
	if err == nil {
		t.Errorf("Error expected in creating file: %v", err)
	}
}

func TestRemove(t *testing.T) {
	fs := &OSfs{}
	_, err := fs.Create("tmp/normal/remove")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	err = fs.Mkdir("tmp/normal/removedir", 0755)
	if err != nil {
		t.Fatalf("Error creating temp dir: %v", err)
	}
	err = fs.Remove("tmp/normal/remove")
	if err != nil {
		t.Errorf("Error removing the files")
	}
	err = fs.Remove("tmp/normal/removedir")
	if err != nil {
		t.Errorf("Error removing the directory")
	}
}

func TestRemoveAll(t *testing.T) {
	fs := &OSfs{}
	if err := fs.MkdirAll("tmp/normal/parent/child", 0755); err != nil {
		t.Fatalf("Error creating temp directories: %v", err)
	}
	_, err := fs.Create("tmp/normal/parent/child/file")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	err = fs.RemoveAll("tmp/normal/parent")
	if err != nil {
		t.Errorf("Error removing the files")
	}
}

func TestRename(t *testing.T) {
	fs := &OSfs{}
	if err := fs.Mkdir("tmp/normal/rename", 0755); err != nil {
		t.Fatalf("Error creating temp directory: %v", err)
	}
	_, err := fs.Create("tmp/normal/rename/file")
	if err != nil {
		t.Fatalf("Error creating temp file %v", err)
	}
	err = fs.Rename("tmp/normal/rename/file", "tmp/normal/rename/newfile")
	if err != nil {
		t.Errorf("Error renaming the file: %v", err)
	}
}

func TestChmod(t *testing.T) {
	fs := &OSfs{}
	if err := fs.Mkdir("tmp/normal/chmod", 0755); err != nil {
		t.Fatalf("Error creating temp directory: %v", err)
	}
	_, err := fs.Create("tmp/normal/chmod/file")
	if err != nil {
		t.Fatalf("Error creating temp file %v", err)
	}
	err = fs.Chmod("tmp/normal/chmod/file", 0777)
	if err != nil {
		t.Errorf("Error with chmod: %v", err)
	}
}

func TestChown(t *testing.T) {
	fs := &OSfs{}
	if err := fs.Mkdir("tmp/normal/chown", 0755); err != nil {
		t.Fatalf("Error creating temp directory: %v", err)
	}
	_, err := fs.Create("tmp/normal/chown/file")
	if err != nil {
		t.Fatalf("Error creating temp file %v", err)
	}
	err = fs.Chown("tmp/normal/chown/file", 0, 0)
	if err != nil {
		t.Errorf("Error with chmod: %v", err)
	}
}

func TestStat(t *testing.T) {
	fs := &OSfs{}
	if err := fs.Mkdir("tmp/normal/stat", 0755); err != nil {
		t.Fatalf("Error creating temp directory: %v", err)
	}
	file := "tmp/normal/stat/file"
	_, err := fs.Create(file)
	if err != nil {
		t.Fatalf("Error creating temp file %v", err)
	}
	_, err = fs.Stat(file)
	if err != nil {
		t.Errorf("Error with stat: %v", err)
	}
}

func TestLstat(t *testing.T) {
	fs := &OSfs{}
	if err := fs.Mkdir("tmp/normal/lstat", 0755); err != nil {
		t.Fatalf("Error creating temp directory: %v", err)
	}
	from := "tmp/normal/lstat/file"
	to := "tmp/normal/lstat/symlink"
	_, err := fs.Create(from)
	if err != nil {
		t.Fatalf("Error creating temp file %v", err)
	}
	err = fs.Symlink(from, to)
	if err != nil {
		t.Fatalf("Error creating temp symlink: %v", err)
	}
	_, err = fs.Lstat(to)
	if err != nil {
		t.Errorf("Error with lstat: %v", err)
	}
}
func TestChtimes(t *testing.T) {
	fs := &OSfs{}
	if err := fs.Mkdir("tmp/normal/chtimes", 0755); err != nil {
		t.Fatalf("Error creating temp directory: %v", err)
	}
	_, err := fs.Create("tmp/normal/chtimes/file")
	if err != nil {
		t.Fatalf("Error creating temp file %v", err)
	}
	err = fs.Chtimes("tmp/normal/chtimes/file", time.Now(), time.Now())
	if err != nil {
		t.Errorf("Error changing times: %v", err)
	}
}
func TestSymlink(t *testing.T) {
	fs := &OSfs{}
	if err := fs.Mkdir("tmp/normal/link", 0755); err != nil {
		t.Fatalf("Error creating temp directory: %v", err)
	}
	_, err := fs.Create("tmp/normal/link/file")
	if err != nil {
		t.Fatalf("Error creating temp file %v", err)
	}
	err = fs.Symlink("tmp/normal/link/file", "tmp/normal/link/symlink")
	if err != nil {
		t.Errorf("Error creating symlink: %v", err)
	}
}

func TestReadlink(t *testing.T) {
	fs := &OSfs{}
	if err := fs.Mkdir("tmp/normal/readlink", 0755); err != nil {
		t.Fatalf("Error creating temp directory: %v", err)
	}
	from := "tmp/normal/readlink/file"
	to := "tmp/normal/readlink/symlink"
	_, err := fs.Create(from)
	if err != nil {
		t.Fatalf("Error creating temp file %v", err)
	}
	err = fs.Symlink(from, to)
	if err != nil {
		t.Fatalf("Error creating temp symlink: %v", err)
	}
	ln, err := fs.Readlink(to)
	if err != nil || ln != from {
		t.Errorf("Error reading link file: %v", err)
	}
}
