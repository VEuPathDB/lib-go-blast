package field

import (
	"github.com/francoispqt/gojay"
	"lib-go-blast/v1/pkg/blast/consts"
)

func NewSumStats(val bool) SumStats {
	return SumStats{true, val}
}

func NewEmptySumStats() SumStats {
	return SumStats{}
}

func DecodeJSONSumStats(dec *gojay.Decoder, val *SumStats) error {
	val.set = true
	return dec.DecodeBool(&val.val)
}

type SumStats struct {
	set bool
	val bool
}

func (s SumStats) IsSet() bool {
	return s.set
}

func (s SumStats) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *SumStats) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *SumStats) Unset() {
	s.set = false
}

func (s SumStats) Flag() string {
	return consts.FlagSumStats
}

func (s SumStats) IsDefault() bool {
	return !s.set
}

func (s SumStats) FlagString() string {
	if s.Get() {
		return s.Flag() + "=true"
	} else {
		return s.Flag() + "=false"
	}
}