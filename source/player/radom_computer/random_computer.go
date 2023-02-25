package radom_computer

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dineshgowda24/tic-tac-toe/source/board"
	"github.com/dineshgowda24/tic-tac-toe/source/player"
)

var _ player.Player = (*RandomComputer)(nil)

// RandomComputer implements Player interface
// Generates a random value when playing does not really care about winning
type RandomComputer struct {
	move player.Move
}

// New returns a new RandomComputer
func New(move player.Move) *RandomComputer {
	rand.Seed(time.Now().UnixNano())
	return &RandomComputer{
		move: move,
	}
}

// Move returns the selected symbol
func (r *RandomComputer) Move() player.Move {
	return r.move
}

// Name returns the name
func (r *RandomComputer) Name() string {
	return "Dumb Computer"
}

// Play returns a random number with the board range
func (r *RandomComputer) Play(grid *board.Board) int {
	fmt.Printf("%s is making a move..\n", r.Name())
	time.Sleep(time.Second * 1) // Simulate time delay
	return rand.Intn((grid.Size()*grid.Size())-1) + 1
}

// Notify is NOOP
func (r *RandomComputer) Notify(data string) error {
	return nil
}

// Exit is NOOP
func (r *RandomComputer) Exit() error {
	return nil
}
