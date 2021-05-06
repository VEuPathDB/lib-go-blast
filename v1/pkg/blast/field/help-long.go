package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewLongHelp(val bool) LongHelp {
	return LongHelp{true, val}
}

func NewEmptyLongHelp() LongHelp {
	return LongHelp{}
}

func DecodeJSONLongHelp(dec *gojay.Decoder, val *LongHelp) error {
	val.set = true
	return dec.Bool(&val.val)
}

type LongHelp struct {
	set bool
	val bool
}

func (s LongHelp) IsSet() bool {
	return s.set
}

func (s LongHelp) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *LongHelp) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *LongHelp) Unset() {
	s.set = false
}

func (s LongHelp) Flag() string {
	return consts.FlagLongHelp
}

func (s LongHelp) IsDefault() bool {
	return !s.set || !s.val
}

func (s LongHelp) FlagString() string {
	return s.Flag()
}

