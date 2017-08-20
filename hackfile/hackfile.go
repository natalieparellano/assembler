package hackfile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func NewPath(path, oldExtn, newExtn string) string {
	fmt.Printf("path: %s\n", path)

	dir := filepath.Dir(path)
	fmt.Printf("dir: %s\n", dir)

	oldfile := filepath.Base(path)
	fmt.Printf("oldfile: %s\n", oldfile)

	newfile := strings.Replace(oldfile, oldExtn, newExtn, 1)
	fmt.Printf("newfile: %s\n", newfile)

	newpath := filepath.Join(dir, newfile)
	fmt.Printf("newpath: %s\n", newpath)

	return newpath
}

func WriteFile(path, str string) {
	f, err := os.Create(path)
	check(err)

	_, err = f.WriteString(str)
	check(err)

	err = f.Sync()
	check(err)

	err = f.Close()
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
