package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewNonGreedy(val bool) NonGreedy {
	return NonGreedy{true, val}
}

func NewEmptyNonGreedy() NonGreedy {
	return NonGreedy{}
}

func DecodeJSONNonGreedy(dec *gojay.Decoder, val *NonGreedy) error {
	val.set = true
	return dec.Bool(&val.val)
}

type NonGreedy struct {
	set bool
	val bool
}

func (s *NonGreedy) IsSet() bool {
	return s.set
}

func (s *NonGreedy) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *NonGreedy) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *NonGreedy) Unset() {
	s.set = false
}

func (s *NonGreedy) Flag() string {
	return consts.FlagNonGreedy
}

func (s *NonGreedy) IsDefault() bool {
	return !s.set || !s.val
}

func (s *NonGreedy) ArgString() string {
	return ""
}

func (s *NonGreedy) FlagString() string {
	return s.Flag()
}

