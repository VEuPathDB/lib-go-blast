package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const DefaultInclusionEThreshold float64 = 0.002

func NewInclusionEThreshold(val float64) InclusionEThreshold {
	return InclusionEThreshold{true, val}
}

func NewEmptyInclusionEThreshold() InclusionEThreshold {
	return InclusionEThreshold{}
}

func DecodeJSONInclusionEThreshold(dec *gojay.Decoder, val *InclusionEThreshold) error {
	val.set = true
	return dec.Float64(&val.val)
}

type InclusionEThreshold struct {
	set bool
	val float64
}

func (b *InclusionEThreshold) IsSet() bool {
	return b.set
}

func (b *InclusionEThreshold) Get() float64 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *InclusionEThreshold) Set(val float64) {
	b.set = true
	b.val = val
}

func (b *InclusionEThreshold) Unset() {
	b.set = false
	b.val = 0
}

func (b *InclusionEThreshold) Flag() string {
	return consts.FlagInclusionEThreshold
}

func (b *InclusionEThreshold) IsDefault() bool {
	return !b.IsSet() || b.val == DefaultInclusionEThreshold
}

func (b *InclusionEThreshold) ArgString() string {
	return strconv.FormatFloat(b.Get(), 'f', -1, 64)
}

func (b *InclusionEThreshold) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}