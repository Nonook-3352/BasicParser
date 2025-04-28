package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var REG [1024]int

func ADD(a int, b int, reg int) int {
	REG[reg] = a + b
	return a + b
}

func SUB(a int, b int, reg int) int {
	REG[reg] = a - b
	return REG[reg]
}

func MULT(a int, b int, reg int) int {
	REG[reg] = a * b
	return REG[reg]
}

func DIV(a int, b int, reg int) int {
	REG[reg] = a / b
	return REG[reg]
}

func ECHO(reg int) {
	fmt.Println(int(REG[reg]))
}

func strToInt(str string) (int) {
	strSlice := strings.Split(str, "")
	intSlice := make([]int, len(strSlice))
	m := make(map[string]int)

	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	m["4"] = 4
	m["5"] = 5
	m["6"] = 6
	m["7"] = 7
	m["8"] = 8
	m["9"] = 9

	for i := range len(strSlice) {
		var x int = m[strSlice[i]]
		intSlice[i] = x
	}

	var number int 

	for y := range len(intSlice) {
		if y == 0 {
			number = intSlice[y]
		} else {
			number = number*10 + intSlice[y]
		}
	}
	return number
}

func parse(bytes []byte) [][]byte {
	var buffer []byte
	var Pbuffer [][]byte
	bytes = append(bytes, 32) // Append 32 at the end (space) so the last array of value is appended to Pbuffer

	for _, value := range bytes {
		if value != 32 {
			buffer = append(buffer, value)
		}

		if value == 32 {
			Pbuffer = append(Pbuffer, buffer)
			buffer = []byte{}
		}
	}

	return Pbuffer
}

func main() {

	file, err := os.Open("code.txt")
	

	r := bufio.NewReaderSize(file, 4096)

	for {
		line, _, err := r.ReadLine()
		if len(line) != 0 {
			parsedLine := parse(line) // Call parse once and store the result
			for x, value := range parsedLine {
   				 // x is the index, value is a []byte (one of the sub-slices)
				if x == 0 {
					switch string(value) {
					case "ADD":
						a := strToInt(string(parsedLine[1]))
						b := strToInt(string(parsedLine[2]))
						reg := strToInt(string(parsedLine[3]))
						ADD(a, b, reg)
						
					case "SUB":
						a := strToInt(string(parsedLine[1]))
						b := strToInt(string(parsedLine[2]))
						reg := strToInt(string(parsedLine[3]))
						SUB(a, b, reg)
						
					case "MULT":
						a := strToInt(string(parsedLine[1]))
						b := strToInt(string(parsedLine[2]))
						reg := strToInt(string(parsedLine[3]))
						MULT(a, b, reg)
					
					case "DIV":
						a := strToInt(string(parsedLine[1]))
						b := strToInt(string(parsedLine[2]))
						reg := strToInt(string(parsedLine[3]))
						DIV(a, b, reg)
					
					case "ECHO":
						reg := strToInt(string(parsedLine[1]))
						ECHO(reg)
					}
				}
			}
		} 

		if err != nil {
			break
		}
	}

	if err != nil {
		fmt.Println(err)
	}
}
