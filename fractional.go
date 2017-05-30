package odds

import (
	"math/big"
	"strings"
)

// Fractional odds represent the ratio of profit/stake.
type Fractional struct {
	prob *big.Float
}

// NewFractional takes an input string like "6/1" or "2.5/7" and
// returns a new Fractional object.
func NewFractional(s string) *Fractional {
	var stake, payout *big.Float
	var err error

	ss := strings.Split(s, "/")
	n := ss[0]
	d := ss[1]

	stake = new(big.Float)
	payout = new(big.Float)

	stake, _, err = stake.Parse(d, 10)
	if err != nil {
		return nil
	}

	payout, _, err = payout.Parse(n, 10)
	if err != nil {
		return nil
	}

	payout.Add(payout, stake)

	return &Fractional{
		prob: stake.Quo(stake, payout),
	}
}

func (f *Fractional) Probability() *big.Float {
	return f.prob
}

func (f *Fractional) RemoveVig(v *big.Float) {
	f.prob.Quo(f.prob, v)
}
