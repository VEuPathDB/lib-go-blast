package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewEmptyUseIndex() UseIndex {
	return UseIndex{}
}

func NewUseIndex(val bool) UseIndex {
	return UseIndex{true, val}
}

func DecodeJSONUseIndex(dec *gojay.Decoder, val *UseIndex) error {
	val.set = true
	return dec.DecodeBool(&val.val)
}

type UseIndex struct {
	set bool
	val bool
}

func (i *UseIndex) IsSet() bool {
	return i.set
}

func (i *UseIndex) Get() bool {
	if !i.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return i.val
}

func (i *UseIndex) Set(val bool) {
	i.set = true
	i.val = val
}

func (i *UseIndex) Unset() {
	i.set = false
}

func (UseIndex) Flag() string {
	return consts.FlagUseIndex
}

func (i *UseIndex) IsDefault() bool {
	return !i.IsSet() || !i.val
}

func (i *UseIndex) ArgString() string {
	if i.Get() {
		return "true"
	} else {
		return "false"
	}
}

func (i *UseIndex) FlagString() string {
	return i.Flag() + "=" + i.ArgString()
}
