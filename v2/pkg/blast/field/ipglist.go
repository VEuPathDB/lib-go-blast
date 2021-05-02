package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func DecodeJSONIPGList(dec *gojay.Decoder, val *IPGList) error {
	val.set = true
	return dec.String(&val.val)
}

type IPGList struct {
	set bool
	val string
}

func (i *IPGList) IsSet() bool {
	return i.set
}

func (i *IPGList) Get() string {
	if !i.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return i.val
}

func (i *IPGList) Set(val string) {
	i.set = true
	i.val = val
}

func (i *IPGList) Unset() {
	i.set = false
	i.val = ""
}

func (i *IPGList) Flag() string {
	return consts.FlagIPGList
}

func (i *IPGList) IsDefault() bool {
	return !i.set || len(i.val) == 0
}

func (i *IPGList) FlagString() string {
	return i.Flag() + "='" + i.Get() + "'"
}

