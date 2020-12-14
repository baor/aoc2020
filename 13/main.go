package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseSchedule(input string) []int {
	busses := []int{}
	schedule := strings.Split(input, ",")
	for _, s := range schedule {
		if s == "x" {
			busses = append(busses, 0)
			continue
		}

		val, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}

		busses = append(busses, int(val))
	}
	return busses
}

func readFile(filename string) (int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	val, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	timeOnStation := int(val)

	scanner.Scan()
	busses := parseSchedule(scanner.Text())

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("time:%v\nbusses:%v\n", timeOnStation, busses)
	return timeOnStation, busses
}

func getWaitTime(myTime int, bus int) int {
	div := myTime / bus
	if div == 0 {
		return bus - myTime
	}

	reminder := myTime % bus
	if reminder == 0 {
		return 0
	}

	return (div+1)*bus - myTime
}

func task1(time int, busses []int) int {
	bestWaitTime := 1000000
	bestBus := 0
	for _, b := range busses {
		if b == 0 {
			continue
		}
		waitTime := getWaitTime(time, b)
		if waitTime < bestWaitTime {
			fmt.Printf("Best result: bus: %v, wait time: %v\n", b, waitTime)
			bestWaitTime = waitTime
			bestBus = b
		}
	}
	return bestBus * bestWaitTime
}

func isK0Good(k0 int, busses []int) bool {
	for i := 1; i < len(busses); i++ {
		if busses[i] == 0 {
			continue
		}

		reminderCurrent := (i + busses[0]*k0) % busses[i]
		if reminderCurrent != 0 {
			return false
		}
	}
	return true
}

func isPrime(num int) bool {
	for i := 1; i < num; i++ {
		if i%num == 0 {
			return false
		}
	}
	return true
}

func getNod(num int) int {
	if num == 1 {
		return 1
	}
	for i := 2; i <= num; i++ {
		if num%i == 0 {
			return i
		}
	}
	return 0
}

func task2(busses []int, kInit int) int {
	sortedBusses := sort.IntSlice{}
	busIndexMap := map[int]int{}
	for i := 0; i < len(busses); i++ {
		if busses[i] == 0 {
			continue
		}

		if !isPrime(busses[i]) {
			panic("number is not prime!")
		}

		sortedBusses = append(sortedBusses, busses[i])
		busIndexMap[busses[i]] = i
		fmt.Printf("index: %v\n", i)
	}
	sort.Sort(sortedBusses)
	fmt.Printf("sortedBusses: %v\n", sortedBusses)

	k := kInit
	kStorage := map[int]int{}
	diffMap := map[int]int{}
	idBusKnown := sortedBusses[len(sortedBusses)-1]

	fmt.Printf("idBusKnown: %v\n", idBusKnown)

	for i := 0; i < len(busses); i++ {
		diffMap[busses[i]] = i - busIndexMap[idBusKnown]
	}

	fmt.Printf("diffMap: %v\n", diffMap)

	multipliers := []int{}
	for _, diff := range diffMap {
		if diff == 0 || diff == 1 || diff == -1 {
			continue
		}
		diff = int(math.Abs(float64(diff)))
		for i := 0; i < len(busses); i++ {
			if busses[i] == 0 || busses[i] == 1 {
				continue
			}
			if busses[i]%diff == 0 {
				fmt.Printf("found multiplier %v!!!\n", diff)
				multipliers = append(multipliers, diff)
			}
		}
	}

	multiplier := 1
	if len(multipliers) > 0 {
		for _, m := range multipliers {
			multiplier *= m
		}
	}
	k = multiplier
	for {
		// start from rarest busses
		kStorage[idBusKnown] = k
		addition := idBusKnown * k
		allGood := true
		for i := len(sortedBusses) - 2; i >= 0; i-- {
			idBus := sortedBusses[i]
			timestampWithShift := diffMap[idBus] + addition
			if timestampWithShift%idBus != 0 {
				allGood = false
				break
			}
			kStorage[idBus] = timestampWithShift / idBus
		}
		if allGood {
			fmt.Printf("final k:%v\n", k)
			return kStorage[busses[0]] * busses[0]
		}
		if k%1000000000 == 0 {
			fmt.Printf("%v\n", k)
		}
		k += multiplier
	}
	//return 0
}

// task1: 2215
// task2: 1058443396696792

func main() {
	time, busses := readFile("input.txt")
	fmt.Printf("Task1: %v\n", task1(time, busses))
	fmt.Printf("Task2: %v\n", task2(busses, 0))
}
