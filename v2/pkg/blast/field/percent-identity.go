package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const (
	MaxPercentIdentity float64 = 100
	MinPercentIdentity float64 = 0
)

func NewPercentIdentity(val float64) PercentIdentity {
	return PercentIdentity{true, val}
}

func NewEmptyPercentIdentity() PercentIdentity {
	return PercentIdentity{}
}

func DecodeJSONPercentIdentity(dec *gojay.Decoder, val *PercentIdentity) error {
	val.set = true
	return dec.DecodeFloat64(&val.val)
}

type PercentIdentity struct {
	set bool
	val float64
}

func (b *PercentIdentity) IsSet() bool {
	return b.set
}

func (b *PercentIdentity) Get() float64 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *PercentIdentity) Set(val float64) {
	b.set = true
	b.val = val
}

func (b *PercentIdentity) Unset() {
	b.set = false
	b.val = 0
}

func (PercentIdentity) Flag() string {
	return consts.FlagPercentIdentity
}

func (b *PercentIdentity) IsDefault() bool {
	return !b.IsSet()
}

func (b *PercentIdentity) ArgString() string {
	return strconv.FormatFloat(b.Get(), 'f', -1, 64)
}

func (b *PercentIdentity) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}

func (b *PercentIdentity) Validate(em bval.ValidationBuilder) {
	_ = em.F64InclusiveRange(b, MinPercentIdentity, MaxPercentIdentity)
}

