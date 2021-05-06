package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

const (
	MaxQueryCoverageHSPPercent float64 = 100
	MinQueryCoverageHSPPercent float64 = 0
)

func NewQueryCoverageHSPPercent(val float64) QueryCoverageHSPPercent {
	return QueryCoverageHSPPercent{true, val}
}

func NewEmptyQueryCoverageHSPPercent() QueryCoverageHSPPercent {
	return QueryCoverageHSPPercent{}
}

func DecodeJSONQueryCoverageHSPPercent(dec *gojay.Decoder, val *QueryCoverageHSPPercent) error {
	val.set = true
	return dec.Float64(&val.val)
}

type QueryCoverageHSPPercent struct {
	set bool
	val float64
}

func (b QueryCoverageHSPPercent) IsSet() bool {
	return b.set
}

func (b QueryCoverageHSPPercent) Get() float64 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *QueryCoverageHSPPercent) Set(val float64) {
	b.set = true
	b.val = val
}

func (b *QueryCoverageHSPPercent) Unset() {
	b.set = false
	b.val = 0
}

func (b QueryCoverageHSPPercent) Flag() string {
	return consts.FlagQueryCoverageHSPPercent
}

func (b QueryCoverageHSPPercent) IsDefault() bool {
	return !b.IsSet()
}

func (b QueryCoverageHSPPercent) ArgString() string {
	return strconv.FormatFloat(b.Get(), 'f', -1, 64)
}

func (b QueryCoverageHSPPercent) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}