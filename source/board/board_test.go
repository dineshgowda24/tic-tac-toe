package board

import (
	"errors"
	"strings"
	"testing"
)

func TestNewBoard(t *testing.T) {
	tests := []struct {
		size       int
		matrixSize int
		err        error
	}{
		{
			size:       3,
			matrixSize: 3*3 + 1,
			err:        nil,
		},
		{
			size:       3,
			matrixSize: 3*3 + 1,
			err:        nil,
		},
		{
			size:       2,
			matrixSize: 0,
			err:        errors.New("invalid size"),
		},
	}
	for _, test := range tests {
		b, err := NewBoard(test.size)
		if err != nil {
			if err.Error() != test.err.Error() {
				t.Errorf("expected %s but got %s", test.err, err)
				t.FailNow()
			}
		} else {
			if test.err == nil {
				if b == nil {
					t.Errorf("board should not be nil")
					t.FailNow()
				}

				if b.Size() != test.size {
					t.Errorf("expected %d but got %d", test.size, b.Size())
					t.FailNow()
				}

				if b.HasEmptyCells() != true {
					t.Errorf("expected %v but got %v", true, b.HasEmptyCells())
					t.FailNow()
				}
				grid := b.Grid()
				if cap(grid) != (test.size*test.size)+1 {
					t.Errorf("expected %d but got %d", (test.size*test.size)+1, len(grid))
					t.FailNow()
				}
			}
		}
	}
}

func TestMovesOnBoard(t *testing.T) {
	tests := []struct {
		cell  int
		value int
		err   error
	}{
		{
			cell:  1,
			value: 3,
			err:   nil,
		},
		{
			cell:  1,
			value: 3,
			err:   ErrMoveAlreadyPlayed,
		},
		{
			cell:  0,
			value: -3,
			err:   ErrMoveOutOfBound,
		},
		{
			cell:  123313133113,
			value: -3,
			err:   ErrMoveOutOfBound,
		},
	}
	b, err := NewBoard(3)
	if err != nil {
		t.Errorf("expected %s but got %s", errors.New(""), err)
		t.FailNow()

	}
	for _, test := range tests {
		err := b.Move(test.cell, test.value)
		if err != test.err {
			t.Errorf("expected error %s but got %s", test.err, err)
			t.FailNow()
		}

		if err != nil {
			continue
		}

		grid := b.Grid()
		if grid[test.cell] != test.value {
			t.Errorf("expected value %d but got %d", test.value, grid[test.cell])
			t.FailNow()
		}

		if b.HasEmptyCells() != true {
			t.Errorf("expected empty cells %v but got %v", true, b.HasEmptyCells())
			t.FailNow()
		}
	}
}

func TestPrintBoard3X3(t *testing.T) {
	b, err := NewBoard(3)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
		t.FailNow()
	}
	str := `0|0|0
0|0|0
0|0|0`
	if strings.TrimSpace(b.String()) != strings.TrimSpace(str) {
		t.Errorf("expected %s  but got %s", str, b.String())
		t.FailNow()
	}

	err = b.Move(6, 3)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
		t.FailNow()
	}
	str = `0|0|0
0|0|3
0|0|0`
	if strings.TrimSpace(b.String()) != strings.TrimSpace(str) {
		t.Errorf("expected %s  but got %s", str, b.String())
		t.FailNow()
	}
}

func TestPrintBoard4X4(t *testing.T) {
	b, err := NewBoard(4)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
		t.FailNow()
	}
	str := `0|0|0|0
0|0|0|0
0|0|0|0
0|0|0|0`
	if strings.TrimSpace(b.String()) != strings.TrimSpace(str) {
		t.Errorf("expected %s  but got %s", str, b.String())
		t.FailNow()
	}

	err = b.Move(16, -3)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
		t.FailNow()
	}
	str = `0|0|0|0
0|0|0|0
0|0|0|0
0|0|0|-3`
	if strings.TrimSpace(b.String()) != strings.TrimSpace(str) {
		t.Errorf("expected %s  but got %s", str, b.String())
		t.FailNow()
	}
}
