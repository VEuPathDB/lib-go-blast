package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewUseSmithWatermanTraceback(val bool) UseSmithWatermanTraceback {
	return UseSmithWatermanTraceback{true, val}
}

func NewEmptyUseSmithWatermanTraceback() UseSmithWatermanTraceback {
	return UseSmithWatermanTraceback{}
}

func DecodeJSONUseSmithWatermanTraceback(dec *gojay.Decoder, val *UseSmithWatermanTraceback) error {
	val.set = true
	return dec.DecodeBool(&val.val)
}

type UseSmithWatermanTraceback struct {
	set bool
	val bool
}

func (s UseSmithWatermanTraceback) IsSet() bool {
	return s.set
}

func (s UseSmithWatermanTraceback) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *UseSmithWatermanTraceback) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *UseSmithWatermanTraceback) Unset() {
	s.set = false
}

func (s UseSmithWatermanTraceback) Flag() string {
	return consts.FlagUseSmithWatermanTraceback
}

func (s UseSmithWatermanTraceback) IsDefault() bool {
	return !s.set || !s.val
}

func (s UseSmithWatermanTraceback) FlagString() string {
	return s.Flag()
}