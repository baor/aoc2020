package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := []string{}
	for scanner.Scan() {
		data = append(data, parseLine(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("data:%v\n", data)

	return data
}

func toRPLTask1(input string, index int) (string, int) {
	out := ""
	op := ""
	for i := index; i < len(input); i++ {
		//fmt.Printf("i: %v, op:%v, out: %v\n", i, op, out)
		ch := string(input[i])
		switch ch {
		case "+":
			op = "+"
		case "*":
			op = "*"
		case ")":
			return out, i
		case "(":
			sub, iNew := toRPLTask1(input, i+1)
			i = iNew
			ch = sub
			fallthrough
		default:
			if op == "" {
				out = fmt.Sprintf("%v%v", out, ch)
			} else {
				out = fmt.Sprintf("%v%v%v", out, ch, op)
				op = ""
			}
		}
	}

	return out, len(input)
}

//var loop int

func evaluateRPL(input string, index int, arg1 int, arg2 int) (int, int) {

	//fmt.Printf("call!")
	for i := index; i < len(input); i++ {
		//loop++
		// if loop > 20000 {
		// 	panic("endless loop!")
		// }
		ch := string(input[i])
		// fmt.Printf("evaluateRPL. i: %v, ch:%v, arg1: %v, arg2:%v\n", i, ch, arg1, arg2)

		switch ch {
		case "+":
			arg1 += arg2
			arg2 = 0
			if index != 0 {
				return arg1, i
			}
		case "*":
			arg1 *= arg2
			arg2 = 0
			if index != 0 {
				return arg1, i
			}
		default:
			v, err := strconv.ParseInt(ch, 10, 64)
			val := int(v)
			if err != nil {
				panic(err)
			}
			if arg1 == 0 {
				arg1 = val
			} else if arg2 == 0 {
				arg2 = val
			} else {
				newVal, iNew := evaluateRPL(input, i+1, arg2, val)

				arg2 = newVal
				//fmt.Printf("back. i:%v, arg1: %v, arg2:%v\n", iNew, arg1, arg2)

				i = iNew
				//loop++
			}
		}
	}
	return arg1, 0
}

func parseLine(input string) string {
	fmt.Printf("parseLine %v\n", input)
	return strings.Join(strings.Split(input, " "), "")
}

func evaluateTask1(expression string) int {
	rpl, _ := toRPLTask1(expression, 0)
	fmt.Printf("rpl: %v\n", rpl)
	res, _ := evaluateRPL(rpl, 0, 0, 0)
	return res
}

func task1(lines []string) int {
	sum := 0
	for _, l := range lines {
		res := evaluateTask1(l)
		sum += res
	}
	return sum
}

func toRPLTask2(input string, index int) (string, int) {
	out := ""
	op := ""
	keepOp := ""
	for i := index; i < len(input); i++ {
		// fmt.Printf("i: %v, op:%v, bufOp:%v, out: %v\n", i, op, keepOp, out)
		ch := string(input[i])
		switch ch {
		case "+":
			if op == "" {
				op = "+"
			} else if op == "+" {
				out = fmt.Sprintf("%v%v", out, op)
			} else if op == "*" {
				keepOp = "*"
				op = "+"
			}
		case "*":
			if keepOp != "" {
				out = fmt.Sprintf("%v%v", out, keepOp)
				keepOp = ""
			}
			if op == "" {
				op = "*"
			} else if op == "*" {
				out = fmt.Sprintf("%v%v", out, op)
			} else {
				panic("strange op:" + ch)
			}
		case ")":
			if op != "" {
				out = fmt.Sprintf("%v%v", out, op)
			}
			if keepOp != "" {
				out = fmt.Sprintf("%v%v", out, keepOp)
			}
			//fmt.Printf("return out: %v\n", out)
			return out, i
		case "(":
			sub, iNew := toRPLTask2(input, i+1)
			i = iNew
			ch = sub
			fallthrough
		default:
			if op == "" {
				out = fmt.Sprintf("%v%v", out, ch)
			} else if op == "+" {
				out = fmt.Sprintf("%v%v%v", out, ch, op)
				op = ""
			} else if op == "*" {
				out = fmt.Sprintf("%v%v", out, ch)
			}
		}
	}

	if op != "" {
		out = fmt.Sprintf("%v%v", out, op)
		op = ""
	}

	if keepOp != "" {
		out = fmt.Sprintf("%v%v", out, keepOp)
		keepOp = ""
	}

	//fmt.Printf("return out: %v\n", out)
	return out, len(input)
}

func evaluateTask2(expression string) int {
	fmt.Printf("expression: %v\n", expression)
	rpl, _ := toRPLTask2(expression, 0)
	fmt.Printf("rpl: %v\n", rpl)
	res, _ := evaluateRPL(rpl, 0, 0, 0)
	return res
}

func task2(lines []string) int {
	sum := 0
	for _, l := range lines {
		res := evaluateTask2(l)
		sum += res
	}
	return sum
}

// task1: 5783053349377
// task2: 74821486966872

func main() {
	input := readFile("input.txt")
	fmt.Printf("Task1: %v\n", task1(input))
	fmt.Printf("Task2: %v\n", task2(input))
}
