package odds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	prec = 4
)

func TestProbability(t *testing.T) {
	tests := []struct {
		odds Odds
		prob string
	}{
		{American(-600), "0.8571"},
		{American(-300), "0.75"},
		{American(-250), "0.7143"},
		{American(-200), "0.6667"},
		{American(-150), "0.6"},
		{American(-110), "0.5238"},
		{American(+100), "0.5"},
		{American(+150), "0.4"},
		{American(+200), "0.3333"},
		{American(+250), "0.2857"},
		{American(+300), "0.25"},
		{American(+450), "0.1818"},

		{Decimal(5.0), "0.2"},
		{Decimal(4.0), "0.25"},
		{Decimal(3.0), "0.3333"},
		{Decimal(2.5), "0.4"},
		{Decimal(2.0), "0.5"},
		{Decimal(1.9090), "0.5238"},
		{Decimal(1.5), "0.6667"},
		{Decimal(1.1), "0.9091"},
		{Decimal(1.0), "1"},

		{Fractional{10, 1}, "0.09091"},
		{Fractional{6, 1}, "0.1429"},
		{Fractional{2, 1}, "0.3333"},
		{Fractional{10, 11}, "0.5238"},
		{Fractional{2, 5}, "0.7143"},
		{Fractional{2.5, 7}, "0.7368"},
		{Fractional{1, 3}, "0.75"},
		{Fractional{8, 11}, "0.5789"},
	}

	for _, test := range tests {
		p := Probability(test.odds)
		assert.Equal(t, test.prob, p.Text('g', prec))
	}
}

func TestVig(t *testing.T) {
	tests := []struct {
		o1  Odds
		o2  Odds
		vig string
	}{
		{American(-750), American(+604), "0.0244"},
		{American(-415), American(+365), "0.02088"},
		{American(-150), Decimal(2.30), "0.03478"},
	}

	for _, test := range tests {
		v := Vig(test.o1, test.o2)
		assert.Equal(t, test.vig, v.Text('g', prec))
	}
}

func TestFairOdds(t *testing.T) {
	tests := []struct {
		o1 Odds
		o2 Odds
		n1 string
		n2 string
	}{
		{American(-750), American(+604), "0.8613", "0.1387"},
	}

	for _, test := range tests {
		n1, n2 := FairOdds(test.o1, test.o2)
		assert.Equal(t, test.n1, n1.Text('g', prec))
		assert.Equal(t, test.n2, n2.Text('g', prec))
	}
}
