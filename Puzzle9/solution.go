package Puzzle9

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type location struct {
	x int
	y int
}

type rope struct {
	location       location
	locationVisits locationVisits
	path           []location
}

type locationVisits map[location]int

// Makes a rope
func makeRope() rope {

	newRope := rope{}
	newRope.location = location{0, 0}
	newRope.locationVisits = locationVisits{}

	newRope.locationVisits[newRope.location] += 1
	newRope.path = []location{}

	return newRope
}

func RunSolution(input string) {

	ropeHead := makeRope()
	ropeTail := makeRope()

	p2Head := makeRope()
	knot1 := makeRope()

	knot2 := makeRope()
	knot3 := makeRope()

	knot4 := makeRope()
	knot5 := makeRope()

	knot6 := makeRope()
	knot7 := makeRope()

	knot8 := makeRope()
	p2Tail := makeRope()

	pathData, err := os.Open(input)
	if err != nil {
		fmt.Println(("There was an error reading input: ") + err.Error())
		return
	}
	defer pathData.Close()

	pathReader := bufio.NewReader(pathData)

	for {

		pathElement, _, err := pathReader.ReadLine()

		if err == io.EOF {
			break
		}

		direction := string(pathElement[0:1])
		distance, err := strconv.Atoi(string(pathElement[2:]))

		if err != nil {
			fmt.Println("Error converting distance value: " + err.Error())
		}

		// Part 1
		for i := 0; i < distance; i++ {
			//moveRope(&ropeHead, &ropeTail, direction)
			moveHead(&ropeHead, direction)
			moveTail(ropeHead, &ropeTail)
		}

		// Part 2
		for i := 0; i < distance; i++ {
			moveHead(&p2Head, direction)
			moveTail(p2Head, &knot1)

			moveTail(knot1, &knot2)
			moveTail(knot2, &knot3)
			moveTail(knot3, &knot4)
			moveTail(knot4, &knot5)
			moveTail(knot5, &knot6)
			moveTail(knot6, &knot7)
			moveTail(knot7, &knot8)

			moveTail(knot8, &p2Tail)

			/* fmt.Println(p2Head.location)
			fmt.Println(p2Tail.location)
			fmt.Println() */

		}

	}

	fmt.Printf("Part 1 - Locations Tail Visited: %d, Locations Head Visited: %d\n", len(ropeTail.locationVisits), len(ropeHead.locationVisits))
	fmt.Printf("Part 2 - Locations Tail Visited: %d, Locations Head Visited: %d\n", len(p2Tail.locationVisits), len(p2Head.locationVisits)) // 2148 WRONG! 2386

}

func moveHead(ropeHead *rope, direction string) {

	switch direction {
	case "U":
		ropeHead.location.y += 1

	case "D":
		ropeHead.location.y -= 1

	case "L":

		ropeHead.location.x -= 1

	case "R":
		ropeHead.location.x += 1

	}

	ropeHead.path = append(ropeHead.path, ropeHead.location)
	ropeHead.locationVisits[ropeHead.location] += 1
}

func moveTail(ropeHead rope, ropeTail *rope) {

	if (ropeTail.location.y + 2) == ropeHead.location.y {

		ropeTail.location.y = ropeTail.location.y + 1
		ropeTail.location.x = ropeHead.location.x
		ropeTail.path = append(ropeTail.path, ropeTail.location)

		ropeTail.locationVisits[ropeTail.location] += 1

	}

	if (ropeTail.location.y - 2) == ropeHead.location.y {

		ropeTail.location.y = ropeTail.location.y - 1
		ropeTail.location.x = ropeHead.location.x
		ropeTail.path = append(ropeTail.path, ropeTail.location)

		ropeTail.locationVisits[ropeTail.location] += 1

	}

	if (ropeTail.location.x - 2) == ropeHead.location.x {

		ropeTail.location.x = ropeTail.location.x - 1
		ropeTail.location.y = ropeHead.location.y
		ropeTail.path = append(ropeTail.path, ropeTail.location)

		ropeTail.locationVisits[ropeTail.location] += 1

	}

	if (ropeTail.location.x + 2) == ropeHead.location.x {

		ropeTail.location.x = ropeTail.location.x + 1
		ropeTail.location.y = ropeHead.location.y
		ropeTail.path = append(ropeTail.path, ropeTail.location)

		ropeTail.locationVisits[ropeTail.location] += 1

	}

}

func moveRope(ropeHead *rope, ropeTail *rope, direction string) {
	switch direction {
	case "U":
		ropeHead.location.y += 1

		if (ropeTail.location.y + 2) == ropeHead.location.y {

			ropeTail.location.y = ropeTail.location.y + 1
			ropeTail.location.x = ropeHead.location.x
			ropeTail.path = append(ropeTail.path, ropeTail.location)

			ropeTail.locationVisits[ropeTail.location] += 1

		}

	case "D":
		ropeHead.location.y -= 1

		if (ropeTail.location.y - 2) == ropeHead.location.y {

			ropeTail.location.y = ropeTail.location.y - 1
			ropeTail.location.x = ropeHead.location.x
			ropeTail.path = append(ropeTail.path, ropeTail.location)

			ropeTail.locationVisits[ropeTail.location] += 1

		}

	case "L":

		ropeHead.location.x -= 1

		if (ropeTail.location.x - 2) == ropeHead.location.x {

			ropeTail.location.x = ropeTail.location.x - 1
			ropeTail.location.y = ropeHead.location.y
			ropeTail.path = append(ropeTail.path, ropeTail.location)

			ropeTail.locationVisits[ropeTail.location] += 1

		}

	case "R":
		ropeHead.location.x += 1

		if (ropeTail.location.x + 2) == ropeHead.location.x {

			ropeTail.location.x = ropeTail.location.x + 1
			ropeTail.location.y = ropeHead.location.y
			ropeTail.path = append(ropeTail.path, ropeTail.location)

			ropeTail.locationVisits[ropeTail.location] += 1

		}

	}

	ropeHead.path = append(ropeHead.path, ropeHead.location)
	ropeHead.locationVisits[ropeHead.location] += 1

}
