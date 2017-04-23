package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/natalieparellano/assembler/symboltable"
)

type Instruction struct {
	Class  string
	Symbol string
	Dest   string
	Comp   string
	Jump   string
}

var nextAddress int = 16

func Parse(filepath string) []Instruction {
	f, err := os.Open(filepath)
	check(err)

	scanner := bufio.NewScanner(f)
	fmt.Printf("\nFIRST PASS\n\n")
	firstPass(scanner)

	f.Seek(0, 0)

	scanner = bufio.NewScanner(f)
	fmt.Printf("\nSECOND PASS\n\n")
	return secondPass(scanner)
}

// Helper Methods

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isAInstruction(line string) bool {
	re := regexp.MustCompile(`@\d+`)
	res := re.FindStringSubmatch(line)
	return res != nil
}

func isLabel(line string) bool {
	return strings.HasPrefix(line, "(")
}

func isSymbol(line string) bool {
	re := regexp.MustCompile(`@\D+`)
	res := re.FindStringSubmatch(line)
	return res != nil
}

func isWhitespace(line string) bool {
	return len(strings.TrimSpace(line)) == 0
}

func trimLine(line string) string {
	line = strings.Split(line, "//")[0]
	line = strings.TrimSpace(line)
	return line
}

// First pass

func firstPass(scanner *bufio.Scanner) {
	var line string
	count := 0

	for scanner.Scan() {
		fmt.Printf("Parsing: %s\n", scanner.Text())
		line = trimLine(scanner.Text())
		if isWhitespace(line) {
			continue
		}
		if isLabel(line) {
			label := line[1 : len(line)-1]
			processLabel(label, count)
			continue
		}
		count += 1 // Line is A or C instruction
	}
	fmt.Printf("Found %d lines of code\n", count)
}

func processLabel(label string, count int) {
	if !symboltable.Contains(label) {
		symboltable.AddEntry(label, count)
	}
}

// Second pass

func secondPass(scanner *bufio.Scanner) []Instruction {
	ret := make([]Instruction, 1)

	var line string
	var instruction Instruction

	for scanner.Scan() {
		fmt.Printf("Parsing: %s\n", scanner.Text())
		line = trimLine(scanner.Text())
		if isWhitespace(line) {
			continue
		}
		if isLabel(line) {
			continue
		}
		instruction = parseLine(line)
		ret = append(ret, instruction)
	}

	return ret
}

func parseLine(line string) Instruction {
	var instruction Instruction
	if isSymbol(line) {
		instruction.Class = "A"
		symbol := line[1:]
		instruction.Symbol = parseSymbol(symbol)
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

func parseSymbol(symbol string) string {
	var address int
	if symboltable.Contains(symbol) {
		address = symboltable.GetAddress(symbol)
	} else {
		address = nextAddress
		symboltable.AddEntry(symbol, address)
		nextAddress += 1
	}
	return strconv.Itoa(address)
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
