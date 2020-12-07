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

type rule struct {
	color      string
	conditions map[string]int64
}

func createRule(input string) rule {
	input = strings.TrimSpace(input)

	rgColor := regexp.MustCompile("(\\w+ \\w+) bags contain (.*)")
	rgCondition := regexp.MustCompile("(\\d) (\\w+ \\w+) bag[s]?")
	rgNoBags := regexp.MustCompile("no other bags")

	match := rgColor.FindStringSubmatch(input)
	if len(match) < 2 {
		panic("match res<2! input: " + input)
	}

	newRule := rule{
		color:      match[1],
		conditions: map[string]int64{},
	}

	if rgNoBags.MatchString(match[2]) {
		fmt.Printf("%+v\n", newRule)
		return newRule
	}
	for _, c := range strings.Split(match[2], ", ") {
		conditionMatch := rgCondition.FindStringSubmatch(c)
		if len(conditionMatch) < 2 {
			panic("conditionMatch res<2! input: " + input)
		}
		num, err := strconv.ParseInt(conditionMatch[1], 10, 64)
		if err != nil {
			panic(err)
		}

		newRule.conditions[conditionMatch[2]] = num
	}

	fmt.Printf("%+v\n", newRule)
	return newRule
}

func readFile(filename string) []rule {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rules := []rule{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rules = append(rules, createRule(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return rules
}

func getRuleForColor(color string, rules []rule) *rule {
	for _, rule := range rules {
		if rule.color == color {
			return &rule
		}
	}
	return nil
}

func conditionsContainColor(color string, r rule, rules []rule) bool {
	for col := range r.conditions {
		subRule := getRuleForColor(col, rules)
		_, ok := subRule.conditions[color]
		if ok {
			return true
		}
	}
	return false
}

func task1(rules []rule) int {
	myBagColor := "shiny gold"

	colorsContainMyBag := map[string]bool{}
	for _, rule := range rules {
		if len(rule.conditions) == 0 {
			continue
		}

		_, ok := rule.conditions[myBagColor]
		if ok {
			colorsContainMyBag[rule.color] = true
		}
	}

	colorsContainMyBagOldSize := 0
	for {
		for _, rule := range rules {
			if len(rule.conditions) == 0 {
				continue
			}

			for col := range rule.conditions {
				if colorsContainMyBag[col] {
					colorsContainMyBag[rule.color] = true
				}
			}
		}
		if colorsContainMyBagOldSize == len(colorsContainMyBag) {
			break
		}

		colorsContainMyBagOldSize = len(colorsContainMyBag)
	}

	return colorsContainMyBagOldSize
}

func bagContainsInside(color string, rules []rule, knownBagsCount map[string]int64) int64 {
	num, ok := knownBagsCount[color]
	if ok {
		//fmt.Printf("%v:%v\n", color, num)
		return num
	}

	for _, rule := range rules {
		if rule.color == color {
			if len(rule.conditions) == 0 {
				knownBagsCount[rule.color] = 0
				//fmt.Printf("%v:%v\n", rule.color, 0)
				return 0
			}
			num := int64(0)
			for col, count := range rule.conditions {
				num += count * (bagContainsInside(col, rules, knownBagsCount) + 1)
			}
			knownBagsCount[rule.color] = num
			//fmt.Printf("%v:%v\n", rule.color, num)
			return num
		}
	}
	panic("not found color: " + color)
}

func task2(rules []rule) int64 {
	myBagColor := "shiny gold"
	num := bagContainsInside(myBagColor, rules, map[string]int64{})
	return num
}

// task1: 177
// task2: 34988

func main() {
	answers := readFile("input.txt")
	fmt.Printf("Task1: %v\n", task1(answers))
	fmt.Printf("Task2: %v\n", task2(answers))
}
