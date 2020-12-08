package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Task1_1(t *testing.T) {
	input := readFileTask1("input_test_task1.txt")
	expectedResponse := 2

	// act
	res := task1(input)

	assert.Equal(t, expectedResponse, res)
}

func Test_IsValid_Task2_1(t *testing.T) {
	validPassports := []string{
		`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
	byr:1937 iyr:2017 cid:147 hgt:183cm`,
		`hcl:#ae17e1 iyr:2013
	eyr:2024
	ecl:brn pid:760753108 byr:1931
	hgt:179cm`,

		`pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
		hcl:#623a2f`,
		`eyr:2029 ecl:blu cid:129 byr:1989
		iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm`,
		`hcl:#888785
		hgt:164cm byr:2001 iyr:2015 cid:88
		pid:545766238 ecl:hzl
		eyr:2022`,
		`iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`,
	}
	for _, validPass := range validPassports {
		_, err := createPassportTask2(validPass)

		assert.Nil(t, err, validPass)
	}
}

func Test_IsValid_Task2_2(t *testing.T) {
	invalidPassports := []string{
		`hcl:#cfa07d eyr:2025 pid:166559648
	iyr:2011 ecl:brn hgt:59in`,
		`iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
	hcl:#cfa07d byr:1929`,

		`eyr:1972 cid:100
		hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926`,
		`iyr:2019
		hcl:#602927 eyr:1967 hgt:170cm
		ecl:grn pid:012533040 byr:1946`,
		`hcl:dab227 iyr:2012
		ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277`,
		`hgt:59cm ecl:zzz
		eyr:2038 hcl:74454a iyr:2023
		pid:3556412378 byr:2007`,
	}
	for _, invalidPass := range invalidPassports {
		_, err := createPassportTask2(invalidPass)

		assert.NotNil(t, err, invalidPass)
	}
}

func Test_createByr(t *testing.T) {
	value, err := createByr("2002")
	assert.Nil(t, err)
	assert.Equal(t, "2002", value)

	value, err = createByr("2003")
	assert.NotNil(t, err)
}

func Test_createHgt(t *testing.T) {
	value, err := createHgt("60in")
	assert.Nil(t, err)
	assert.Equal(t, "60in", value)

	value, err = createHgt("190cm")
	assert.Nil(t, err)
	assert.Equal(t, "190cm", value)

	value, err = createHgt("190in")
	assert.NotNil(t, err)

	value, err = createHgt("190")
	assert.NotNil(t, err)
}

func Test_createHcl(t *testing.T) {
	value, err := createHcl("#123abc")
	assert.Nil(t, err)
	assert.Equal(t, "#123abc", value)

	value, err = createHcl("#123abz")
	assert.NotNil(t, err)

	value, err = createHcl("123abc")
	assert.NotNil(t, err)
}

func Test_createEcl(t *testing.T) {
	value, err := createEcl("brn")
	assert.Nil(t, err)
	assert.Equal(t, "brn", value)

	value, err = createEcl("wat")
	assert.NotNil(t, err)
}

func Test_createPid(t *testing.T) {
	value, err := createPid("000000001")
	assert.Nil(t, err)
	assert.Equal(t, "000000001", value)

	value, err = createPid("0123456789")
	assert.NotNil(t, err)
}
