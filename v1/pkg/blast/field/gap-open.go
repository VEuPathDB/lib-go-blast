package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"lib-go-blast/v1/pkg/blast/consts"
)

func NewGapOpen(val int32) GapOpen {
	return GapOpen{true, val}
}

func NewEmptyGapOpen() GapOpen {
	return GapOpen{}
}

func DecodeJSONGapOpen(dec *gojay.Decoder, val *GapOpen) error {
	return dec.DecodeInt32(&val.val)
}

type GapOpen struct {
	set bool
	val int32
}

func (b GapOpen) IsSet() bool {
	return b.set
}

func (b GapOpen) Get() int32 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *GapOpen) Set(val int32) {
	b.set = true
	b.val = val
}

func (b *GapOpen) Unset() {
	b.set = false
	b.val = 0
}

func (b GapOpen) Flag() string {
	return consts.FlagGapOpen
}

func (b GapOpen) IsDefault() bool {
	return !b.set
}

func (b GapOpen) ArgString() string {
	return strconv.Itoa(int(b.Get()))
}

func (b GapOpen) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}
