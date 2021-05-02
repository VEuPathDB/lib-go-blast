package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const DefaultNumDescriptions uint32 = 500

func NewNumDescriptions(val uint32) NumDescriptions {
	return NumDescriptions{true, val}
}

func NewEmptyNumDescriptions() NumDescriptions {
	return NumDescriptions{}
}

func DecodeJSONNumDescriptions(dec *gojay.Decoder, val *NumDescriptions) error {
	val.set = true
	return dec.Uint32(&val.val)
}

type NumDescriptions struct {
	set bool
	val uint32
}

func (b *NumDescriptions) IsSet() bool {
	return b.set
}

func (b *NumDescriptions) Get() uint32 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *NumDescriptions) Set(val uint32) {
	b.set = true
	b.val = val
}

func (b *NumDescriptions) Unset() {
	b.set = false
	b.val = 0
}

func (b *NumDescriptions) Flag() string {
	return consts.FlagNumDescriptions
}

func (b *NumDescriptions) IsDefault() bool {
	return !b.set || b.val == DefaultNumDescriptions
}

func (b *NumDescriptions) ArgString() string {
	return strconv.FormatUint(uint64(b.Get()), 10)
}

func (b *NumDescriptions) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}

