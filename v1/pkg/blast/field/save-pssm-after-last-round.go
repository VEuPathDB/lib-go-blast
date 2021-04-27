package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewSavePSSMAfterLastRound(val bool) SavePSSMAfterLastRound {
	return SavePSSMAfterLastRound{true, val}
}

func NewEmptySavePSSMAfterLastRound() SavePSSMAfterLastRound {
	return SavePSSMAfterLastRound{}
}

func DecodeJSONSavePSSMAfterLastRound(dec *gojay.Decoder, val *SavePSSMAfterLastRound) error {
	val.set = true
	return dec.DecodeBool(&val.val)
}

type SavePSSMAfterLastRound struct {
	set bool
	val bool
}

func (s SavePSSMAfterLastRound) IsSet() bool {
	return s.set
}

func (s SavePSSMAfterLastRound) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *SavePSSMAfterLastRound) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *SavePSSMAfterLastRound) Unset() {
	s.set = false
}

func (s SavePSSMAfterLastRound) Flag() string {
	return consts.FlagSavePSSMAfterLastRound
}

func (s SavePSSMAfterLastRound) IsDefault() bool {
	return !s.set || !s.val
}

func (s SavePSSMAfterLastRound) FlagString() string {
	return s.Flag()
}