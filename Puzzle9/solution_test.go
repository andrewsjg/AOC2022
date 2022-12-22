package Puzzle9

import (
	"fmt"
	"testing"
)

func TestBasicMoveLeft(t *testing.T) {

	ropeHead := makeRope()
	ropeTail := makeRope()

	moveRope(&ropeHead, &ropeTail, "L")

	expectedHeadLocation := location{-1, 0}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation := location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "L")

	fmt.Println(ropeHead.location)
	fmt.Println(ropeTail.location)

	expectedHeadLocation = location{-2, 0}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{-1, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

}

func TestBasicMoveRight(t *testing.T) {

	ropeHead := makeRope()
	ropeTail := makeRope()

	moveRope(&ropeHead, &ropeTail, "R")

	expectedHeadLocation := location{1, 0}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation := location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "R")

	fmt.Println(ropeHead.location)
	fmt.Println(ropeTail.location)

	expectedHeadLocation = location{2, 0}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{1, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

}

func TestBasicMoveUp(t *testing.T) {
	ropeHead := makeRope()
	ropeTail := makeRope()

	moveRope(&ropeHead, &ropeTail, "U")

	expectedHeadLocation := location{0, 1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation := location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "U")

	expectedHeadLocation = location{0, 2}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{0, 1}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

}

func TestBasicMoveDown(t *testing.T) {
	ropeHead := makeRope()
	ropeTail := makeRope()

	moveRope(&ropeHead, &ropeTail, "D")

	expectedHeadLocation := location{0, -1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation := location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "D")

	expectedHeadLocation = location{0, -2}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{0, -1}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

}

func TestMoveLeftUp(t *testing.T) {
	ropeHead := makeRope()
	ropeTail := makeRope()

	moveRope(&ropeHead, &ropeTail, "L")

	expectedHeadLocation := location{-1, 0}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation := location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "U")

	expectedHeadLocation = location{-1, 1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "U")

	expectedHeadLocation = location{-1, 2}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{-1, 1}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

}

func TestMoveLeftDown(t *testing.T) {
	ropeHead := makeRope()
	ropeTail := makeRope()

	moveRope(&ropeHead, &ropeTail, "L")

	expectedHeadLocation := location{-1, 0}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation := location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "D")

	expectedHeadLocation = location{-1, -1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "D")

	expectedHeadLocation = location{-1, -2}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{-1, -1}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

}

func TestMoveRightUp(t *testing.T) {
	ropeHead := makeRope()
	ropeTail := makeRope()

	moveRope(&ropeHead, &ropeTail, "R")

	expectedHeadLocation := location{1, 0}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation := location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "U")

	expectedHeadLocation = location{1, 1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "U")

	expectedHeadLocation = location{1, 2}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{1, 1}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

}

func TestMoveRightDown(t *testing.T) {
	ropeHead := makeRope()
	ropeTail := makeRope()

	moveRope(&ropeHead, &ropeTail, "R")

	expectedHeadLocation := location{1, 0}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation := location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "D")

	expectedHeadLocation = location{1, -1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "D")

	expectedHeadLocation = location{1, -2}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{1, -1}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

}

