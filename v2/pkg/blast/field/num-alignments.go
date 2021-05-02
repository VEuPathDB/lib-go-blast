package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const DefaultNumAlignments uint32 = 250

func NewNumAlignments(val uint32) NumAlignments {
	return NumAlignments{true, val}
}

func NewEmptyNumAlignments() NumAlignments {
	return NumAlignments{}
}

func DecodeJSONNumAlignments(dec *gojay.Decoder, val *NumAlignments) error {
	val.set = true
	return dec.Uint32(&val.val)
}

type NumAlignments struct {
	set bool
	val uint32
}

func (b *NumAlignments) IsSet() bool {
	return b.set
}

func (b *NumAlignments) Get() uint32 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *NumAlignments) Set(val uint32) {
	b.set = true
	b.val = val
}

func (b *NumAlignments) Unset() {
	b.set = false
	b.val = 0
}

func (b *NumAlignments) Flag() string {
	return consts.FlagNumAlignments
}

func (b *NumAlignments) IsDefault() bool {
	return !b.set || b.val == DefaultNumAlignments
}

func (b *NumAlignments) ArgString() string {
	return strconv.FormatUint(uint64(b.Get()), 10)
}

func (b *NumAlignments) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}

