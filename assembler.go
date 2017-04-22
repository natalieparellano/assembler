package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/natalieparellano/assembler/code"
	"github.com/natalieparellano/assembler/parser"
)

func main() {
	path := os.Args[1]
	fmt.Printf("Parsing file: %s\n", path)

	instructions := parser.Parse(path) // TODO: determine the right pattern for sharing structs across packages
	var ret string

	for i := 0; i < len(instructions); i++ {
		instruction := instructions[i]
		var res string
		if instruction.Class == "L" {
			continue // TODO: implement labels
		} else if instruction.Class == "A" {
			val, err := strconv.Atoi(instruction.Symbol)
			check(err)
			res = "0" + zeroPad(fmt.Sprintf("%b", val)) + "\n"
		} else if instruction.Class == "C" {
			res = "111" + code.Comp(instruction.Comp) + code.Dest(instruction.Dest) +
				code.Jump(instruction.Jump) + "\n"
		}
		ret += res
	}

	newpath := newPath(path)
	writeFile(newpath, ret)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func newPath(path string) string {
	fmt.Printf("path: %s\n", path)

	dir := filepath.Dir(path)
	fmt.Printf("dir: %s\n", dir)

	oldfile := filepath.Base(path)
	fmt.Printf("oldfile: %s\n", oldfile)

	newfile := strings.Replace(oldfile, "asm", "hack", 1)
	fmt.Printf("newfile: %s\n", newfile)

	newpath := filepath.Join(dir, newfile)
	fmt.Printf("newpath: %s\n", newpath)

	return newpath
}

func writeFile(path, str string) {
	f, err := os.Create(path)
	check(err)

	_, err = f.WriteString(str)
	check(err)

	err = f.Sync()
	check(err)

	err = f.Close()
	check(err)
}

func zeroPad(val string) string {
	for len(val) < 15 {
		val = "0" + val
	}
	return val
}
