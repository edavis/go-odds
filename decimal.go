package odds

import (
	"math/big"
)

// Decimal odds represent the amount won (profit + stake) for every $1 staked.
type Decimal struct {
	prob *big.Float
}

// NewDecimal takes an input string like "1.9090" or "2.4" and returns
// a new Decimal object.
func NewDecimal(s string) *Decimal {
	f := new(big.Float)
	f, _, err := f.Parse(s, 10)
	if err != nil {
		return nil
	}
	return &Decimal{
		prob: f.Quo(big.NewFloat(1), f),
	}
}

func (d *Decimal) Probability() *big.Float {
	return d.prob
}

func (d *Decimal) RemoveVig(v *big.Float) {
	d.prob.Quo(d.prob, v)
}
