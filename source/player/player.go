package player

import (
	"strings"

	"github.com/dineshgowda24/tic-tac-toe/source/board"
)

// Player is an interface which represents someone who can play
// It will be an abstraction of a concrete type like human player, random computer, smart computer
type Player interface {
	Play(*board.Board) int
	Notify(data string) error
	Move() Move
	Name() string
	Exit() error
}

// Move represents the move made by a player
type Move int

var (
	// X is Player X move
	X Move = 3
	// O is Player O move
	O Move = -3
	// NP move which has not been played yet
	NP Move = 0
)

var moveToString = map[Move]string{
	X:  "X",
	O:  "O",
	NP: "-",
}

var stringToMove = map[string]Move{
	"x": X,
	"o": O,
	"-": NP,
}

// String returns the string representation of move
func (m Move) String() string {
	switch m {
	case X, O, NP:
		return moveToString[m]
	default:
		panic("this should never happen")
	}
}

// NewMove returns a new Move from string
func NewMove(move string) Move {
	switch strings.ToLower(move) {
	case "o", "x", "-":
		return stringToMove[strings.ToLower(move)]
	default:
		panic("this should never happen")
	}
}

// Type represents the type of a player
type Type int

var (
	Human          Type = 1
	SmartComputer  Type = 2
	RandomComputer Type = 3
)