func TestMoveUpLeft(t *testing.T) {
	ropeHead := makeRope()
	ropeTail := makeRope()

	moveRope(&ropeHead, &ropeTail, "U")

	expectedHeadLocation := location{0, 1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation := location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "L")

	expectedHeadLocation = location{-1, 1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "L")

	expectedHeadLocation = location{-2, 1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{-1, 1}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

}

func TestMoveDownLeft(t *testing.T) {
	ropeHead := makeRope()
	ropeTail := makeRope()

	moveRope(&ropeHead, &ropeTail, "D")

	expectedHeadLocation := location{0, -1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation := location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "L")

	expectedHeadLocation = location{-1, -1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "L")

	expectedHeadLocation = location{-2, -1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{-1, -1}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

}

func TestMoveDownRight(t *testing.T) {
	ropeHead := makeRope()
	ropeTail := makeRope()

	moveRope(&ropeHead, &ropeTail, "D")

	expectedHeadLocation := location{0, -1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation := location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "R")

	expectedHeadLocation = location{1, -1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "R")

	expectedHeadLocation = location{2, -1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{1, -1}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

}

func TestMoveUpRight(t *testing.T) {
	ropeHead := makeRope()
	ropeTail := makeRope()

	moveRope(&ropeHead, &ropeTail, "U")

	expectedHeadLocation := location{0, 1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation := location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "R")

	expectedHeadLocation = location{1, 1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveRope(&ropeHead, &ropeTail, "R")

	expectedHeadLocation = location{2, 1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{1, 1}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	expectedTailVisits := 2

	if expectedTailVisits != len(ropeTail.locationVisits) {
		t.Errorf("Expected tail to have visted 2 locations. Got: %d", len(ropeTail.locationVisits))

	}

}

func TestPath(t *testing.T) {
	ropeHead := makeRope()
	ropeTail := makeRope()

	moveHead(&ropeHead, "U")
	moveTail(ropeHead, &ropeTail)

	expectedHeadLocation := location{0, 1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation := location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveHead(&ropeHead, "R")
	moveTail(ropeHead, &ropeTail)

	expectedHeadLocation = location{1, 1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{0, 0}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	moveHead(&ropeHead, "R")
	moveTail(ropeHead, &ropeTail)

	expectedHeadLocation = location{2, 1}
	if ropeHead.location != expectedHeadLocation {
		t.Errorf("Expected Location for RopeHead was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedHeadLocation.x, expectedHeadLocation.y, ropeHead.location.x, ropeHead.location.y)

	}

	expectedTailLocation = location{1, 1}
	if ropeTail.location != expectedTailLocation {
		t.Errorf("Expected Location for RopeTail was wrong. Expected: {%d,%d} Got: {%d,%d}", expectedTailLocation.x, expectedTailLocation.y, ropeTail.location.x, ropeTail.location.y)

	}

	if len(ropeTail.locationVisits) != 2 {
		t.Errorf("Expected RopeTail locationvisits = 2, got: %d}", len(ropeTail.locationVisits))

	}
}

func TestLongRope(t *testing.T) {
	head := makeRope()
	knot1 := makeRope()
	knot2 := makeRope()
	knot3 := makeRope()
	knot4 := makeRope()
	knot5 := makeRope()
	knot6 := makeRope()
	knot7 := makeRope()
	knot8 := makeRope()
	tail := makeRope()

	for i := 1; i <= 10; i++ {
		moveHead(&head, "U")
		moveTail(head, &knot1)

		moveTail(knot1, &knot2)
		moveTail(knot2, &knot3)
		moveTail(knot3, &knot4)
		moveTail(knot4, &knot5)
		moveTail(knot5, &knot6)
		moveTail(knot6, &knot7)
		moveTail(knot7, &knot8)

		moveTail(knot8, &tail)
	}

	if tail.location.y != 1 {
		t.Errorf("Expected Location tail Y location to be 1 got %d", tail.location.y)

	}

	if len(tail.locationVisits) != 2 {
		t.Errorf("Expected Locations to be 2 got %d", len(tail.locationVisits))

	}

	fmt.Println(tail.location)

	for i := 1; i <= 10; i++ {
		moveHead(&head, "R")
		moveTail(head, &knot1)

		moveTail(knot1, &knot2)
		moveTail(knot2, &knot3)
		moveTail(knot3, &knot4)
		moveTail(knot4, &knot5)
		moveTail(knot5, &knot6)
		moveTail(knot6, &knot7)
		moveTail(knot7, &knot8)

		moveTail(knot8, &tail)
		fmt.Println(tail.location)
	}

	if tail.location.x != 1 {
		t.Errorf("Expected Location tail X location to be 1 got %d", tail.location.x)

	}

	if len(tail.locationVisits) != 3 {
		t.Errorf("Expected Locations to be 3 got %d", len(tail.locationVisits))

	}
}
