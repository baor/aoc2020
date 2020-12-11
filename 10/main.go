package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type adapters sort.IntSlice

func readFile(filename string) adapters {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ads := adapters{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		joltage := int(value)
		ads = append(ads, joltage)
	}

	sort.Sort(sort.IntSlice(ads))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	jl := getDeviceJoltage(ads)
	ads = append(ads, jl)

	return ads
}

func getDeviceJoltage(ads adapters) int {
	return ads[len(ads)-1] + 3
}

var deviceJoltage int

func getAdapterForJoltage(jl int, ads adapters) []int {
	validAds := []int{}
	for _, joltage := range ads {
		if joltage <= jl {
			continue
		}

		if joltage > jl+3 {
			break
		}
		validAds = append(validAds, joltage)
	}

	return validAds
}

type connections []int

type nodes map[int]connections

func (n nodes) print(ads adapters) {
	fmt.Println("nodes:")
	for _, jl := range ads {
		fmt.Printf("%v:%v\n", jl, n[jl])
	}
}

func buildNodes(ads adapters) nodes {
	n := nodes{}

	for i := 0; i < len(ads); i++ {
		n[ads[i]] = connections{}
	}

	for i := 0; i < len(ads)-1; i++ {
		for j := i + 1; j < len(ads); j++ {
			if ads[j]-ads[i] > 3 {
				break
			}
			n[ads[i]] = append(n[ads[i]], ads[j])
		}
	}

	return n
}

func getValidChains(ads adapters) int {
	nodes := buildNodes(ads)
	nodes.print(ads)
	validChains := map[int]int{}
	validChainsCount := 1
	for i := len(ads) - 1; i >= 0; i-- {
		jl := ads[i]
		connects := nodes[jl]
		if len(connects) < 2 {
			validChains[jl] = validChainsCount
			fmt.Printf("validChainsCount: %v for %v\n", validChainsCount, jl)
			continue
		}
		validChainsCount = 0
		for _, con := range connects {
			validChainsCount += validChains[con]
		}
		validChains[jl] = validChainsCount
		fmt.Printf("validChainsCount: %v for %v\n", validChainsCount, jl)
	}

	return validChainsCount
}

func task1(ads adapters) int {
	num1 := 0
	num3 := 0
	fmt.Printf("ads: %v\n", ads)

	for i := 0; i < len(ads); i++ {
		diff := ads[i]
		if i > 0 {
			diff = ads[i] - ads[i-1]
		}
		//fmt.Printf("ch[i]: %v, diff: %v\n", ch[i], diff)

		if diff == 1 {
			num1++
			continue
		}
		if diff == 3 {
			num3++
			continue
		}
	}
	fmt.Printf("num1: %v, num3:%v\n", num1, num3)
	return num1 * num3
}

func task2(ads adapters) int {
	fmt.Printf("len ad: %v\n", len(ads))
	fmt.Printf("ads: %v\n", ads)

	ads = append(adapters{0}, ads...)

	// fmt.Printf("blocks: %v\n", blocks)
	validChainsCount := getValidChains(ads)

	fmt.Printf("valid chains: %v\n", validChainsCount)
	return validChainsCount
}

// task1: 3000
// task2: 193434623148032

func main() {
	answers := readFile("input.txt")
	fmt.Printf("Task1: %v\n", task1(answers))
	fmt.Printf("Task2: %v\n", task2(answers))
}
