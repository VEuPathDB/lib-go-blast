package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

const (
	MaxBestHitScoreEdge float64 = 0.5
	MinBestHitScoreEdge float64 = 0
)

func NewBestHitScoreEdge(val float64) BestHitScoreEdge {
	return BestHitScoreEdge{true, val}
}

func NewEmptyBestHitScoreEdge() BestHitScoreEdge {
	return BestHitScoreEdge{}
}

func DecodeJSONBestHitScoreEdge(dec *gojay.Decoder, val *BestHitScoreEdge) error {
	val.set = true
	return dec.Float64(&val.val)
}

type BestHitScoreEdge struct {
	set bool
	val float64
}

func (b BestHitScoreEdge) IsSet() bool {
	return b.set
}

func (b BestHitScoreEdge) Get() float64 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *BestHitScoreEdge) Set(val float64) {
	b.set = true
	b.val = val
}

func (b *BestHitScoreEdge) Unset() {
	b.set = false
	b.val = 0
}

func (b BestHitScoreEdge) Flag() string {
	return consts.FlagBestHitScoreEdge
}

func (b BestHitScoreEdge) IsDefault() bool {
	return !b.IsSet()
}

func (b BestHitScoreEdge) ArgString() string {
	return strconv.FormatFloat(b.Get(), 'f', -1, 64)
}

func (b BestHitScoreEdge) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}
