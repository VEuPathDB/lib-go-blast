package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewExtensionDropoffUngapped(val float64) ExtensionDropoffUngapped {
	return ExtensionDropoffUngapped{true, val}
}

func NewEmptyExtensionDropoffUngapped() ExtensionDropoffUngapped {
	return ExtensionDropoffUngapped{}
}

func DecodeJSONExtensionDropoffUngapped(dec *gojay.Decoder, val *ExtensionDropoffUngapped) error {
	val.set = true
	return dec.Float64(&val.val)
}

type ExtensionDropoffUngapped struct {
	set bool
	val float64
}

func (b *ExtensionDropoffUngapped) IsSet() bool {
	return b.set
}

func (b *ExtensionDropoffUngapped) Get() float64 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *ExtensionDropoffUngapped) Set(val float64) {
	b.set = true
	b.val = val
}

func (b *ExtensionDropoffUngapped) Unset() {
	b.set = false
	b.val = 0
}

func (b *ExtensionDropoffUngapped) Flag() string {
	return consts.FlagExtensionDropoffUngapped
}

func (b *ExtensionDropoffUngapped) IsDefault() bool {
	return !b.IsSet()
}

func (b *ExtensionDropoffUngapped) ArgString() string {
	return strconv.FormatFloat(b.Get(), 'f', -1, 64)
}

func (b *ExtensionDropoffUngapped) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}