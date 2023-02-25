package human

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/dineshgowda24/tic-tac-toe/source/board"
	"github.com/dineshgowda24/tic-tac-toe/source/player"
)

var _ player.Player = (*Human)(nil)

// Human implements player interface
type Human struct {
	move   player.Move        // Move represents X or 0
	stream io.ReadWriteCloser // something where we can read, write and close
	name   string
}

// New returns a new human
func New(move player.Move, name string, stream io.ReadWriteCloser) *Human {
	return &Human{
		move:   move,
		stream: stream,
		name:   name,
	}
}

// Move returns the selected symbol
func (h *Human) Move() player.Move {
	return h.move
}

func (h *Human) Name() string {
	return h.name
}

// Play reads the input of player from STDIN
func (h *Human) Play(grid *board.Board) int {
	reader := bufio.NewReader(h.stream)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("unable to read input, try again")
			continue
		}
		fmt.Printf("%#v\n", input)
		input = strings.Replace(strings.Replace(input, "\r", "", -1), "\n", "", -1)
		fmt.Printf("%#v\n", input)
		index, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input")
			continue
		}
		return index
	}
}

// Notify is notifies some data to player
func (h *Human) Notify(data string) error {
	_, err := h.stream.Write([]byte(data))
	return err
}

// Exit closes the stream
func (h *Human) Exit() error {
	return h.stream.Close()
}
