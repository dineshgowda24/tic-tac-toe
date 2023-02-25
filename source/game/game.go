package game

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/dineshgowda24/tic-tac-toe/source/board"
	"github.com/dineshgowda24/tic-tac-toe/source/player"
)

// Game represents the current context of a tic tac toe game
// Its has board information and all the relevant players in the game
type Game struct {
	board     *board.Board
	playerOne player.Player
	playerTwo player.Player
}

// NewGame returns a new Game
func NewGame(b *board.Board, playerOne player.Player, playerTwo player.Player) *Game {
	return &Game{
		board:     b,
		playerOne: playerOne,
		playerTwo: playerTwo,
	}
}

// String returns the string representation of the board
func (g *Game) String() string {
	size := g.board.Size()
	grid := g.board.Grid()
	var buf bytes.Buffer
	var rows []string
	offset := 1
	for i := 1; i <= size*size; i++ {
		rows = append(rows, fmt.Sprintf("%s", player.Move(grid[i]).String()))
		offset++
		if offset > size {
			fmt.Fprintf(&buf, fmt.Sprintf("%s\n", strings.Join(rows, "|")))
			offset = 1
			rows = []string{}
		}
	}
	return buf.String()
}

type Status string

var (
	Draw       Status = "Draw"
	InProgress Status = "In Progress"
	XPlayerWon Status = "Player X Won"
	OPlayerWon Status = "Player O Won"
)

func hasPlayerWon(status Status) bool {
	if status == XPlayerWon || status == OPlayerWon {
		return true
	}
	return false
}

// Status returns the  current status of the game
// It validates the board horizontally, vertically and diagonally
// If there is a winner, it returns it else returns the state of the board
func (g *Game) Status() Status {
	grid := g.board.Grid()
	if status := g.checkHorizontal(grid); hasPlayerWon(status) {
		return status
	}
	if status := g.checkVertical(grid); hasPlayerWon(status) {
		return status
	}
	if status := g.checkDiagonal(grid); hasPlayerWon(status) {
		return status
	}
	if g.board.HasEmptyCells() {
		return InProgress
	}
	return Draw
}

// checkHorizontal checks for a winner in all the rows, if a winner is identified then it returns
// It adds all the cells horizontally and checks if the sum is equal size * move
func (g *Game) checkHorizontal(grid []int) Status {
	size := g.board.Size()
	for cell := 1; cell <= (size * size); {
		offset := 1
		sum := 0
		for {
			sum += grid[cell]
			cell++
			offset++
			if offset > size {
				break
			}
		}
		if sum == total(g.playerTwo.Move(), size) {
			if g.playerTwo.Move() == player.X {
				return XPlayerWon
			} else if g.playerTwo.Move() == player.O {
				return OPlayerWon
			}
		}

		if sum == total(g.playerOne.Move(), size) {
			if g.playerOne.Move() == player.X {
				return XPlayerWon
			} else if g.playerOne.Move() == player.O {
				return OPlayerWon
			}
		}
	}
	if g.board.HasEmptyCells() {
		return InProgress
	}
	return Draw
}

func total(move player.Move, size int) int {
	val := int(move)
	res := 0
	for i := size; i >= 1; i-- {
		res += val
	}
	return res
}

// checkVertical checks for a winner in all the columns, if a winner is identified then it returns
// It adds all the cells vertically and checks if the sum is equal size * move
func (g *Game) checkVertical(grid []int) Status {
	size := g.board.Size()
	for cell := 1; cell <= size; cell++ {
		offset := 1
		sum := 0
		col := cell
		for {
			sum += grid[col]
			offset++
			col += size
			if offset > size {
				break
			}
		}
		if sum == total(g.playerTwo.Move(), size) {
			if g.playerTwo.Move() == player.X {
				return XPlayerWon
			} else if g.playerTwo.Move() == player.O {
				return OPlayerWon
			}
		}

		if sum == total(g.playerOne.Move(), size) {
			if g.playerOne.Move() == player.X {
				return XPlayerWon
			} else if g.playerOne.Move() == player.O {
				return OPlayerWon
			}
		}
	}
	if g.board.HasEmptyCells() {
		return InProgress
	}
	return Draw
}

// checkDiagonal checks for a winner, if a winner is identified then it returns
// It adds all the cells diagonally and checks if the sum is equal size * move
func (g *Game) checkDiagonal(grid []int) Status {
	size := g.board.Size()
	offset := 0
	cell := 1
	sum := 0
	// start from the 1st cell and diagonally move down
	for {
		sum += grid[cell]
		cell += size + 1
		offset++
		if offset == size {
			break
		}
	}
	if sum == total(g.playerTwo.Move(), size) {
		if g.playerTwo.Move() == player.X {
			return XPlayerWon
		} else if g.playerTwo.Move() == player.O {
			return OPlayerWon
		}
	}

	if sum == total(g.playerOne.Move(), size) {
		if g.playerOne.Move() == player.X {
			return XPlayerWon
		} else if g.playerOne.Move() == player.O {
			return OPlayerWon
		}
	}

	offset = 0
	cell = size * size
	sum = 0
	// start from the last cell from left and diagonally move up
	for {
		cell = cell - size + 1
		sum += grid[cell]
		offset++
		if offset == size {
			break
		}
	}
	if sum == total(g.playerTwo.Move(), size) {
		if g.playerTwo.Move() == player.X {
			return XPlayerWon
		} else if g.playerTwo.Move() == player.O {
			return OPlayerWon
		}
	}

	if sum == total(g.playerOne.Move(), size) {
		if g.playerOne.Move() == player.X {
			return XPlayerWon
		} else if g.playerOne.Move() == player.O {
			return OPlayerWon
		}
	}

	// none of the players won yet, check for empty cells
	if g.board.HasEmptyCells() {
		return InProgress
	}
	return Draw
}

// Start is the entry point for the game
// Its a blocking call, terminates when users wishes to not play any more
// It switches between two players
func (g *Game) Start() {
	fmt.Println("Initializing Game, please wait...")
	currentPlayer := g.playerOne
	g.notifyPlayers("Initializing Game, please wait...\n")
	g.printBoard()
	for {
		fmt.Println(fmt.Sprintf("%s enter your move", currentPlayer.Name()))
		currentPlayer.Notify(fmt.Sprintf("%s enter your move", currentPlayer.Name()))
		move := currentPlayer.Play(g.board)
		err := g.board.Move(move, int(currentPlayer.Move()))
		if err != nil {
			fmt.Println(err.Error())
			currentPlayer.Notify(err.Error())

			continue
		}
		g.printBoard()

		// Check for status
		switch g.Status() {
		case XPlayerWon, OPlayerWon, Draw:
			fmt.Println(g.Status())
			g.notifyPlayers(fmt.Sprintf("%s", g.Status()))
			g.finish()
			return
		}

		// set the next player
		if currentPlayer == g.playerOne {
			currentPlayer = g.playerTwo
		} else {
			currentPlayer = g.playerOne
		}
	}
}

func (g *Game) printBoard() {
	brd := fmt.Sprintf("%s\n", g.String())
	fmt.Println(brd)
	g.notifyPlayers(brd)
}

func (g *Game) notifyPlayers(data string) {
	g.playerOne.Notify(data)
	g.playerTwo.Notify(data)
}

func (g *Game) finish() {
	g.playerOne.Exit()
	g.playerTwo.Exit()
}
