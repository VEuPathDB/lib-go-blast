package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

const DefaultGapTrigger float64 = 22

func NewGapTrigger(val float64) GapTrigger {
	return GapTrigger{true, val}
}

func NewEmptyGapTrigger() GapTrigger {
	return GapTrigger{}
}

func DecodeJSONGapTrigger(dec *gojay.Decoder, val *GapTrigger) error {
	val.set = true
	return dec.DecodeFloat64(&val.val)
}

type GapTrigger struct {
	set bool
	val float64
}

func (b GapTrigger) IsSet() bool {
	return b.set
}

func (b GapTrigger) Get() float64 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *GapTrigger) Set(val float64) {
	b.set = true
	b.val = val
}

func (b *GapTrigger) Unset() {
	b.set = false
	b.val = 0
}

func (b GapTrigger) Flag() string {
	return consts.FlagGapTrigger
}

func (b GapTrigger) IsDefault() bool {
	return !b.IsSet() || b.val == DefaultGapTrigger
}

func (b GapTrigger) ArgString() string {
	return strconv.FormatFloat(b.Get(), 'f', -1, 64)
}

func (b GapTrigger) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}