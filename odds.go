// Package odds provides functionality for working with betting odds.
package odds

import (
	"math/big"
)

// Odds is an interface that each odds format implements. It exists so
// functions can take any of the available odds formats.
type Odds interface {
	Probability() *big.Float
}

// Probability returns the implied probability of the given odds.
func Probability(o Odds) *big.Float {
	return o.Probability()
}

// Vig returns the vigorish (aka juice) between two odds.
func Vig(o1, o2 Odds) *big.Float {
	p := o1.Probability()
	p.Add(p, o2.Probability())
	p.Add(p, big.NewFloat(-1))
	return p
}

// FairOdds returns what the odds would be without the vig.
func FairOdds(o1, o2 Odds) (*big.Float, *big.Float) {
	v := Vig(o1, o2)
	v.Add(v, big.NewFloat(1))

	p1 := o1.Probability()
	p2 := o2.Probability()

	return p1.Quo(p1, v), p2.Quo(p2, v)
}
