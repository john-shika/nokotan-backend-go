package bigfloat

import (
	"fmt"
	"math/big"
)

type BigFloat struct {
	Value *big.Float
}

func New(value float64) *BigFloat {
	return &BigFloat{
		Value: new(big.Float).SetFloat64(value),
	}
}

func (b *BigFloat) MarshalBinary() (data []byte, err error) {
	var buff []byte
	return b.Value.Append(buff, 'g', -1), nil
}

func (b *BigFloat) UnmarshalBinary(data []byte) (err error) {
	var f *big.Float

	if f, _, err = b.Value.Parse(string(data), 0); err != nil {
		return err
	}

	b.Value = f
	return nil
}

func (b *BigFloat) MarshalText() (text []byte, err error) {
	return b.Value.MarshalText()
}

func (b *BigFloat) UnmarshalText(text []byte) (err error) {
	return b.Value.UnmarshalText(text)
}

func (b *BigFloat) Scan(s fmt.ScanState, ch rune) error {
	return b.Value.Scan(s, ch)
}

func (b *BigFloat) String() string {
	return b.Value.String()
}

func (b *BigFloat) SetPrec(prec uint) *BigFloat {
	b.Value.SetPrec(prec)
	return b
}

func (b *BigFloat) SetMode(mode big.RoundingMode) *BigFloat {
	b.Value.SetMode(mode)
	return b
}

func (b *BigFloat) SetFloat64(value float64) *BigFloat {
	b.Value.SetFloat64(value)
	return b
}

func (b *BigFloat) SetInt64(value int64) *BigFloat {
	b.Value.SetInt64(value)
	return b
}

func (b *BigFloat) SetInt(value *big.Int) *BigFloat {
	b.Value.SetInt(value)
	return b
}

func (b *BigFloat) SetUint64(value uint64) *BigFloat {
	b.Value.SetUint64(value)
	return b
}

// TODO: continued, not implemented yet
