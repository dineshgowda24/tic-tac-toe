package factory

import (
	"io"

	"github.com/dineshgowda24/tic-tac-toe/source/player"
	"github.com/dineshgowda24/tic-tac-toe/source/player/human"
	"github.com/dineshgowda24/tic-tac-toe/source/player/radom_computer"
	"github.com/dineshgowda24/tic-tac-toe/source/player/smart_computer"
)

// NewPlayer is a factory method that returns a new player
func NewPlayer(playerType player.Type, name string, mv player.Move, stream io.ReadWriteCloser) player.Player {
	switch playerType {
	case player.Human:
		return human.New(mv, name, stream)
	case player.SmartComputer:
		return smart_computer.New(mv)
	case player.RandomComputer:
		return radom_computer.New(mv)
	}
	return nil
}
