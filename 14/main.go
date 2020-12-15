package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type write struct {
	address uint
	value   uint
}

type blockTask1 struct {
	forceZero uint
	forceOne  uint
	writes    []write
}

type programTask1 struct {
	blocks []blockTask1
	mem    map[uint]uint
}

type forcePair struct {
	zero uint
	one  uint
}
type blockTask2 struct {
	forces []forcePair
	writes []write
}

type programTask2 struct {
	blocks []blockTask2
	mem    map[uint]uint
}

func parseMaskTask1(input string) (uint, uint) {
	forceZero := uint(0)
	forceOne := uint(0)

	for i := 0; i < len(input); i++ {

		forceZero <<= 1
		forceOne <<= 1

		if string(input[i]) == "X" {
			continue
		}

		v, err := strconv.ParseInt(string(input[i]), 10, 64)
		if err != nil {
			panic(err)
		}
		if v == 0 {
			forceZero |= 1
			continue
		}
		if v == 1 {
			forceOne |= 1
			continue
		}
		panic("unknown v!")
	}
	fmt.Printf("input: %v\n", input)
	fmt.Printf("0: %b\n", forceZero)
	fmt.Printf("1: %b\n", forceOne)
	return forceZero, forceOne
}

func parseMaskTask2(input string) []forcePair {
	pairs := []forcePair{
		forcePair{
			zero: uint(0),
			one:  uint(0),
		},
	}

	for i := 0; i < len(input); i++ {
		for i := range pairs {
			pairs[i].zero <<= 1
			pairs[i].one <<= 1
		}

		if string(input[i]) == "X" {
			size := len(pairs)
			for i := 0; i < size; i++ {
				pairs = append(pairs, forcePair{
					zero: pairs[i].zero,
					one:  pairs[i].one | 1,
				})
				pairs[i].zero |= 1

			}
			continue
		}

		v, err := strconv.ParseInt(string(input[i]), 10, 64)
		if err != nil {
			panic(err)
		}
		if v == 0 {
			continue
		}
		if v == 1 {
			for i := range pairs {
				pairs[i].one |= 1
			}
			continue
		}
		panic("unknown v!")
	}
	fmt.Printf("input: %v\n", input)
	for i := range pairs {
		fmt.Printf("0: %b\n", pairs[i].zero)
		fmt.Printf("1: %b\n", pairs[i].one)
	}
	return pairs
}

func processBlockTask1(b blockTask1, mem map[uint]uint) {
	for _, w := range b.writes {
		//fmt.Printf("bef--: %b\n", w.value)
		val := w.value | b.forceOne
		//fmt.Printf("aft-1: %b\n", val)
		val = val ^ (val & b.forceZero)
		//fmt.Printf("aft-0: %b\n", val)
		mem[w.address] = val
	}
}

func processBlockTask2(b blockTask2, mem map[uint]uint) {
	for _, w := range b.writes {
		//fmt.Printf("bef--: %b\n", w.value)
		for _, p := range b.forces {
			// fmt.Printf("address-before: %b\n", w.address)
			address := w.address | p.one
			// fmt.Printf("aft-1: %b\n", address)
			address = address ^ (address & p.zero)
			// fmt.Printf("aft-0: %b\n", address)
			// fmt.Printf("address: %b\n", address)
			// fmt.Printf("mem[%v]=%v\n", address, w.value)
			mem[address] = w.value
		}
	}
}

func readFileTask1(filename string) programTask1 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	prog := programTask1{}
	prog.blocks = []blockTask1{}
	prog.mem = map[uint]uint{}

	b := blockTask1{}
	rgMask := regexp.MustCompile("mask = ([10X]+)$")
	rgWrite := regexp.MustCompile("mem\\[([0-9]+)\\] = ([0-9]+)")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if rgMask.MatchString(line) {
			if len(b.writes) > 0 {
				prog.blocks = append(prog.blocks, b)
			}
			maskMatch := rgMask.FindStringSubmatch(line)

			f0, f1 := parseMaskTask1(maskMatch[1])
			b = blockTask1{
				forceZero: f0,
				forceOne:  f1,
			}
			//fmt.Printf("b mask:%v\n", b)
			continue
		}
		if rgWrite.MatchString(line) {
			writeMatch := rgWrite.FindStringSubmatch(line)

			if len(writeMatch) < 3 {
				panic("len(writeMatch) <3")
			}
			a, err := strconv.ParseUint(writeMatch[1], 10, 64)
			if err != nil {
				panic(err)
			}
			v, err := strconv.ParseUint(writeMatch[2], 10, 64)
			if err != nil {
				panic(err)
			}
			b.writes = append(b.writes, write{address: uint(a), value: uint(v)})
			//fmt.Printf("b writes:%v\n", b)
			continue
		}
	}
	prog.blocks = append(prog.blocks, b)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("prog:%+v\n", prog)
	return prog
}

func readFileTask2(filename string) programTask2 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	prog := programTask2{}
	prog.blocks = []blockTask2{}
	prog.mem = map[uint]uint{}

	b := blockTask2{}
	rgMask := regexp.MustCompile("mask = ([10X]+)$")
	rgWrite := regexp.MustCompile("mem\\[([0-9]+)\\] = ([0-9]+)")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if rgMask.MatchString(line) {
			if len(b.writes) > 0 {
				prog.blocks = append(prog.blocks, b)
			}
			maskMatch := rgMask.FindStringSubmatch(line)

			fPairs := parseMaskTask2(maskMatch[1])
			b = blockTask2{
				forces: fPairs,
			}
			//fmt.Printf("b mask:%v\n", b)
			continue
		}
		if rgWrite.MatchString(line) {
			writeMatch := rgWrite.FindStringSubmatch(line)

			if len(writeMatch) < 3 {
				panic("len(writeMatch) <3")
			}
			a, err := strconv.ParseUint(writeMatch[1], 10, 64)
			if err != nil {
				panic(err)
			}
			v, err := strconv.ParseUint(writeMatch[2], 10, 64)
			if err != nil {
				panic(err)
			}
			b.writes = append(b.writes, write{address: uint(a), value: uint(v)})
			//fmt.Printf("b writes:%v\n", b)
			continue
		}
	}
	prog.blocks = append(prog.blocks, b)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("prog:%+v\n", prog)
	return prog
}

func task1(prog programTask1) uint {
	for _, b := range prog.blocks {
		processBlockTask1(b, prog.mem)
	}
	sum := uint(0)
	for _, val := range prog.mem {
		sum += val
	}
	return sum
}

func task2(prog programTask2) uint {
	for _, b := range prog.blocks {
		processBlockTask2(b, prog.mem)
	}
	sum := uint(0)
	for _, val := range prog.mem {
		sum += val
	}
	return sum
}

// task1: 15403588588538
// task2: 3260587250457

func main() {
	prog1 := readFileTask1("input.txt")
	fmt.Printf("Task1: %v\n", task1(prog1))
	prog2 := readFileTask2("input.txt")
	fmt.Printf("Task2: %v\n", task2(prog2))
}
