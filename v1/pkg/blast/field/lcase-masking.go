package field

import (
	"github.com/francoispqt/gojay"
	"lib-go-blast/v1/pkg/blast/consts"
)

func NewLowercaseMasking(val bool) LowercaseMasking {
	return LowercaseMasking{true, val}
}

func NewEmptyLowercaseMasking() LowercaseMasking {
	return LowercaseMasking{}
}

func DecodeJSONLowercaseMasking(dec *gojay.Decoder, val *LowercaseMasking) error {
	val.set = true
	return dec.DecodeBool(&val.val)
}

type LowercaseMasking struct {
	set bool
	val bool
}

func (s LowercaseMasking) IsSet() bool {
	return s.set
}

func (s LowercaseMasking) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *LowercaseMasking) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *LowercaseMasking) Unset() {
	s.set = false
}

func (s LowercaseMasking) Flag() string {
	return consts.FlagLowercaseMasking
}

func (s LowercaseMasking) IsDefault() bool {
	return !s.set || !s.val
}

func (s LowercaseMasking) FlagString() string {
	return s.Flag()
}