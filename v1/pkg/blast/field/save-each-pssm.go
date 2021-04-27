package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewSaveEachPSSM(val bool) SaveEachPSSM {
	return SaveEachPSSM{true, val}
}

func NewEmptySaveEachPSSM() SaveEachPSSM {
	return SaveEachPSSM{}
}

func DecodeJSONSaveEachPSSM(dec *gojay.Decoder, val *SaveEachPSSM) error {
	val.set = true
	return dec.DecodeBool(&val.val)
}

type SaveEachPSSM struct {
	set bool
	val bool
}

func (s SaveEachPSSM) IsSet() bool {
	return s.set
}

func (s SaveEachPSSM) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *SaveEachPSSM) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *SaveEachPSSM) Unset() {
	s.set = false
}

func (s SaveEachPSSM) Flag() string {
	return consts.FlagSaveEachPSSM
}

func (s SaveEachPSSM) IsDefault() bool {
	return !s.set || !s.val
}

func (s SaveEachPSSM) FlagString() string {
	return s.Flag()
}