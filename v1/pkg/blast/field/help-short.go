package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewShortHelp(val bool) ShortHelp {
	return ShortHelp{true, val}
}

func NewEmptyShortHelp() ShortHelp {
	return ShortHelp{}
}

func DecodeJSONShortHelp(dec *gojay.Decoder, val *ShortHelp) error {
	val.set = true
	return dec.DecodeBool(&val.val)
}

type ShortHelp struct {
	set bool
	val bool
}

func (s ShortHelp) IsSet() bool {
	return s.set
}

func (s ShortHelp) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *ShortHelp) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *ShortHelp) Unset() {
	s.set = false
}

func (s ShortHelp) Flag() string {
	return consts.FlagShortHelp
}

func (s ShortHelp) IsDefault() bool {
	return !s.set || !s.val
}

func (s ShortHelp) FlagString() string {
	return s.Flag()
}

