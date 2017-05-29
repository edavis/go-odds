package odds

import (
	"math/big"
)

// Fractional odds represent the ratio of profit / stake.
type Fractional struct {
	Profit, Stake float64
}

// Probability returns the implied probability of the odds.
func (f Fractional) Probability() *big.Float {
	stake := big.NewFloat(float64(f.Stake))
	payout := big.NewFloat(float64(f.Profit + f.Stake))
	return stake.Quo(stake, payout)
}
