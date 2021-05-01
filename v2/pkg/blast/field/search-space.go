package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const (
	MaxSearchSpace uint8 = 127
)

func NewSearchSpace(val uint8) SearchSpace {
	return SearchSpace{true, val}
}

func NewEmptySearchSpace() SearchSpace {
	return SearchSpace{}
}

func DecodeJSONSearchSpace(dec *gojay.Decoder, val *SearchSpace) error {
	val.set = true
	return dec.DecodeUint8(&val.val)
}

type SearchSpace struct {
	set bool
	val uint8
}

func (s *SearchSpace) IsSet() bool {
	return s.set
}

func (s *SearchSpace) Get() uint8 {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *SearchSpace) Set(val uint8) {
	s.set = true
	s.val = val
}

func (s *SearchSpace) Unset() {
	s.set = false
}

func (s *SearchSpace) Flag() string {
	return consts.FlagSearchSpace
}

func (s *SearchSpace) IsDefault() bool {
	return !s.set
}

func (s *SearchSpace) ArgString() string {
	return strconv.Itoa(int(s.Get()))
}

func (s *SearchSpace) FlagString() string {
	return s.Flag() + "=" + s.ArgString()
}

func (s *SearchSpace) Validate(em bval.ValidationBuilder) {
	_ = em.Uint8LTEQ(s, MaxSearchSpace)
}
