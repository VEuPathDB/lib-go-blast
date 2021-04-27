package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"lib-go-blast/v1/pkg/blast/consts"
)

func NewGapExtend(val int32) GapExtend {
	return GapExtend{true, val}
}

func NewEmptyGapExtend() GapExtend {
	return GapExtend{}
}

func DecodeJSONGapExtend(dec *gojay.Decoder, val *GapExtend) error {
	return dec.DecodeInt32(&val.val)
}

type GapExtend struct {
	set bool
	val int32
}

func (b GapExtend) IsSet() bool {
	return b.set
}

func (b GapExtend) Get() int32 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *GapExtend) Set(val int32) {
	b.set = true
	b.val = val
}

func (b *GapExtend) Unset() {
	b.set = false
	b.val = 0
}

func (b GapExtend) Flag() string {
	return consts.FlagGapExtend
}

func (b GapExtend) IsDefault() bool {
	return !b.set
}

func (b GapExtend) ArgString() string {
	return strconv.Itoa(int(b.Get()))
}

func (b GapExtend) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}
