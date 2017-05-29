package odds

import (
	"math/big"
)

// Decimal odds represent the amount won (profit + stake) for every $1 staked.
type Decimal float64

// Probability returns the implied probability of the odds.
func (d Decimal) Probability() *big.Float {
	stake := big.NewFloat(1.0)
	odds := big.NewFloat(float64(d))
	stake.Quo(stake, odds)
	return stake
}
