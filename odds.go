package odds

type American int
type Decimal float64
type Fractional struct {
	Numerator, Denominator int
}
type OddsFormat interface {
	ToDecimal() float64
}

func (a American) ToDecimal() float64 {
	f := float64(a)
	if f < 0 {
		return (-f + 100) / -f
	} else {
		return (f + 100) / 100
	}
}

func (a American) Prob() float64 {
	return 1 / a.ToDecimal()
}
