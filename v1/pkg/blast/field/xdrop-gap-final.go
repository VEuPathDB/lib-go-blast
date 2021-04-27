package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"lib-go-blast/v1/pkg/blast/consts"
)

func NewExtensionDropoffFinalGapped(val float64) ExtensionDropoffFinalGapped {
	return ExtensionDropoffFinalGapped{true, val}
}

func NewEmptyExtensionDropoffFinalGapped() ExtensionDropoffFinalGapped {
	return ExtensionDropoffFinalGapped{}
}

func DecodeJSONExtensionDropoffFinalGapped(dec *gojay.Decoder, val *ExtensionDropoffFinalGapped) error {
	val.set = true
	return dec.DecodeFloat64(&val.val)
}

type ExtensionDropoffFinalGapped struct {
	set bool
	val float64
}

func (b ExtensionDropoffFinalGapped) IsSet() bool {
	return b.set
}

func (b ExtensionDropoffFinalGapped) Get() float64 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *ExtensionDropoffFinalGapped) Set(val float64) {
	b.set = true
	b.val = val
}

func (b *ExtensionDropoffFinalGapped) Unset() {
	b.set = false
	b.val = 0
}

func (b ExtensionDropoffFinalGapped) Flag() string {
	return consts.FlagExtensionDropoffFinalGapped
}

func (b ExtensionDropoffFinalGapped) IsDefault() bool {
	return !b.IsSet()
}

func (b ExtensionDropoffFinalGapped) ArgString() string {
	return strconv.FormatFloat(b.Get(), 'f', -1, 64)
}

func (b ExtensionDropoffFinalGapped) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}