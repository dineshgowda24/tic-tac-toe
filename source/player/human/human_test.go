package human

import (
	"bytes"
	"testing"

	"github.com/dineshgowda24/tic-tac-toe/source/player"
)

type ReadWriteNopCloser struct {
	*bytes.Buffer
}

func (r *ReadWriteNopCloser) Close() error {
	return nil
}

func TestNew(t *testing.T) {

	human := New(player.O, "Dinesh", &ReadWriteNopCloser{bytes.NewBuffer([]byte("Sample Input"))})
	if human == nil {
		t.Error("human can not be nil")
	}
}

func TestHumanPlay(t *testing.T) {
	human := New(player.O, "Dinesh", &ReadWriteNopCloser{bytes.NewBuffer([]byte("1\n2\n3\n"))})
	if human == nil {
		t.Error("human can not be nil")
		t.FailNow()
	}

	tests := []int{1, 2, 3}
	for _, val := range tests {
		actual := human.Play(nil)
		if val != actual {
			t.Errorf("Expected %d, but got %d", val, actual)
		}
	}
}
