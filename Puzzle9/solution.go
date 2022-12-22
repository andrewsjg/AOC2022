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

	pathData, err := os.Open(input)
	defer pathData.Close()

	if err != nil {
		fmt.Println(("There was an error: ") + err.Error())
		return
	}

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

		// fmt.Printf("Direction: %s, Distance: %d\n", direction, distance)

		for i := 0; i < distance; i++ {
			moveRope(&ropeHead, &ropeTail, direction)
		}

	}

	fmt.Printf("Locations Tail Visited: %d, Locations Head Visited: %d\n", len(ropeTail.locationVisits), len(ropeHead.locationVisits)) // 6181
	//printLocations(ropeTail)

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
