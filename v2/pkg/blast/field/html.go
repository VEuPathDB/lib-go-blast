package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewHTML(val bool) HTML {
	return HTML{true, val}
}

func NewEmptyHTML() HTML {
	return HTML{}
}

func DecodeJSONHTML(dec *gojay.Decoder, val *HTML) error {
	val.set = true
	return dec.Bool(&val.val)
}

type HTML struct {
	set bool
	val bool
}

func (s *HTML) IsSet() bool {
	return s.set
}

func (s *HTML) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *HTML) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *HTML) Unset() {
	s.set = false
}

func (s *HTML) Flag() string {
	return consts.FlagHTML
}

func (s *HTML) IsDefault() bool {
	return !s.set || !s.val
}

func (s *HTML) FlagString() string {
	return s.Flag()
}