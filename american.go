package odds

import (
	"math/big"
)

// American odds represent either the stake required to win $100 or
// the amount won per $100 staked (when negative or positive,
// respectively).
type American struct {
	prob *big.Float
}

// NewAmerican takes an input string like "-110" or "+250" and returns
// a new American object.
func NewAmerican(s string) *American {
	var f, stake, payout *big.Float

	f = new(big.Float)
	f, _, err := f.Parse(s, 10)
	if err != nil {
		return nil
	}

	stake = new(big.Float)
	payout = new(big.Float)

	if f.Sign() == -1 {
		f.Neg(f)
		stake.Copy(f)
	} else {
		stake = big.NewFloat(100)
	}

	payout.Add(f, big.NewFloat(100))

	return &American{
		prob: stake.Quo(stake, payout),
	}
}

// NewAmericanFromProb takes a *big.Float and returns a new American object.
func NewAmericanFromProb(p *big.Float) *American {
	return &American{prob: p}
}

func (a *American) Probability() *big.Float {
	return a.prob
}

func (a *American) RemoveVig(v *big.Float) {
	a.prob.Quo(a.prob, v)
}
