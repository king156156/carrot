package weights

import "math/rand"

// weidata  weidata
type weidata struct {
	Beforeindex int
	Alterindex  int
	Value       string
}

// Weights weights
type Weights struct {
	wei    []*weidata
	nummax int
}

// Add num and value
func (w *Weights) Add(num int, value string) {
	w.wei = append(w.wei, &weidata{
		Beforeindex: w.nummax,
		Alterindex:  w.nummax + num,
		Value:       value,
	})
	w.nummax += num
}

// Draw random num
func (w *Weights) Draw(r *rand.Rand) string {
	if len(w.wei) > 0 {
		myrand := random(r, 1, w.nummax)
		for _, data := range w.wei {
			if myrand > data.Beforeindex && myrand <= data.Alterindex {
				return data.Value
			}
		}
	}
	panic("no add")
}

// random random
func random(r *rand.Rand, min int, max int) int {
	return r.Intn(max-min+1) + min
}
