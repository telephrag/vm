package fu64

type Fu64 struct {
	Value float64
	Prec  float64
}

func New(v, p float64) *Fu64 {
	return &Fu64{
		Value: v,
		Prec:  p,
	}
}

func NewFromF64(v float64) *Fu64 {
	return New(v, 0)
}

func Copy(f *Fu64) *Fu64 {
	return New(f.Value, f.Prec)
}

func (f *Fu64) RemoveTail() {

}
