package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const MinLineLength uint16 = 1 // TODO: this should be a newtype in v2

const DefaultLineLength uint16 = 60

func NewLineLength(val uint16) LineLength {
	return LineLength{true, val}
}

func NewEmptyLineLength() LineLength {
	return LineLength{}
}

func DecodeJSONLineLength(dec *gojay.Decoder, val *LineLength) error {
	val.set = true
	return dec.Uint16(&val.val)
}

type LineLength struct {
	set bool
	val uint16
}

func (b *LineLength) IsSet() bool {
	return b.set
}

func (b *LineLength) Get() uint16 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *LineLength) Set(val uint16) {
	b.set = true
	b.val = val
}

func (b *LineLength) Unset() {
	b.set = false
	b.val = 0
}

func (b *LineLength) Flag() string {
	return consts.FlagLineLength
}

func (b *LineLength) IsDefault() bool {
	return !b.set || b.val == DefaultLineLength
}

func (b *LineLength) ArgString() string {
	return strconv.FormatUint(uint64(b.Get()), 10)
}

func (b *LineLength) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}

func (b *LineLength) Validate(em bval.ValidationBuilder) {
	_ = em.Uint16GTEQ(b, MinLineLength)
}
