package Puzzle12

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func RunSolution(inputFileLocation string) {
	_, start, end := readMap(inputFileLocation)
	path, _, found := path(start, end)

	if found {

		/*for stepNo, step := range path {
			fmt.Printf("Step No: %d: %d,%d\n", stepNo, step.(*Location).x, step.(*Location).y)
		} */
		fmt.Printf("Part 1 - Total Steps: %d\n", len(path)-1)

	} else {
		fmt.Println("No path found")
	}
}

func readMap(inputFileLocation string) (world topoMap, start Pather, end Pather) {

	Map := topoMap{}

	mapData, err := os.Open(inputFileLocation)
	if err != nil {
		fmt.Println(("There was an error reading input: ") + err.Error())
		return Map, nil, nil
	}
	defer mapData.Close()

	mapReader := bufio.NewReader(mapData)

	y := 0
	for {
		line, _, err := mapReader.ReadLine()

		mapLine := string(line)

		for x, chr := range mapLine {
			switch chr {
			case 'S':
				Map.setLocation(&Location{topoType: 'a'}, x, y)
				start = Map.getLocation(x, y)

			case 'E':
				Map.setLocation(&Location{topoType: 'z'}, x, y)
				end = Map.getLocation(x, y)

			default:
				Map.setLocation(&Location{topoType: int(chr)}, x, y)
			}
		}

		if err == io.EOF {
			break
		}

		y++
	}

	return Map, start, end
}
