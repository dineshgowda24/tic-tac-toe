package board

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"sync"
)

const (
	// startIndex grid start index
	startIndex = 1
	// notOccupied are the cells which are not played and has 0 value
	notOccupied = 0
)

var (
	ErrNoEmptyCells      = errors.New("no empty cells")
	ErrMoveAlreadyPlayed = errors.New("move already played")
	ErrMoveOutOfBound    = errors.New("move out of bound")
)

// Board has context of the grid in tic tac toe
// 0 index is ignored and not used for simplicity and from player stand point
type Board struct {
	mu    sync.RWMutex // mutex for synchronization
	moves int          // moves represents the total number of valid moves made on board
	size  int          // size of NxN matrix, represents N
	grid  []int        // one dimensional array representing NxN matrix
}

// NewBoard returns a new Board
// Initializes the slice to size+1 since we will be ignoring index : 0
func NewBoard(size int) (*Board, error) {
	if size < 3 {
		return nil, errors.New("invalid size")
	}
	return &Board{
		size: size,
		grid: make([]int, (size*size)+1, (size*size)+1),
	}, nil
}

// Grid returns a copy of the grid at the time
// The actual grid is not shared since it can be modified
func (b *Board) Grid() []int {
	b.mu.RLock()
	defer b.mu.RUnlock()

	grid := make([]int, (b.size*b.size)+1, (b.size*b.size)+1)
	copy(grid, b.grid)
	return grid
}

func (b *Board) DeepCopy() *Board {
	b.mu.RLock()
	defer b.mu.RUnlock()

	grid := make([]int, (b.size*b.size)+1, (b.size*b.size)+1)
	copy(grid, b.grid)
	return &Board{grid: grid, size: b.size, mu: sync.RWMutex{}}
}

// GridRef returns a reference to the actual grid
// The actual grid is not shared since it can be modified
func (b *Board) GridRef() []int {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.grid[:]
}

// Size returns the size of the matrix
func (b *Board) Size() int {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.size
}

// Move adds a given value to an index
// If index is out of bound or already filled an error is returned
func (b *Board) Move(cell int, value int) error {
	valid, err := b.validMove(cell)
	if !valid {
		return err
	}

	b.mu.Lock()
	defer b.mu.Unlock()
	b.grid[cell] = value
	b.moves++
	return nil
}

func (b *Board) Reset(cell int) error {
	if cell > b.size*b.size || cell < startIndex {
		return ErrMoveOutOfBound
	}

	b.mu.Lock()
	defer b.mu.Unlock()
	b.grid[cell] = 0
	b.moves--
	return nil
}

// validMove move checks if the move is valid
// returns true if the move is valid and false if not with a reason
func (b *Board) validMove(index int) (bool, error) {
	if !b.HasEmptyCells() {
		return false, ErrNoEmptyCells
	}

	if index > b.size*b.size || index < startIndex {
		return false, ErrMoveOutOfBound
	}

	if b.grid[index] != notOccupied {
		return false, ErrMoveAlreadyPlayed
	}
	return true, nil
}

// HasEmptyCells returns true if there are any empty indexes in the matrix
func (b *Board) HasEmptyCells() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if b.moves < b.size*b.size {
		return true
	}
	return false
}

func (b *Board) String() string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	var buf bytes.Buffer
	var rows []string
	offset := 1
	for i := 1; i <= b.size*b.size; i++ {
		rows = append(rows, fmt.Sprintf("%d", b.grid[i]))
		offset++
		if offset > b.size {
			fmt.Fprintf(&buf, fmt.Sprintf("%s\n", strings.Join(rows, "|")))
			offset = 1
			rows = []string{}
		}
	}
	return buf.String()
}
