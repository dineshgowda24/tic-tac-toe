package smart_computer

import (
	"github.com/dineshgowda24/tic-tac-toe/source/board"
	"github.com/dineshgowda24/tic-tac-toe/source/game"
	"github.com/dineshgowda24/tic-tac-toe/source/player"
	"github.com/dineshgowda24/tic-tac-toe/source/player/radom_computer"
)

type SmartComputer struct {
	move player.Move
}

// NewSmartComputer returns a new SmartComputer
func NewSmartComputer(move player.Move) *SmartComputer {
	return &SmartComputer{
		move: move,
	}
}

// Move returns the selected symbol
func (r *SmartComputer) Move() player.Move {
	return r.move
}

// Name returns the name of the string
func (r *SmartComputer) Name() string {
	return "Smart Computer"
}

// Play returns a random number with the board range
func (r *SmartComputer) Play(grid *board.Board) int {
	return r.bestMove(grid)
}

// computes the optimal move for computer to win
func (r *SmartComputer) minimax(brd *board.Board, depth int, isMe bool) int {
	gm := game.NewGame(brd, r, radom_computer.NewRandomComputer(player.O))
	switch gm.Status() {
	case game.XPlayerWon:
		return int(player.X)
	case game.OPlayerWon:
		return int(player.O)
	case game.Draw:
		return 0
	}

	grid := brd.GridRef()
	if isMe {
		best := -1000
		for i := 1; i <= brd.Size()*brd.Size(); i++ {
			if grid[i] == 0 {
				brd.Move(i, int(r.Move()))
				best = max(best, r.minimax(brd, depth+1, false))
				brd.Reset(i)
			}
		}
		return best
	} else {
		best := 1000
		for i := 1; i <= brd.Size()*brd.Size(); i++ {
			if grid[i] == 0 {
				brd.Move(i, int(player.O))
				best = min(best, r.minimax(brd, depth+1, true))
				brd.Reset(i)
			}
		}
		return best
	}
}

// bestMove returns the best move for smart_computer to win
// It starts with computer making the move 1st
func (r *SmartComputer) bestMove(b *board.Board) int {
	var bestVal, move int
	bestVal = -1000
	brd := b.DeepCopy()
	grid := brd.GridRef()
	for i := 1; i <= b.Size()*b.Size(); i++ {
		if grid[i] == 0 {
			brd.Move(i, int(r.Move()))
			value := r.minimax(brd, 0, false)
			brd.Reset(i)
			if value > bestVal {
				bestVal = value
				move = i
			}
		}
	}
	return move
}

// min returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// min returns the smaller of x or y.
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// Notify is NOOP
func (r *SmartComputer) Notify(data string) error {
	return nil
}

// Exit is NOOP
func (r *SmartComputer) Exit() error {
	return nil
}
