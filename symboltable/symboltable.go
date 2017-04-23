package symboltable

import "fmt"

var symbolDict = map[string]int{
	"SP":     0,
	"LCL":    1,
	"ARG":    2,
	"THIS":   3,
	"THAT":   4,
	"R0":     0,
	"R1":     1,
	"R2":     2,
	"R3":     3,
	"R4":     4,
	"R5":     5,
	"R6":     6,
	"R7":     7,
	"R8":     8,
	"R9":     9,
	"R10":    10,
	"R11":    11,
	"R12":    12,
	"R13":    13,
	"R14":    14,
	"R15":    15,
	"SCREEN": 16384,
	"KBD":    24576,
}

func Contains(key string) bool {
	_, exists := symbolDict[key]
	return exists
}

func AddEntry(key string, val int) {
	fmt.Printf("Adding key %s with value %d\n", key, val)
	fmt.Printf("Table has old size %d\n", len(symbolDict))
	symbolDict[key] = val
	fmt.Printf("Table has new size %d\n", len(symbolDict))
}

func GetAddress(key string) int {
	fmt.Printf("Searching for %s\n", key)
	val, exists := symbolDict[key]
	if !exists {
		panic(fmt.Sprintf("Could not find %s in table", key))
	}
	return val
}
