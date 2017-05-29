// Package odds provides functionality for working with betting odds.
package odds

import (
	"math/big"
)

// Odds is an interface that each odds format implements. It exists so
// functions can take any of the available odds formats.
type Odds interface {
	Probability() *big.Float
	RemoveVig(*big.Float)
	// String() string
}

// Probability returns the implied probability of the given odds.
func Probability(o Odds) *big.Float {
	return o.Probability()
}

// Vig returns the vigorish (aka juice) between two odds.
func Vig(o1, o2 Odds) *big.Float {
	p := new(big.Float)
	p.Add(o1.Probability(), o2.Probability())
	p.Add(p, big.NewFloat(-1))
	return p
}

// FairOdds returns what the odds would be without the vig.
func FairOdds(o1, o2 Odds) {
	v := Vig(o1, o2)
	v.Add(v, big.NewFloat(1))
	o1.RemoveVig(v)
	o2.RemoveVig(v)
}
