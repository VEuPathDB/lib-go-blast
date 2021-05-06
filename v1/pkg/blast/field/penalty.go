package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

const MaxPenalty int32 = 0

func NewPenalty(val int32) Penalty {
	return Penalty{true, val}
}

func NewEmptyPenalty() Penalty {
	return Penalty{}
}

func DecodeJSONPenalty(dec *gojay.Decoder, val *Penalty) error {
	return dec.Int32(&val.val)
}

type Penalty struct {
	set bool
	val int32
}

func (b Penalty) IsSet() bool {
	return b.set
}

func (b Penalty) Get() int32 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *Penalty) Set(val int32) {
	b.set = true
	b.val = val
}

func (b *Penalty) Unset() {
	b.set = false
	b.val = 0
}

func (b Penalty) Flag() string {
	return consts.FlagPenalty
}

func (b Penalty) IsDefault() bool {
	return !b.set
}

func (b Penalty) ArgString() string {
	return strconv.Itoa(int(b.Get()))
}

func (b Penalty) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}
