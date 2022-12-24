package Puzzle10

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// To syncronishe the goroutines
var wg sync.WaitGroup

type CPU struct {
	clock            chan int
	clockSpeed       int // setting to allow execution to be slowed so it can be traced visually
	cycleCount       int
	xRegister        int
	memory           map[int]int
	instructionQueue chan []string
	controlChan      chan string
}

func newCPU() CPU {
	cpu := CPU{}
	cpu.xRegister = 1
	cpu.clockSpeed = 0 // 0 is fast, 1 is slow. Use 1 to debug execution
	cpu.cycleCount = 0
	cpu.clock = make(chan int)
	cpu.controlChan = make(chan string) // not used right now
	cpu.instructionQueue = make(chan []string)
	cpu.memory = map[int]int{}

	// Load the signal checker inputs into the 'special' memory addresses
	// used by the regchecker function
	cpu.memory[-1] = 20
	cpu.memory[-2] = 60

	return cpu
}

func (c *CPU) start() {

	// execution pipeline
	wg.Add(1)
	go func() {

		for {

			// sync to clock
			c.cycleCount = <-c.clock

			// Check the X register after every clock cycle
			c.regChecker()

			instruction := <-c.instructionQueue
			//fmt.Printf("instruction: %s\n\n", instruction)
			if instruction[0] == "halt" {
				wg.Done()
				//c.controlChan <- "halt"
				return
			}

			// if instruction is addx, a clock cycle gets burnt here
			c.executeInstruction(instruction)

		}

	}()

	// cpu clock. Probably should stop the timer on halt, but easier to
	// let it terminate with main for now. Avoids a race
	go func() {
		timer := 1
		for {

			c.clock <- timer
			timer++

			// Execution speed based on clock speed setting
			time.Sleep(time.Duration(c.clockSpeed) * (time.Second / 2))

		}

	}()

}

func (c *CPU) executeInstruction(instruction []string) {

	switch instruction[0] {

	case "addx":
		c.cycleCount = <-c.clock // consume a cycle
		// Check the X register since we consumed a cycle
		c.regChecker()

		if len(instruction) < 1 {
			// Instruction does nothing if malformed
			fmt.Println("Bad instruction")
			return
		}

		addVal, err := strconv.Atoi(instruction[1])

		if err != nil {
			fmt.Println("Bad value passed to addx: " + err.Error())
			return
		}
		c.xRegister += addVal

	case "noop":
		//fmt.Println("noop")
		return

	case "halt":
		return

	}
}

// There might be a cooler way of doing this, but for now the CPU checks the X register on
// every cycle and can do something with it. In this case it stores the signal strength
// values it its memory

func (c *CPU) regChecker() {

	// Read the signal Cycle values from the 'special CPU memory' locations
	startSignalCycle := c.memory[-1] //20
	signalCycle := c.memory[-2]      //60

	//fmt.Printf("RegChecker Cycle: %d Value of X: %d\n", c.cycleCount, c.xRegister)

	if startSignalCycle != 0 && signalCycle != 0 {
		if c.cycleCount == startSignalCycle {
			c.memory[c.cycleCount] = c.xRegister
		}

		if c.cycleCount == signalCycle {
			c.memory[c.cycleCount] = c.xRegister
			c.memory[-2] = c.memory[-2] + 40
		}
	}

}

func RunSolution(inputFileName string) {

	programData, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println(("There was an error reading input: ") + err.Error())
		return
	}
	defer programData.Close()

	programReader := bufio.NewReader(programData)

	cpu := newCPU()
	cpu.start()

	for {

		instructionLine, _, err := programReader.ReadLine()

		if err == io.EOF {

			cpu.instructionQueue <- []string{"halt"}
			wg.Wait()
			break
		}

		instruction := strings.Split(string(instructionLine), " ")
		cpu.instructionQueue <- instruction

	}

	// Read out the CPU memory to calculate signal strength
	signalStrengh := 0
	for key, val := range cpu.memory {
		//fmt.Printf("key: %d, val: %d\n", key, val)

		// Ignore the 'special' memory addresses below zero
		if key > 0 {
			signalStrengh += key * val
		}
	}

	fmt.Printf("Part 1 - Signal stregth is: %d\n", signalStrengh)
}
