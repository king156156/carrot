package carrot

import (
	"testing"
)

func Benchmark_weicontrol_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := GetWeights()
		m.Add(1, "A")

		p := New()
		p.Add("test", m)
	}
}

func Benchmark_Run(b *testing.B) {
	testkey := "test"
	m := GetWeights()
	m.Add(1, "A")

	p := New()
	p.Add(testkey, m)

	for i := 0; i < b.N; i++ {
		p.Run(testkey)
	}
}

func Test_Add(t *testing.T) {
	m := GetWeights()
	m.Add(2, "BonusGame")
	m.Add(3, "FreeGame")
	m.Add(700, "BaseGame")

	waggre := New()
	waggre.Add("test", m)
	_, ok := waggre.aggre["test"]
	if !ok {
		t.Error("Error Null value Need to exist.")
	}
}

func Test_Run(t *testing.T) {
	testA := "Probability"
	m := GetWeights()
	m.Add(100, "Bonus1")
	m.Add(30, "Bonus2")
	m.Add(2, "Bonus3")

	testB := "Probability"
	m1 := GetWeights()
	m1.Add(1, "A")
	m1.Add(3, "B")

	waggre := New()
	waggre.Add(testA, m)
	waggre.Add(testB, m1)

	v := waggre.Run(testA)
	if v != "Bonus1" && v != "Bonus2" && v != "Bonus3" {
		t.Error("Error does not exist.")
	}
	v = waggre.Run(testB)
	if v != "A" && v != "B" {
		t.Error("Error does not exist.")
	}
}
