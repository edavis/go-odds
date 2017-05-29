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

// American odds represent either: (1) The stake required to win $100
// (when negative); or (2) the amount won per $100 staked (when
// positive).
type American float64

// Decimal odds represent the amount won (profit + stake) for every $1 staked.
type Decimal float64

// Fractional odds represent the ratio of profit / stake.
type Fractional struct {
	Profit, Stake float64
}

func (a American) Probability() *big.Float {
	var stake, payout *big.Float

	if i := float64(a); i < 0 {
		stake = big.NewFloat(float64(-i))
		payout = big.NewFloat(float64(-i + 100))
	} else {
		stake = big.NewFloat(float64(100))
		payout = big.NewFloat(float64(i + 100))
	}

	return stake.Quo(stake, payout)
}

func (d Decimal) Probability() *big.Float {
	stake := big.NewFloat(1.0)
	odds := big.NewFloat(float64(d))
	stake.Quo(stake, odds)
	return stake
}

func (f Fractional) Probability() *big.Float {
	stake := big.NewFloat(float64(f.Stake))
	payout := big.NewFloat(float64(f.Profit + f.Stake))
	return stake.Quo(stake, payout)
}

// Probability returns the implied probability of the given odds.
func Probability(o Odds) *big.Float {
	return o.Probability()
}

// Vig returns the vigorish (aka juice) between two odds.
func Vig(o1, o2 Odds) *big.Float {
	p := new(big.Float).Add(o1.Probability(), o2.Probability())
	p.Add(p, big.NewFloat(float64(-1)))
	return p
}
