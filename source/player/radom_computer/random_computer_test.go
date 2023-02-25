package radom_computer

import (
	"testing"

	"github.com/dineshgowda24/tic-tac-toe/source/board"
	player2 "github.com/dineshgowda24/tic-tac-toe/source/player"
)

func TestNew(t *testing.T) {
	var player player2.Player
	player = New(player2.X)
	if player == nil {
		t.Error("human can not be nil")
	}
}

func TestRandomComputerPlay(t *testing.T) {
	tests := []int{3, 4, 5, 6}
	for _, size := range tests {
		var player player2.Player
		player = New(player2.X)
		if player == nil {
			t.Error("human can not be nil")
			t.FailNow()
		}

		brd, _ := board.NewBoard(3)
		moveOne := player.Play(brd)
		if moveOne < 0 || moveOne > size*size {
			t.Errorf("move can not be %d", moveOne)
			t.FailNow()
		}

		moveTwo := player.Play(brd)
		if moveTwo < 0 || moveTwo > size*size {
			t.Errorf("move can not be %d", moveTwo)
			t.FailNow()
		}

		moveThree := player.Play(brd)
		if moveThree < 0 || moveThree > size*size {
			t.Errorf("move can not be %d", moveThree)
			t.FailNow()
		}

		if moveOne == moveTwo && moveTwo == moveThree {
			t.Errorf("moves should have been random but got %d %d %d", moveOne, moveTwo, moveThree)
			t.FailNow()
		}
	}
}
