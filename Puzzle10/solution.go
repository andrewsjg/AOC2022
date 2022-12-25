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

// Use a waitgroup to syncronishe the goroutines
var wg sync.WaitGroup

type CPU struct {
	clock            chan int
	clockSpeed       int // setting to allow execution to be slowed so it can be traced visually
	cycleCount       int
	xRegister        int
	cpuState         string
	memory           map[int]int
	instructionQueue chan []string
	busConnector     chan busMessage
}

type busMessage struct {
	clockSignal    int
	xRegisterValue int
	controlMsg     string
}

type scanLine [40]string

type CRT struct {
	displayBuffer [6]scanLine
	currentLine   int
	currentPixel  int

	busConnector chan busMessage
}

type Machine struct {
	cpu CPU
	crt CRT

	crtbus chan busMessage
	cpubus chan busMessage
}

// Configure a CPU
func newCPU(busConnector chan busMessage) CPU {
	cpu := CPU{}
	cpu.xRegister = 1
	cpu.clockSpeed = 0 // 0 is fast, 1 is slow. Use 1 to debug execution
	cpu.cycleCount = 0
	cpu.clock = make(chan int)
	cpu.instructionQueue = make(chan []string)
	cpu.memory = map[int]int{}
	cpu.busConnector = busConnector

	// Load the signal checker inputs into the 'special' memory addresses
	// used by the busprocessor function
	cpu.memory[-1] = 20
	cpu.memory[-2] = 60

	return cpu
}

// Configure a CRT
func newCRT(busconnector chan busMessage) CRT {
	crt := CRT{}

	crt.displayBuffer = [6]scanLine{}
	crt.currentLine = -1 // This is a little hack so the display logic is slightly easier
	crt.currentPixel = 0
	crt.busConnector = busconnector

	return crt
}

// Configure a  machine with a CPU and a CRT
func newMachine() Machine {
	machine := Machine{}

	machine.cpubus = make(chan busMessage)
	machine.cpu = newCPU(machine.cpubus)

	machine.crtbus = make(chan busMessage)
	machine.crt = newCRT(machine.crtbus)

	return machine
}

func (m *Machine) start() {

	// The bus processor brokers messages between
	// peripherals and the CPU
	//wg.Add(1)
	go m.busProcessor()

	// Start the CRT
	wg.Add(1)
	go m.crt.start()

	// Start the CPU
	m.cpu.start()

}

// Connect all the various peripherals together
// Only a CRT and a CPU. So all this does is relay
// messages from the CPU to the CRT
//
// NOTE: Channels arent actually a great analogy for a bus.
// but it works for now.

func (m *Machine) busProcessor() {

	//defer wg.Done()
	for {
		cpuBusMsg := <-m.cpubus
		m.crtbus <- cpuBusMsg
	}

}

func (crt *CRT) start() {

	defer wg.Done()

	for {
		// Get messages via the bus
		busMsg := <-crt.busConnector

		// Because the clock starts at 1 and the crt line starts at 0
		// need to do some messing around to sort out the current pixel
		pixel := (busMsg.clockSignal % 40) - 1
		if pixel == -1 {
			pixel = 39
		}

		crt.currentPixel = pixel

		if crt.currentPixel == 0 {
			crt.currentLine++

			if crt.currentLine == 6 {
				crt.currentLine = 0
			}
		}

		spritePixel1 := busMsg.xRegisterValue - 1
		spritePixel2 := busMsg.xRegisterValue
		spritePixel3 := busMsg.xRegisterValue + 1

		// At this point the current pixel and display line are determined and the location of the sprite recorded

		// Debug line to show the state of the CRT
		//fmt.Printf("CRT clock signal: %d xVal: %d Current Pixel: %d Current Line: %d Sprite Pixels: {%d,%d,%d}\n", busMsg.clockSignal, busMsg.xRegisterValue, crt.currentPixel, crt.currentLine, spritePixel1, spritePixel2, spritePixel3)

		// Write to the display buffer

		if crt.currentPixel == spritePixel1 || crt.currentPixel == spritePixel2 || crt.currentPixel == spritePixel3 {
			crt.displayBuffer[crt.currentLine][crt.currentPixel] = "#"
		} else {
			// On my screen it was easier to see the message without rendering a '.' for the none
			// 'unlit' pixels
			crt.displayBuffer[crt.currentLine][crt.currentPixel] = " "
		}

		// There should probably also be a 'flush' message to signal when the display
		// buffer can be safely read. But since we just need to render one 'screen' we will
		// never fill the buffer more than once. So simply halt and render afterwards.
		if busMsg.controlMsg == "HALT" {
			return
		}

	}
}

// Render the screen based on what is in the display buffer
func (crt *CRT) renderScreen() {
	for _, line := range crt.displayBuffer {
		for _, pixel := range line {
			fmt.Print(pixel)
		}
		fmt.Println()
	}
}

func (c *CPU) start() {

	c.cpuState = "RUN"
	// The execution pipeline
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {

			// sync to clock
			c.cycleCount = <-c.clock

			// Call the bus updater every clock cycle
			c.busUpdater()

			instruction := <-c.instructionQueue
			//fmt.Printf("instruction: %s\n\n", instruction)
			if instruction[0] == "halt" {
				c.cpuState = "HALT"
				// Call bus updater to propagate the halt
				c.busUpdater()

				return
			}

			// Remember, if the instruction is addx, a clock cycle gets burnt here
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
		// Because we used a clock cycle, need to call the bus updater
		c.busUpdater()

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

// There might be a cooler way of doing this, but for now the CPU checks the X register and
// updates the bus on every cycle and can do something with it.
// In this case it stores the signal strength values it its memory and propagates bus messages
// from the CPU

func (c *CPU) busUpdater() {

	// Read the signal Cycle values from the 'special CPU memory' locations
	startSignalCycle := c.memory[-1] //20
	signalCycle := c.memory[-2]      //60

	//fmt.Printf("BusUpdater Cycle: %d Value of X: %d\n", c.cycleCount, c.xRegister)

	if startSignalCycle != 0 && signalCycle != 0 {
		if c.cycleCount == startSignalCycle {
			c.memory[c.cycleCount] = c.xRegister
		}

		if c.cycleCount == signalCycle {
			c.memory[c.cycleCount] = c.xRegister
			c.memory[-2] = c.memory[-2] + 40
		}
	}

	if c.busConnector != nil {
		msg := busMessage{}
		msg.clockSignal = c.cycleCount
		msg.xRegisterValue = c.xRegister
		msg.controlMsg = c.cpuState

		c.busConnector <- msg
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

	machine := newMachine()
	machine.start()

	for {

		instructionLine, _, err := programReader.ReadLine()

		if err == io.EOF {

			machine.cpu.instructionQueue <- []string{"halt"}
			wg.Wait()
			break
		}

		instruction := strings.Split(string(instructionLine), " ")
		machine.cpu.instructionQueue <- instruction

	}

	// Read out the CPU memory to calculate signal strength
	signalStrengh := 0
	for key, val := range machine.cpu.memory {
		//fmt.Printf("key: %d, val: %d\n", key, val)

		// Ignore the 'special' memory addresses below zero
		if key > 0 {
			signalStrengh += key * val
		}
	}

	fmt.Printf("Part 1 - Signal stregth is: %d\n", signalStrengh)
	fmt.Println("Part 2 - The display contains: ")
	machine.crt.renderScreen()

}
