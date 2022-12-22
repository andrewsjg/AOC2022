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

	if len(ropeTail.locationVisits) != 2 {
		t.Errorf("Expected RopeTail locationvisits = 2, got: %d}", len(ropeTail.locationVisits))

	}
}
