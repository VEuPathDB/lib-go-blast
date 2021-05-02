package field

import (
	"errors"
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const DefaultShortCompBasedStats = ShortCompBasedStatsStatistics

func DecodeJSONShortCompBasedStats(dec *gojay.Decoder, val *ShortCompBasedStats) error {
	var tmp string
	if err := dec.String(&tmp); err != nil {
		return err
	}

	val.set = true

	switch tmp {
	case "D", "d":
		val.val = ShortCompBasedStatsStatistics
	case "0", "F", "f":
		val.val = ShortCompBasedStatsNone
	case "1", "T", "t":
		val.val = ShortCompBasedStatsStatistics
	default:
		return errors.New("Invalid comp based stats flag value.")
	}

	return nil
}

type ShortCompBasedStats struct {
	set bool
	val ShortCompBasedStatsEnum
}

func (f *ShortCompBasedStats) IsSet() bool {
	return f.set
}

func (f *ShortCompBasedStats) Get() ShortCompBasedStatsEnum {
	if !f.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return f.val
}

func (f *ShortCompBasedStats) Set(val ShortCompBasedStatsEnum) {
	f.set = true
	f.val = val
}

func (f *ShortCompBasedStats) Unset() {
	f.set = false
}

func (f *ShortCompBasedStats) Flag() string {
	return consts.FlagCompBasedStats
}

func (f *ShortCompBasedStats) IsDefault() bool {
	return !f.set || f.val == DefaultShortCompBasedStats
}

func (f *ShortCompBasedStats) FlagString() string {
	return f.Flag() + "=" + f.Get().String()
}

func (f *ShortCompBasedStats) Validate(em bval.ValidationBuilder) {
	if !f.IsDefault() {
		em.InvalidUint8Enum(
			f.Flag(),
			uint8(f.Get()),
			uint8(ShortCompBasedStatsNone),
			uint8(ShortCompBasedStatsStatistics),
		)
	}
}

type ShortCompBasedStatsEnum uint8

const (
	ShortCompBasedStatsNone ShortCompBasedStatsEnum = iota
	ShortCompBasedStatsStatistics
)

func (f ShortCompBasedStatsEnum) String() string {
	return strconv.Itoa(int(f))
}

func (f ShortCompBasedStatsEnum) IsValid() bool {
	return f <= ShortCompBasedStatsStatistics
}
