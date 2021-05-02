package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewWordSize(val uint32) WordSize {
	return WordSize{true, val}
}

func NewEmptyWordSize() WordSize {
	return WordSize{}
}

func DecodeJSONWordSize(dec *gojay.Decoder, val *WordSize) error {
	return dec.Uint32(&val.val)
}

type WordSize struct {
	set bool
	val uint32
}

func (b *WordSize) IsSet() bool {
	return b.set
}

func (b *WordSize) Get() uint32 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *WordSize) Set(val uint32) {
	b.set = true
	b.val = val
}

func (b *WordSize) Unset() {
	b.set = false
	b.val = 0
}

func (b *WordSize) Flag() string {
	return consts.FlagWordSize
}

func (b *WordSize) IsDefault() bool {
	return !b.set
}

func (b *WordSize) ArgString() string {
	return strconv.FormatUint(uint64(b.Get()), 10)
}

func (b *WordSize) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}

