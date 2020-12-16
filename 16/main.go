package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type interval struct {
	from int
	to   int
}

type rule struct {
	field string
	i1    interval
	i2    interval
}

type ticket []int

type data struct {
	rules         []rule
	myTicket      ticket
	nearbyTickets []ticket
	validTickets  []ticket
}

func createTicket(input string) ticket {
	t := ticket{}
	nums := strings.Split(input, ",")
	for _, n := range nums {
		v, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			panic(err)
		}
		t = append(t, int(v))
	}
	return t
}

func readFile(filename string) data {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	d := data{rules: []rule{}, nearbyTickets: []ticket{}}

	rgRule := regexp.MustCompile("([^:]+): ([\\d]+)-([\\d]+) or ([\\d]+)-([\\d]+)")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		ruleMatch := rgRule.FindStringSubmatch(line)
		if len(ruleMatch) < 6 {
			panic("len(ruleMatch) < 6 ")
		}

		v2, err := strconv.ParseInt(ruleMatch[2], 10, 64)
		if err != nil {
			panic(err)
		}
		v3, err := strconv.ParseInt(ruleMatch[3], 10, 64)
		if err != nil {
			panic(err)
		}
		v4, err := strconv.ParseInt(ruleMatch[4], 10, 64)
		if err != nil {
			panic(err)
		}
		v5, err := strconv.ParseInt(ruleMatch[5], 10, 64)
		if err != nil {
			panic(err)
		}
		i1 := interval{
			from: int(v2),
			to:   int(v3),
		}
		i2 := interval{
			from: int(v4),
			to:   int(v5),
		}

		d.rules = append(d.rules, rule{field: ruleMatch[1], i1: i1, i2: i2})
	}

	// scan my ticket
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if strings.Contains(line, "your ticket") {
			continue
		}
		d.myTicket = createTicket(line)
	}

	// scan nearby tickets
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if strings.Contains(line, "nearby tickets") {
			continue
		}
		d.nearbyTickets = append(d.nearbyTickets, createTicket(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("data:%+v\n", d)
	fmt.Printf("rules[%v]:%+v\n", len(d.rules), d.rules)
	return d
}

func checkErrors(d *data) int {
	errorRate := 0
	for _, ticket := range d.nearbyTickets {
		isValidTicket := true
		for _, ticketNumber := range ticket {
			isValidNumber := false
			for _, r := range d.rules {
				if (ticketNumber >= r.i1.from && ticketNumber <= r.i1.to) ||
					(ticketNumber >= r.i2.from && ticketNumber <= r.i2.to) {
					isValidNumber = true
					break
				}
			}
			if !isValidNumber {
				fmt.Printf("wrong number: %v\n", ticketNumber)
				errorRate += ticketNumber
				isValidTicket = false
				break
			}
		}
		if isValidTicket {
			d.validTickets = append(d.validTickets, ticket)
		}
	}
	return errorRate
}

func task1(d data) int {
	return checkErrors(&d)
}

func identifyMapping(d data) map[string]int {
	d.validTickets = []ticket{}
	checkErrors(&d)

	d.validTickets = append(d.validTickets, d.myTicket)
	fmt.Printf("valid tickets len: %v\n", len(d.validTickets))

	possibleMapping := map[int]map[string]bool{}
	//knownIndexes := map[int]bool{}

	for _, t := range d.validTickets {
		for ticketNumberIndex, ticketNumber := range t {
			if possibleMapping[ticketNumberIndex] == nil {
				possibleMapping[ticketNumberIndex] = map[string]bool{}
				for _, r := range d.rules {
					possibleMapping[ticketNumberIndex][r.field] = true
				}
			}

			for _, r := range d.rules {
				if (ticketNumber >= r.i1.from && ticketNumber <= r.i1.to) ||
					(ticketNumber >= r.i2.from && ticketNumber <= r.i2.to) {
					continue
				}
				possibleMapping[ticketNumberIndex][r.field] = false
			}
		}
	}

	fmt.Printf("possibleMapping: %v\n", possibleMapping)

	knownMapping := map[string]int{}
	for i := 0; i < len(d.rules); i++ {
		for index, fields := range possibleMapping {
			numTrue := 0
			fieldTrueName := ""
			for f, v := range fields {
				if _, ok := knownMapping[f]; ok {
					continue
				}
				if v {
					fieldTrueName = f
					numTrue++
				}
			}
			if numTrue == 1 {
				knownMapping[fieldTrueName] = index
			}
		}
		//fmt.Printf("knownFields: %v\n", knownMapping)
	}

	return knownMapping
}

func task2(d data) int {
	res := 1
	m := identifyMapping(d)
	fmt.Printf("mapping: %v\n", m)
	for fieldName, index := range m {
		if strings.Index(fieldName, "departure") >= 0 {
			res *= d.myTicket[index]
		}
	}
	return res
}

// task1: 25972
// task2: 622670335901

func main() {
	input := readFile("input.txt")
	fmt.Printf("Task1: %v\n", task1(input))
	fmt.Printf("Task2: %v\n", task2(input))
}
