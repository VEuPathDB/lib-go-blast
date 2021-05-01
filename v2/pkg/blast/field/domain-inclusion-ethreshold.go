package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const DefaultDomainInclusionEThreshold float64 = 0.05

func NewDomainInclusionEThreshold(val float64) DomainInclusionEThreshold {
	return DomainInclusionEThreshold{true, val}
}

func NewEmptyDomainInclusionEThreshold() DomainInclusionEThreshold {
	return DomainInclusionEThreshold{}
}

func DecodeJSONDomainInclusionEThreshold(dec *gojay.Decoder, val *DomainInclusionEThreshold) error {
	val.set = true
	return dec.DecodeFloat64(&val.val)
}

type DomainInclusionEThreshold struct {
	set bool
	val float64
}

func (b *DomainInclusionEThreshold) IsSet() bool {
	return b.set
}

func (b *DomainInclusionEThreshold) Get() float64 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *DomainInclusionEThreshold) Set(val float64) {
	b.set = true
	b.val = val
}

func (b *DomainInclusionEThreshold) Unset() {
	b.set = false
	b.val = 0
}

func (b *DomainInclusionEThreshold) Flag() string {
	return consts.FlagDomainInclusionEThreshold
}

func (b *DomainInclusionEThreshold) IsDefault() bool {
	return !b.IsSet() || b.val == DefaultDomainInclusionEThreshold
}

func (b *DomainInclusionEThreshold) ArgString() string {
	return strconv.FormatFloat(b.Get(), 'f', -1, 64)
}

func (b *DomainInclusionEThreshold) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}