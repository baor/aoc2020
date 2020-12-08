package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func createByr(input string) (string, error) {
	var value int
	fmt.Sscanf(input, "%d", &value)
	if value >= 1920 && value <= 2002 {
		return input, nil
	}
	return "", fmt.Errorf("%v is invalid", input)
}

func createIyr(input string) (string, error) {
	var value int
	fmt.Sscanf(input, "%d", &value)
	if value >= 2010 && value <= 2020 {
		return input, nil
	}
	return "", fmt.Errorf("%v is invalid", input)
}

func createEyr(input string) (string, error) {
	var value int
	fmt.Sscanf(input, "%d", &value)
	if value >= 2020 && value <= 2030 {
		return input, nil
	}
	return "", fmt.Errorf("%v is invalid", input)
}

func createHgt(input string) (string, error) {
	var value int
	var unit string
	fmt.Sscanf(input, "%d%s", &value, &unit)
	if unit == "cm" && value >= 150 && value <= 193 {
		return input, nil
	}

	if unit == "in" && value >= 59 && value <= 76 {
		return input, nil
	}

	return "", fmt.Errorf("%v is invalid", input)
}

func createHcl(input string) (string, error) {
	var value string
	fmt.Sscanf(input, "#%s", &value)
	if len(value) != 6 {
		return "", fmt.Errorf("%v is invalid, length is not 6", input)
	}

	matched, _ := regexp.MatchString("^[0-9a-f]+$", value)
	if matched {
		return input, nil
	}

	return "", fmt.Errorf("%v is invalid", input)
}

func createEcl(input string) (string, error) {
	switch input {
	case "amb":
		fallthrough
	case "blu":
		fallthrough
	case "brn":
		fallthrough
	case "gry":
		fallthrough
	case "grn":
		fallthrough
	case "hzl":
		fallthrough
	case "oth":
		return input, nil
	}

	return "", fmt.Errorf("%v is invalid", input)
}

func createPid(input string) (string, error) {
	matched, _ := regexp.MatchString("^[0-9]{9}$", input)
	if matched {
		return input, nil
	}

	return "", fmt.Errorf("%v is invalid", input)
}

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func createPassportTask1(input string) (passport, error) {
	input = strings.ReplaceAll(input, "\n", " ")
	pairs := strings.Split(input, " ")

	pass := passport{}
	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}

		vals := strings.Split(pair, ":")
		switch vals[0] {
		case "byr":
			pass.byr = vals[1]
		case "iyr":
			pass.iyr = vals[1]
		case "eyr":
			pass.eyr = vals[1]
		case "hgt":
			pass.hgt = vals[1]
		case "hcl":
			pass.hcl = vals[1]
		case "ecl":
			pass.ecl = vals[1]
		case "pid":
			pass.pid = vals[1]
		case "cid":
			pass.cid = vals[1]
		default:
			panic("unknown key \"" + vals[0] + "\" in pair \"" + pair + "\"")
		}
	}

	hasMandatoryFields := pass.byr != "" &&
		pass.ecl != "" &&
		pass.eyr != "" &&
		pass.hcl != "" &&
		pass.hgt != "" &&
		pass.iyr != "" &&
		pass.pid != ""

	if !hasMandatoryFields {
		return pass, fmt.Errorf("Misses mandatory fields")
	}

	return pass, nil
}

func createPassportTask2(input string) (passport, error) {
	input = strings.ReplaceAll(input, "\n", " ")
	pairs := strings.Split(input, " ")

	pass := passport{}
	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}

		var err error
		vals := strings.Split(pair, ":")
		switch vals[0] {
		case "byr":
			pass.byr, err = createByr(vals[1])
		case "iyr":
			pass.iyr, err = createIyr(vals[1])
		case "eyr":
			pass.eyr, err = createEyr(vals[1])
		case "hgt":
			pass.hgt, err = createHgt(vals[1])
		case "hcl":
			pass.hcl, err = createHcl(vals[1])
		case "ecl":
			pass.ecl, err = createEcl(vals[1])
		case "pid":
			pass.pid, err = createPid(vals[1])
		case "cid":
			pass.cid = vals[1]
		default:
			panic("unknown key \"" + vals[0] + "\" in pair \"" + pair + "\"")
		}
		if err != nil {
			return pass, err
		}
	}

	hasMandatoryFields := pass.byr != "" &&
		pass.ecl != "" &&
		pass.eyr != "" &&
		pass.hcl != "" &&
		pass.hgt != "" &&
		pass.iyr != "" &&
		pass.pid != ""

	if !hasMandatoryFields {
		return pass, fmt.Errorf("Misses mandatory fields")
	}

	return pass, nil
}

func readFileTask1(filename string) []passport {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	passports := []passport{}
	scanner := bufio.NewScanner(file)
	inputForOnePassport := ""
	for scanner.Scan() {
		if scanner.Text() == "" {
			pass, err := createPassportTask1(inputForOnePassport)
			if err == nil {
				passports = append(passports, pass)
			}
			inputForOnePassport = ""
			continue
		}
		inputForOnePassport += " " + scanner.Text()
	}
	if inputForOnePassport != "" {
		pass, err := createPassportTask1(inputForOnePassport)
		if err == nil {
			passports = append(passports, pass)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%+v\n", passports)
	return passports
}

func readFileTask2(filename string) []passport {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	passports := []passport{}
	scanner := bufio.NewScanner(file)
	inputForOnePassport := ""
	for scanner.Scan() {
		if scanner.Text() == "" {
			pass, err := createPassportTask2(inputForOnePassport)
			if err == nil {
				passports = append(passports, pass)
			}
			inputForOnePassport = ""
			continue
		}
		inputForOnePassport += " " + scanner.Text()
	}
	if inputForOnePassport != "" {
		pass, err := createPassportTask2(inputForOnePassport)
		if err == nil {
			passports = append(passports, pass)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%+v\n", passports)
	return passports
}

func task1(passports []passport) int {
	return len(passports)
}

func task2(passports []passport) int {
	return len(passports)
}

// task1: 226
// task2: 160

func main() {
	fmt.Printf("Task1: %v\n", task1(readFileTask1("input.txt")))
	fmt.Printf("Task2: %v\n", task2(readFileTask2("input.txt")))
}
