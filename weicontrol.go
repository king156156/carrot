package carrot

import (
	"math/rand"
	"time"

	"github.com/king156156/carrot/weights"
)

// Weinfo weight number return value
type Weinfo struct {
	Number int
	Value  string
}

// WeightsAggre  WeightsAggre
type WeightsAggre struct {
	randSoure *rand.Rand
	aggre     map[string]*weights.Weights
}

// New new
func New() *WeightsAggre {
	return &WeightsAggre{
		randSoure: getNewRand(),
		aggre:     map[string]*weights.Weights{},
	}
}

// GetWeights GetWeights
func GetWeights(data ...Weinfo) (n *weights.Weights) {
	n = &weights.Weights{}
	for _, v := range data {
		n.Add(v.Number, v.Value)
	}
	return
}

// Add Inject this mode weight
func (g *WeightsAggre) Add(key string, w *weights.Weights) {
	g.aggre[key] = w
}

// Run Probability result after executing this mode operation
func (g *WeightsAggre) Run(key string) (value string) {
	if w, ok := g.aggre[key]; ok {
		value = w.Draw(g.randSoure)
	}
	return
}

// getNewRand getNewRand
func getNewRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
