package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewWindowMaskerTaxID(val int32) WindowMaskerTaxID {
	return WindowMaskerTaxID{true, val}
}

func NewEmptyWindowMaskerTaxID() WindowMaskerTaxID {
	return WindowMaskerTaxID{}
}

func DecodeJSONWindowMaskerTaxID(dec *gojay.Decoder, val *WindowMaskerTaxID) error {
	return dec.Int32(&val.val)
}

type WindowMaskerTaxID struct {
	set bool
	val int32
}

func (b WindowMaskerTaxID) IsSet() bool {
	return b.set
}

func (b WindowMaskerTaxID) Get() int32 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *WindowMaskerTaxID) Set(val int32) {
	b.set = true
	b.val = val
}

func (b *WindowMaskerTaxID) Unset() {
	b.set = false
	b.val = 0
}

func (b WindowMaskerTaxID) Flag() string {
	return consts.FlagWindowMaskerTaxID
}

func (b WindowMaskerTaxID) IsDefault() bool {
	return !b.set
}

func (b WindowMaskerTaxID) ArgString() string {
	return strconv.Itoa(int(b.Get()))
}

func (b WindowMaskerTaxID) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}
