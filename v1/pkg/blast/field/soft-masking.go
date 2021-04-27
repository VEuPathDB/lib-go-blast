package field

import (
	"github.com/francoispqt/gojay"
	"lib-go-blast/v1/pkg/blast/consts"
)

func NewEmptySoftMasking() SoftMasking {
	return SoftMasking{}
}

func NewSoftMasking(val bool) SoftMasking {
	return SoftMasking{true, val}
}

func DecodeJSONSoftMasking(dec *gojay.Decoder, val *SoftMasking) error {
	val.set = true
	return dec.DecodeBool(&val.val)
}

type SoftMasking struct {
	set bool
	val bool
}

func (s SoftMasking) IsSet() bool {
	return s.set
}

func (s SoftMasking) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *SoftMasking) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *SoftMasking) Unset() {
	s.set = false
}

func (SoftMasking) Flag() string {
	return consts.FlagSoftMasking
}

func (s SoftMasking) IsDefault() bool {
	return !s.IsSet()
}

func (s SoftMasking) ArgString() string {
	if s.Get() {
		return "true"
	} else {
		return "false"
	}
}

func (s SoftMasking) FlagString() string {
	return s.Flag() + "=" + s.ArgString()
}
