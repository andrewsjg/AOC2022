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

type knot struct {
	location       location
	locationVisits locationVisits
}
type rope struct {
	knots []knot
}

type locationVisits map[location]int

// Makes a rope
func makeRope(knots int) rope {

	newRope := rope{}

	for i := 0; i < knots; i++ {
		knot := knot{}
		knot.location = location{0, 0}
		knot.locationVisits = map[location]int{}
		knot.locationVisits[knot.location] += 1
		newRope.knots = append(newRope.knots, knot)
	}

	return newRope
}

func RunSolution(input string) {

	p1Rope := makeRope(2)
	p2Rope := makeRope(10)

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

		for i := 0; i < distance; i++ {

			p1Rope.moveHead(direction)
			p2Rope.moveHead(direction)

		}

	}

	p1tail := p1Rope.knots[len(p1Rope.knots)-1]
	p1head := p1Rope.knots[0]

	p2tail := p2Rope.knots[len(p2Rope.knots)-1]
	p2head := p2Rope.knots[0]

	fmt.Printf("Part 1 - Locations Tail Visited: %d, Locations Head Visited: %d\n", len(p1tail.locationVisits), len(p1head.locationVisits))
	fmt.Printf("Part 2 - Locations Tail Visited: %d, Locations Head Visited: %d\n", len(p2tail.locationVisits), len(p2head.locationVisits))

}

func (r *rope) moveHead(direction string) {
	switch direction {
	case "U":
		r.knots[0].location.y += 1

	case "D":
		r.knots[0].location.y -= 1

	case "L":

		r.knots[0].location.x -= 1

	case "R":
		r.knots[0].location.x += 1

	}

	r.knots[0].locationVisits[r.knots[0].location] += 1

	for i := 1; i < len(r.knots); i++ {

		diffX := r.knots[i-1].location.x - r.knots[i].location.x
		diffY := r.knots[i-1].location.y - r.knots[i].location.y

		// Only move if the tail is within 1 of the head
		if !(diffX < 2 && diffX > -2 && diffY < 2 && diffY > -2) {
			if diffY > 0 {

				r.knots[i].location.y += (diffY + 1) / 2

			} else {

				r.knots[i].location.y += (diffY - 1) / 2

			}

			if diffX > 0 {

				r.knots[i].location.x += (diffX + 1) / 2

			} else {

				r.knots[i].location.x += (diffX - 1) / 2

			}
		}
		r.knots[i].locationVisits[r.knots[i].location] += 1
	}

}
