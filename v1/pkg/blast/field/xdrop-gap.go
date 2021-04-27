package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewExtensionDropoffPrelimGapped(val float64) ExtensionDropoffPrelimGapped {
	return ExtensionDropoffPrelimGapped{true, val}
}

func NewEmptyExtensionDropoffPrelimGapped() ExtensionDropoffPrelimGapped {
	return ExtensionDropoffPrelimGapped{}
}

func DecodeJSONExtensionDropoffPrelimGapped(dec *gojay.Decoder, val *ExtensionDropoffPrelimGapped) error {
	val.set = true
	return dec.DecodeFloat64(&val.val)
}

type ExtensionDropoffPrelimGapped struct {
	set bool
	val float64
}

func (b ExtensionDropoffPrelimGapped) IsSet() bool {
	return b.set
}

func (b ExtensionDropoffPrelimGapped) Get() float64 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *ExtensionDropoffPrelimGapped) Set(val float64) {
	b.set = true
	b.val = val
}

func (b *ExtensionDropoffPrelimGapped) Unset() {
	b.set = false
	b.val = 0
}

func (b ExtensionDropoffPrelimGapped) Flag() string {
	return consts.FlagExtensionDropoffPrelimGapped
}

func (b ExtensionDropoffPrelimGapped) IsDefault() bool {
	return !b.IsSet()
}

func (b ExtensionDropoffPrelimGapped) ArgString() string {
	return strconv.FormatFloat(b.Get(), 'f', -1, 64)
}

func (b ExtensionDropoffPrelimGapped) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}