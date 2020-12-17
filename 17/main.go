package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cubeSpace map[int]map[int]map[int]map[int]bool

func copy(c cubeSpace) cubeSpace {
	newC := cubeSpace{}
	for w := range c {
		newC[w] = make(map[int]map[int]map[int]bool)
		for z := range c[w] {
			newC[w][z] = map[int]map[int]bool{}
			for y := range c[w][z] {
				newC[w][z][y] = map[int]bool{}
				for x := range c[w][z][y] {
					v, ok := c[w][z][y][x]
					if !ok {
						newC[w][z][y][x] = false
					}
					newC[w][z][y][x] = v
				}
			}
		}
	}
	return newC
}

func (c cubeSpace) getMinW() int {
	wMin := 0
	for w := range c {
		if w < wMin {
			wMin = w
		}
	}
	return wMin
}

func (c cubeSpace) getMinZ() int {
	zMin := 0
	for z := range c[0] {
		if z < zMin {
			zMin = z
		}
	}
	return zMin
}

func (c cubeSpace) getMinY() int {
	yMin := 0
	for y := range c[0][0] {
		if y < yMin {
			yMin = y
		}
	}
	return yMin
}

func (c cubeSpace) getMinX() int {
	xMin := 0
	for x := range c[0][0][0] {
		if x < xMin {
			xMin = x
		}
	}
	return xMin
}

func (c cubeSpace) printAll() {
	for w := range c {
		for z := range c[w] {
			fmt.Printf("z=%v,w=%v\n", z, w)
			c.print(w, z)
		}
	}
}

func (c cubeSpace) print(w, z int) {
	fmt.Println()
	for y := range c[w][z] {
		for x := range c[w][z][y] {
			if c[w][z][y][x] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func (c cubeSpace) numberOfActive(w, z int) int {
	active := 0
	for y := range c[w][z] {
		for x := range c[w][z][y] {
			if c[w][z][y][x] {
				active++
			}
		}
	}
	return active
}

func (c cubeSpace) totalActive() int {
	active := 0
	for w := range c {
		for z := range c[w] {
			active += c.numberOfActive(w, z)
		}
	}
	return active
}

func parseToCube(data []string) cubeSpace {
	c := cubeSpace{}
	rowLen := 0
	c[0] = map[int]map[int]map[int]bool{}
	c[0][0] = map[int]map[int]bool{}
	boolRow := map[int]bool{}
	for y, line := range data {
		boolRow = map[int]bool{}
		line = strings.TrimSpace(line)
		row := strings.Split(line, "")
		if rowLen == 0 {
			rowLen = len(row)
		}
		if rowLen != len(row) {
			panic("wrong row size!")
		}
		for i := range row {
			switch row[i] {
			case ".":
				boolRow[i] = false
			case "#":
				boolRow[i] = true
			default:
				panic("wrong input")
			}
		}

		c[0][0][y] = boolRow
	}

	c.print(0, 0)
	return c
}

func readFile(filename string) cubeSpace {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := []string{}
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return parseToCube(data)
}

func getNumberActiveNeighbors(xA, yA, zA, wA int, cubes cubeSpace) int {
	numberActive := 0
	for w := wA - 1; w <= wA+1; w++ {
		for z := zA - 1; z <= zA+1; z++ {
			for y := yA - 1; y <= yA+1; y++ {
				for x := xA - 1; x <= xA+1; x++ {
					if w == wA && z == zA && y == yA && x == xA {
						continue
					}
					if cubes[w][z][y][x] {
						numberActive++
					}
				}
			}
		}
	}

	return numberActive
}

func doIteration(cubes cubeSpace, wShift int) cubeSpace {
	minW := cubes.getMinW()
	minZ := cubes.getMinZ()
	minY := cubes.getMinY()
	minX := cubes.getMinX()
	lengthW := len(cubes)
	lengthZ := len(cubes[0])
	lengthY := len(cubes[0][0])
	lengthX := len(cubes[0][0][0])
	newCubes := copy(cubes)

	//newCubes.printAll()
	for w := minW - wShift; w < lengthW+wShift; w++ {
		//for w := minW; w < lengthW; w++ {
		for z := minZ - 1; z < lengthZ+1; z++ {
			for y := minY - 1; y < lengthY+1; y++ {
				for x := minX - 1; x < lengthX+1; x++ {
					if _, ok := cubes[w]; !ok {
						cubes[w] = map[int]map[int]map[int]bool{}
						newCubes[w] = map[int]map[int]map[int]bool{}
					}
					if _, ok := cubes[w][z]; !ok {
						cubes[w][z] = map[int]map[int]bool{}
						newCubes[w][z] = map[int]map[int]bool{}
					}
					if _, ok := cubes[w][z][y]; !ok {
						cubes[w][z][y] = map[int]bool{}
						newCubes[w][z][y] = map[int]bool{}

					}
					if _, ok := cubes[w][z][y][x]; !ok {
						cubes[w][z][y][x] = false
						newCubes[w][z][y][x] = false
					}

					n := getNumberActiveNeighbors(x, y, z, w, cubes)
					////fmt.Printf("cell w:%v,z:%v,%v,%v, neighbors: %v - ", w, z, y, x, n)
					if newCubes[w][z][y][x] == true {
						if n == 2 || n == 3 {
							newCubes[w][z][y][x] = true
							////fmt.Printf("on\n")
						} else {
							newCubes[w][z][y][x] = false
							////fmt.Printf("off\n")
						}
					} else if n == 3 {
						newCubes[w][z][y][x] = true
						////fmt.Printf("on\n")
					} else {
						newCubes[w][z][y][x] = false
						////fmt.Printf("off\n")
					}
				}
			}
		}
	}

	return newCubes
}

func task1(cubes cubeSpace) int {
	for i := 0; i < 6; i++ {
		cubes = doIteration(cubes, 0)
	}

	return cubes.totalActive()
}

func task2(cubes cubeSpace) int {
	for i := 0; i < 6; i++ {
		cubes = doIteration(cubes, 1)
	}

	return cubes.totalActive()
}

// task1: 359
// task2: 2228

func main() {
	input := readFile("input.txt")
	fmt.Printf("Task1: %v\n", task1(input))
	input = readFile("input.txt")
	fmt.Printf("Task2: %v\n", task2(input))
}
