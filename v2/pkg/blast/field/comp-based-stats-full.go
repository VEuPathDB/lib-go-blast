package field

import (
	"errors"
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const DefaultFullCompBasedStats = FullCompBasedStatsScoreAdjustment

func DecodeJSONFullCompBasedStats(dec *gojay.Decoder, val *FullCompBasedStats) error {
	var tmp string
	if err := dec.String(&tmp); err != nil {
		return err
	}

	val.set = true

	switch tmp {
	case "D", "d":
		val.val = FullCompBasedStatsScoreAdjustment
	case "0", "F", "f":
		val.val = FullCompBasedStatsNone
	case "1":
		val.val = FullCompBasedStatsStatistics
	case "2", "T", "t":
		val.val = FullCompBasedStatsScoreAdjustment
	case "3":
		val.val = FullCompBasedStatsScoreAdjustmentUnconditional
	default:
		return errors.New("Invalid comp based stats flag value.")
	}

	return nil
}

type FullCompBasedStats struct {
	set bool
	val FullCompBasedStatsEnum
}

func (f *FullCompBasedStats) IsSet() bool {
	return f.set
}

func (f *FullCompBasedStats) Get() FullCompBasedStatsEnum {
	if !f.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return f.val
}

func (f *FullCompBasedStats) Set(val FullCompBasedStatsEnum) {
	f.set = true
	f.val = val
}

func (f *FullCompBasedStats) Unset() {
	f.set = false
}

func (f *FullCompBasedStats) Flag() string {
	return consts.FlagCompBasedStats
}

func (f *FullCompBasedStats) IsDefault() bool {
	return !f.set || f.val == DefaultFullCompBasedStats
}

func (f *FullCompBasedStats) FlagString() string {
	return f.Flag() + "=" + f.Get().String()
}

func (f *FullCompBasedStats) Validate(dm bval.ValidationBuilder) {
	if !f.IsDefault() {
		switch f.val {
		case FullCompBasedStatsNone, FullCompBasedStatsStatistics, FullCompBasedStatsScoreAdjustment, FullCompBasedStatsScoreAdjustmentUnconditional:
			// do nothing
		default:
			dm.InvalidUint8Enum(
				f.Flag(),
				uint8(f.Get()),
				uint8(FullCompBasedStatsNone),
				uint8(FullCompBasedStatsStatistics),
				uint8(FullCompBasedStatsScoreAdjustment),
				uint8(FullCompBasedStatsScoreAdjustmentUnconditional),
			)
		}
	}
}

type FullCompBasedStatsEnum uint8

const (
	FullCompBasedStatsNone FullCompBasedStatsEnum = iota
	FullCompBasedStatsStatistics
	FullCompBasedStatsScoreAdjustment
	FullCompBasedStatsScoreAdjustmentUnconditional
)

func (f FullCompBasedStatsEnum) String() string {
	return strconv.Itoa(int(f))
}

func (f FullCompBasedStatsEnum) IsValid() bool {
	return f <= FullCompBasedStatsScoreAdjustmentUnconditional
}
