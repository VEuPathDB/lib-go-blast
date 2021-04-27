package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewMaxIntronLength(val uint32) MaxIntronLength {
	return MaxIntronLength{true, val}
}

func NewEmptyMaxIntronLength() MaxIntronLength {
	return MaxIntronLength{}
}

func DecodeJSONMaxIntronLength(dec *gojay.Decoder, val *MaxIntronLength) error {
	val.set = true
	return dec.DecodeUint32(&val.val)
}

type MaxIntronLength struct {
	set bool
	val uint32
}

func (b MaxIntronLength) IsSet() bool {
	return b.set
}

func (b MaxIntronLength) Get() uint32 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *MaxIntronLength) Set(val uint32) {
	b.set = true
	b.val = val
}

func (b *MaxIntronLength) Unset() {
	b.set = false
	b.val = 0
}

func (b MaxIntronLength) Flag() string {
	return consts.FlagMaxIntronLength
}

func (b MaxIntronLength) IsDefault() bool {
	return !b.set
}

func (b MaxIntronLength) ArgString() string {
	return strconv.FormatUint(uint64(b.Get()), 10)
}

func (b MaxIntronLength) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}

