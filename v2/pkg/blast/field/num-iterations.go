package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const DefaultNumIterations uint32 = 1

func NewNumIterations(val uint32) NumIterations {
	return NumIterations{true, val}
}

func NewEmptyNumIterations() NumIterations {
	return NumIterations{}
}

func DecodeJSONNumIterations(dec *gojay.Decoder, val *NumIterations) error {
	val.set = true
	return dec.Uint32(&val.val)
}

type NumIterations struct {
	set bool
	val uint32
}

func (b *NumIterations) IsSet() bool {
	return b.set
}

func (b *NumIterations) Get() uint32 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *NumIterations) Set(val uint32) {
	b.set = true
	b.val = val
}

func (b *NumIterations) Unset() {
	b.set = false
	b.val = 0
}

func (b *NumIterations) Flag() string {
	return consts.FlagNumIterations
}

func (b *NumIterations) IsDefault() bool {
	return !b.set || b.val == DefaultNumIterations
}

func (b *NumIterations) ArgString() string {
	return strconv.FormatUint(uint64(b.Get()), 10)
}

func (b *NumIterations) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}

