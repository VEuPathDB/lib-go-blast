package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"lib-go-blast/v1/pkg/blast/consts"
)

func NewThreshold(val float64) Threshold {
	return Threshold{true, val}
}

func NewEmptyThreshold() Threshold {
	return Threshold{}
}

func DecodeJSONThreshold(dec *gojay.Decoder, val *Threshold) error {
	val.set = true
	return dec.DecodeFloat64(&val.val)
}

type Threshold struct {
	set bool
	val float64
}

func (b Threshold) IsSet() bool {
	return b.set
}

func (b Threshold) Get() float64 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *Threshold) Set(val float64) {
	b.set = true
	b.val = val
}

func (b *Threshold) Unset() {
	b.set = false
	b.val = 0
}

func (b Threshold) Flag() string {
	return consts.FlagThreshold
}

func (b Threshold) IsDefault() bool {
	return !b.IsSet()
}

func (b Threshold) ArgString() string {
	return strconv.FormatFloat(b.Get(), 'f', -1, 64)
}

func (b Threshold) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}