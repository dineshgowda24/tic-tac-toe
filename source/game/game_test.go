package game

import (
	"strings"
	"testing"

	"github.com/dineshgowda24/tic-tac-toe/source/board"
	"github.com/dineshgowda24/tic-tac-toe/source/player"
	"github.com/dineshgowda24/tic-tac-toe/source/player/radom_computer"
)

func TestPrintGameBoard3X3(t *testing.T) {
	brd, _ := board.NewBoard(3)
	gm := NewGame(brd, nil, nil)
	str := `-|-|-
-|-|-
-|-|-`
	if strings.TrimSpace(gm.String()) != strings.TrimSpace(str) {
		t.Errorf("expected %s  but got %s", str, gm.String())
		t.FailNow()
	}

	err := brd.Move(9, -3)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
		t.FailNow()
	}
	str = `-|-|-
-|-|-
-|-|O`
	if strings.TrimSpace(gm.String()) != strings.TrimSpace(str) {
		t.Errorf("expected %s  but got %s", str, gm.String())
		t.FailNow()
	}
}

func TestPrintGameBoard4X4(t *testing.T) {
	brd, _ := board.NewBoard(4)
	gm := NewGame(brd, nil, nil)
	str := `-|-|-|-
-|-|-|-
-|-|-|-
-|-|-|-`
	if strings.TrimSpace(gm.String()) != strings.TrimSpace(str) {
		t.Errorf("expected %s  but got %s", str, gm.String())
		t.FailNow()
	}

	err := brd.Move(16, -3)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
		t.FailNow()
	}

	err = brd.Move(1, 3)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
		t.FailNow()
	}

	str = `X|-|-|-
-|-|-|-
-|-|-|-
-|-|-|O`
	if strings.TrimSpace(gm.String()) != strings.TrimSpace(str) {
		t.Errorf("expected %s  but got %s", str, gm.String())
		t.FailNow()
	}
}

func TestVerticalResult3x3(t *testing.T) {
	const boardSize = 3
	brd, _ := board.NewBoard(boardSize)
	gm := NewGame(brd, radom_computer.NewRandomComputer(player.X), radom_computer.NewRandomComputer(player.O))
	if gm.checkVertical(brd.Grid()) != InProgress {
		t.Errorf("expected %v but got %v", InProgress, gm.Status())
		t.FailNow()
	}
	tests := []int{1, 4}
	for _, tt := range tests {
		err := brd.Move(tt, -3)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
			t.FailNow()
		}
		if gm.checkVertical(brd.Grid()) != InProgress {
			t.Errorf("expected %v but got %v", InProgress, gm.Status())
			t.FailNow()

		}
	}
	err := brd.Move(7, -3)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
		t.FailNow()
	}
	if gm.checkVertical(brd.Grid()) != OPlayerWon {
		t.Errorf("expected %v but got %v", OPlayerWon, gm.Status())
		t.FailNow()
	}
}

func TestVerticalResult4x4(t *testing.T) {
	const boardSize = 4
	brd, _ := board.NewBoard(boardSize)
	gm := NewGame(brd, radom_computer.NewRandomComputer(player.X), radom_computer.NewRandomComputer(player.O))
	if gm.checkVertical(brd.Grid()) != InProgress {
		t.Errorf("expected %v but got %v", InProgress, gm.Status())
		t.FailNow()
	}

	tests := []int{1, 5, 9}
	for _, tt := range tests {
		err := brd.Move(tt, 3)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
			t.FailNow()
		}
		if gm.checkVertical(brd.Grid()) != InProgress {
			t.Errorf("expected %v but got %v", InProgress, gm.Status())
			t.FailNow()

		}
	}
	err := brd.Move(13, 3)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
		t.FailNow()
	}
	if gm.checkVertical(brd.Grid()) != XPlayerWon {
		t.Errorf("expected %v but got %v", OPlayerWon, gm.Status())
		t.FailNow()
	}
}

