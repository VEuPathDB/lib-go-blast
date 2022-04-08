package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewPseudocount(val int32) Pseudocount {
	return Pseudocount{true, val}
}

func NewEmptyPseudocount() Pseudocount {
	return Pseudocount{}
}

func DecodeJSONPseudocount(dec *gojay.Decoder, val *Pseudocount) error {
	val.set = true
	return dec.Int32(&val.val)
}

type Pseudocount struct {
	set bool
	val int32
}

func (b *Pseudocount) IsSet() bool {
	return b.set
}

func (b *Pseudocount) Get() int32 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *Pseudocount) Set(val int32) {
	b.set = true
	b.val = val
}

func (b *Pseudocount) Unset() {
	b.set = false
	b.val = 0
}

func (b *Pseudocount) Flag() string {
	return consts.FlagPseudocount
}

func (b *Pseudocount) IsDefault() bool {
	return !b.set
}

func (b *Pseudocount) ArgString() string {
	return strconv.Itoa(int(b.Get()))
}

func (b *Pseudocount) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}
