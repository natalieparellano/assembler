package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Instruction struct {
	Class  string
	Symbol string
	Dest   string
	Comp   string
	Jump   string
}

func Parse(filepath string) []Instruction {
	ret := make([]Instruction, 1)

	f, err := os.Open(filepath)
	check(err)

	scanner := bufio.NewScanner(f)
	var line string
	var instruction Instruction

	for scanner.Scan() {
		fmt.Printf("Parsing: %s\n", scanner.Text())
		line = trimLine(scanner.Text())
		if isWhitespace(line) {
			continue
		}
		instruction = parseLine(line)
		ret = append(ret, instruction)
	}

	return ret
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isAInstruction(line string) bool {
	return strings.HasPrefix(line, "@")
}

func isLabel(line string) bool {
	return strings.HasPrefix(line, "(")
}

func isWhitespace(line string) bool {
	return len(strings.TrimSpace(line)) == 0
}

func parseCInstruction(line string) (string, string, string) {
	var dest, comp, jump string

	if strings.Contains(line, "=") && strings.Contains(line, ";") {
		re := regexp.MustCompile(`(.+)=(.+);(.+)`)
		res := re.FindStringSubmatch(line)
		dest = res[1]
		comp = res[2]
		jump = res[3]
	} else if strings.Contains(line, "=") {
		re := regexp.MustCompile(`(.+)=(.+)`)
		res := re.FindStringSubmatch(line)
		dest = res[1]
		comp = res[2]
	} else {
		re := regexp.MustCompile(`(.+);(.+)`)
		res := re.FindStringSubmatch(line)
		comp = res[1]
		jump = res[2]
	}

	return strings.TrimSpace(dest), strings.TrimSpace(comp), strings.TrimSpace(jump)
}

func parseLine(line string) Instruction {
	var instruction Instruction
	if isLabel(line) {
		instruction.Class = "L"
		instruction.Symbol = line[1 : len(line)-1]
	} else if isAInstruction(line) {
		instruction.Class = "A"
		instruction.Symbol = line[1:]
	} else {
		instruction.Class = "C"
		dest, comp, jump := parseCInstruction(line)
		instruction.Dest = dest
		instruction.Comp = comp
		instruction.Jump = jump
	}
	return instruction
}

func trimLine(line string) string {
	line = strings.Split(line, "//")[0]
	line = strings.TrimSpace(line)
	return line
}