func TestHorizontalResult3x3(t *testing.T) {
	const boardSize = 3
	brd, _ := board.NewBoard(boardSize)
	gm := NewGame(brd, radom_computer.NewRandomComputer(player.X), radom_computer.NewRandomComputer(player.O))
	if gm.checkHorizontal(brd.Grid()) != InProgress {
		t.Errorf("expected %v but got %v", InProgress, gm.Status())
		t.FailNow()
	}
	tests := []int{4, 5}
	for _, tt := range tests {
		err := brd.Move(tt, -3)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
			t.FailNow()
		}
		if gm.checkHorizontal(brd.Grid()) != InProgress {
			t.Errorf("expected %v but got %v", InProgress, gm.Status())
			t.FailNow()

		}
	}
	err := brd.Move(6, -3)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
		t.FailNow()
	}
	if gm.checkHorizontal(brd.Grid()) != OPlayerWon {
		t.Errorf("expected %v but got %v", OPlayerWon, gm.Status())
		t.FailNow()
	}
}

func TestHorizontalResult4x4(t *testing.T) {
	const boardSize = 4
	brd, _ := board.NewBoard(boardSize)
	gm := NewGame(brd, radom_computer.NewRandomComputer(player.X), radom_computer.NewRandomComputer(player.O))
	if gm.checkHorizontal(brd.Grid()) != InProgress {
		t.Errorf("expected %v but got %v", InProgress, gm.Status())
		t.FailNow()
	}

	tests := []int{13, 14, 15}
	for _, tt := range tests {
		err := brd.Move(tt, 3)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
			t.FailNow()
		}
		if gm.checkHorizontal(brd.Grid()) != InProgress {
			t.Errorf("expected %v but got %v", InProgress, gm.Status())
			t.FailNow()

		}
	}
	err := brd.Move(16, 3)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
		t.FailNow()
	}
	if gm.checkHorizontal(brd.Grid()) != XPlayerWon {
		t.Errorf("expected %v but got %v", OPlayerWon, gm.Status())
		t.FailNow()
	}
}

func TestDiagonalResult3x3(t *testing.T) {
	const boardSize = 3
	brd, _ := board.NewBoard(boardSize)
	gm := NewGame(brd, radom_computer.NewRandomComputer(player.X), radom_computer.NewRandomComputer(player.O))
	if gm.checkHorizontal(brd.Grid()) != InProgress {
		t.Errorf("expected %v but got %v", InProgress, gm.Status())
		t.FailNow()
	}
	tests := []int{1, 5}
	for _, tt := range tests {
		err := brd.Move(tt, -3)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
			t.FailNow()
		}
		if gm.checkDiagonal(brd.Grid()) != InProgress {
			t.Errorf("expected %v but got %v", InProgress, gm.Status())
			t.FailNow()
		}
	}
	err := brd.Move(9, -3)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
		t.FailNow()
	}
	if gm.checkDiagonal(brd.Grid()) != OPlayerWon {
		t.Errorf("expected %v but got %v", OPlayerWon, gm.Status())
		t.FailNow()
	}
}

func TestDiagonalResult4x4(t *testing.T) {
	const boardSize = 4
	brd, _ := board.NewBoard(boardSize)
	gm := NewGame(brd, radom_computer.NewRandomComputer(player.X), radom_computer.NewRandomComputer(player.O))
	if gm.checkHorizontal(brd.Grid()) != InProgress {
		t.Errorf("expected %v but got %v", InProgress, gm.Status())
		t.FailNow()
	}

	tests := []int{1, 6, 11}
	for _, tt := range tests {
		err := brd.Move(tt, 3)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
			t.FailNow()
		}
		if gm.checkDiagonal(brd.Grid()) != InProgress {
			t.Errorf("expected %v but got %v", InProgress, gm.Status())
			t.FailNow()

		}
	}
	err := brd.Move(16, 3)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
		t.FailNow()
	}
	if gm.checkDiagonal(brd.Grid()) != XPlayerWon {
		t.Errorf("expected %v but got %v", OPlayerWon, gm.Status())
		t.FailNow()
	}
}
