package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/natalieparellano/assembler/code"
	"github.com/natalieparellano/assembler/hackfile"
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
			continue
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

	newpath := hackfile.NewPath(path, "asm", "hack")
	hackfile.WriteFile(newpath, ret)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func zeroPad(val string) string {
	for len(val) < 15 {
		val = "0" + val
	}
	return val
}
