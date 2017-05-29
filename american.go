package odds

import (
	"math/big"
)

// American odds represent either: (1) The stake required to win $100
// (when negative); or (2) the amount won per $100 staked (when
// positive).
type American float64

// Probability returns the implied probability of the odds.
func (a American) Probability() *big.Float {
	var stake, payout *big.Float

	if i := float64(a); i < 0 {
		stake = big.NewFloat(-i)
		payout = big.NewFloat(-i + 100)
	} else {
		stake = big.NewFloat(100)
		payout = big.NewFloat(i + 100)
	}

	return stake.Quo(stake, payout)
}
