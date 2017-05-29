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
		{NewAmerican("-750"), "0.8824"},
		{NewAmerican("-600"), "0.8571"},
		{NewAmerican("-300"), "0.75"},
		{NewAmerican("-250"), "0.7143"},
		{NewAmerican("-200"), "0.6667"},
		{NewAmerican("-150"), "0.6"},
		{NewAmerican("-110"), "0.5238"},
		{NewAmerican("+100"), "0.5"},
		{NewAmerican("+150"), "0.4"},
		{NewAmerican("+200"), "0.3333"},
		{NewAmerican("+250"), "0.2857"},
		{NewAmerican("+300"), "0.25"},
		{NewAmerican("+450"), "0.1818"},

		{NewDecimal("5.0"), "0.2"},
		{NewDecimal("4.0"), "0.25"},
		{NewDecimal("3.0"), "0.3333"},
		{NewDecimal("2.5"), "0.4"},
		{NewDecimal("2.0"), "0.5"},
		{NewDecimal("1.9090"), "0.5238"},
		{NewDecimal("1.5"), "0.6667"},
		{NewDecimal("1.1"), "0.9091"},
		{NewDecimal("1.0"), "1"},

		{NewFractional("10/1"), "0.09091"},
		{NewFractional("6/1"), "0.1429"},
		{NewFractional("2/1"), "0.3333"},
		{NewFractional("10/11"), "0.5238"},
		{NewFractional("2/5"), "0.7143"},
		{NewFractional("2.5/7"), "0.7368"},
		{NewFractional("1/3"), "0.75"},
		{NewFractional("8/11"), "0.5789"},
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
		{NewAmerican("-750"), NewAmerican("+604"), "0.0244"},
		{NewAmerican("-415"), NewAmerican("+365"), "0.02088"},
		{NewAmerican("-150"), NewDecimal("2.30"), "0.03478"},
	}

	for _, test := range tests {
		v := Vig(test.o1, test.o2)
		assert.Equal(t, test.vig, v.Text('g', prec))
	}
}

func TestFairOdds(t *testing.T) {
	f := func(o Odds) string {
		return o.Probability().Text('g', prec)
	}

	a1 := NewAmerican("-750")
	a2 := NewAmerican("+604")

	assert.Equal(t, "0.8824", f(a1))
	assert.Equal(t, "0.142", f(a2))

	FairOdds(a1, a2)

	assert.Equal(t, "0.8613", f(a1))
	assert.Equal(t, "0.1387", f(a2))
}
