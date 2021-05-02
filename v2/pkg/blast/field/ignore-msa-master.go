package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewIgnoreMSAMaster(val bool) IgnoreMSAMaster {
	return IgnoreMSAMaster{true, val}
}

func NewEmptyIgnoreMSAMaster() IgnoreMSAMaster {
	return IgnoreMSAMaster{}
}

func DecodeJSONIgnoreMSAMaster(dec *gojay.Decoder, val *IgnoreMSAMaster) error {
	val.set = true
	return dec.Bool(&val.val)
}

type IgnoreMSAMaster struct {
	set bool
	val bool
}

func (s *IgnoreMSAMaster) IsSet() bool {
	return s.set
}

func (s *IgnoreMSAMaster) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *IgnoreMSAMaster) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *IgnoreMSAMaster) Unset() {
	s.set = false
}

func (s *IgnoreMSAMaster) Flag() string {
	return consts.FlagIgnoreMSAMaster
}

func (s *IgnoreMSAMaster) IsDefault() bool {
	return !s.set || !s.val
}

func (s *IgnoreMSAMaster) FlagString() string {
	return s.Flag()
}

