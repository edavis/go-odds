package odds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAmerican(t *testing.T) {
	var tests = []struct {
		original int
		decimal  float64
		prob     float64
	}{
		{-139, 1.719, 0.5817},
		{+126, 2.260, 0.4424},
		{-170, 1.588, 0.6297},
		{+160, 2.600, 0.3846},
	}
	for _, obj := range tests {
		a := American(obj.original)
		assert.InEpsilon(t, a.ToDecimal(), obj.decimal, 0.001)
		assert.InEpsilon(t, a.Prob(), obj.prob, 0.001)
	}
}
