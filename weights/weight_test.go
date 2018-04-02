package weights

import (
	"math/rand"
	"testing"
	"time"
)

func Test_Add(t *testing.T) {
	s := &Weights{}
	var test = []struct {
		input    int
		callback string
		want     int
		wantbool bool
	}{
		{input: 1, callback: "A", want: 1, wantbool: true},
		{input: 15, callback: "B", want: 15, wantbool: false},
		{input: 700, callback: "C", want: 716, wantbool: true},
	}

	for _, value := range test {
		nummaxed := s.nummax
		if s.Add(value.input, value.callback); (s.nummax == value.want) == !value.wantbool {
			t.Errorf("weights.nummax =%v weights.Add(%v)  weights.nummax =%v, want = %v , %v ", nummaxed, value.input, s.nummax, value.want, !value.wantbool)
		}
	}
}

func Test_Weights(t *testing.T) {
	s := &Weights{}
	var test = []struct {
		input    int
		callback string
		want     int
		wantbool bool
	}{
		{input: 1, callback: "A", want: 1, wantbool: true},
		{input: 15, callback: "B", want: 15, wantbool: false},
		{input: 700, callback: "C", want: 716, wantbool: true},
	}

	defer func() {
		if r := recover(); r == nil {
			return
		}
	}()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, value := range test {
		nummaxed := s.nummax
		if s.Add(value.input, value.callback); (s.nummax != value.want) == value.wantbool {
			t.Errorf("weights.nummax =%v weights.Add(%v)  weights.nummax =%v, want = %v , %v ", nummaxed, value.input, s.nummax, value.want, value.wantbool)
		}
		result := s.Draw(r)
		ok := false
		for _, data := range test {
			if result == data.callback {
				ok = true
			}
		}
		if !ok {
			t.Errorf("weights.MugleSymbol = %v not callback message in test", result)
		}
	}
	// panic
	errwei := &Weights{}
	errwei.Draw(r)
}
