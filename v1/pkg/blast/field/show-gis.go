package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewShowGIs(val bool) ShowGIs {
	return ShowGIs{true, val}
}

func NewEmptyShowGIs() ShowGIs {
	return ShowGIs{}
}

func DecodeJSONShowGIs(dec *gojay.Decoder, val *ShowGIs) error {
	val.set = true
	return dec.Bool(&val.val)
}

type ShowGIs struct {
	set bool
	val bool
}

func (s ShowGIs) IsSet() bool {
	return s.set
}

func (s ShowGIs) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *ShowGIs) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *ShowGIs) Unset() {
	s.set = false
}

func (s ShowGIs) Flag() string {
	return consts.FlagShowGIs
}

func (s ShowGIs) IsDefault() bool {
	return !s.set || !s.val
}

func (s ShowGIs) FlagString() string {
	return s.Flag()
}