package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
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

func (s UseIndex) IsSet() bool {
	return s.set
}

func (s UseIndex) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *UseIndex) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *UseIndex) Unset() {
	s.set = false
}

func (UseIndex) Flag() string {
	return consts.FlagUseIndex
}

func (s UseIndex) IsDefault() bool {
	return !s.IsSet() || !s.val
}

func (s UseIndex) ArgString() string {
	if s.Get() {
		return "true"
	} else {
		return "false"
	}
}

func (s UseIndex) FlagString() string {
	return s.Flag() + "=" + s.ArgString()
}
