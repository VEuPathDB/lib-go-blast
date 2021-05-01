package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const (
	MaxBestHitOverhang float64 = 0.5
	MinBestHitOverhang float64 = 0
)

func NewBestHitOverhang(val float64) BestHitOverhang {
	return BestHitOverhang{true, val}
}

func NewEmptyBestHitOverhang() BestHitOverhang {
	return BestHitOverhang{}
}

func DecodeJSONBestHitOverhang(dec *gojay.Decoder, val *BestHitOverhang) error {
	val.set = true
	return dec.DecodeFloat64(&val.val)
}

type BestHitOverhang struct {
	set bool
	val float64
}

func (b *BestHitOverhang) IsSet() bool {
	return b.set
}

func (b *BestHitOverhang) Get() float64 {
	if !b.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return b.val
}

func (b *BestHitOverhang) Set(val float64) {
	b.set = true
	b.val = val
}

func (b *BestHitOverhang) Unset() {
	b.set = false
	b.val = 0
}

func (b *BestHitOverhang) Flag() string {
	return consts.FlagBestHitOverhang
}

func (b *BestHitOverhang) IsDefault() bool {
	return !b.IsSet()
}

func (b *BestHitOverhang) ArgString() string {
	return strconv.FormatFloat(b.Get(), 'f', -1, 64)
}

func (b *BestHitOverhang) FlagString() string {
	return b.Flag() + "=" + b.ArgString()
}

func (b *BestHitOverhang) Validate(em bval.ValidationBuilder) {
	_ = em.F64ExclusiveRange(b, MinBestHitOverhang, MaxBestHitOverhang)
}
