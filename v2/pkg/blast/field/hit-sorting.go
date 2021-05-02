package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewHitSorting(val HitSortingType) HitSorting {
	return HitSorting{true, val}
}

func NewEmptyHitSorting() HitSorting {
	return HitSorting{}
}

func DecodeJSONHitSorting(dec *gojay.Decoder, val *HitSorting) error {
	val.set = true
	return dec.Uint8((*uint8)(&val.val))
}

type HitSorting struct {
	set bool
	val HitSortingType
}

func (h *HitSorting) Flag() string {
	return consts.FlagSortHits
}

func (h *HitSorting) IsDefault() bool {
	return !h.set
}

func (h *HitSorting) FlagString() string {
	return h.Flag() + "=" + h.Get().String()
}

func (h *HitSorting) IsSet() bool {
	return h.set
}

func (h *HitSorting) Get() HitSortingType {
	if !h.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return h.val
}

func (h *HitSorting) Set(kind HitSortingType) {
	h.set = true
	h.val = kind
}

func (h *HitSorting) Unset() {
	h.set = false
}

func (h *HitSorting) Validate(em bval.ValidationBuilder) {
	if !h.IsDefault() && !h.Get().IsValid() {
		_ = em.InvalidUint8Enum(
			h.Flag(),
			uint8(h.Get()),
			uint8(HitSortingTypeByExpectValue),
			uint8(HitSortingTypeByBitScore),
			uint8(HitSortingTypeByTotalScore),
			uint8(HitSortingTypeByPercentIdentity),
			uint8(HitSortingTypeByQueryCoverage),
		)
	}
}

type HitSortingType uint8

const (
	HitSortingTypeByExpectValue HitSortingType = iota
	HitSortingTypeByBitScore
	HitSortingTypeByTotalScore
	HitSortingTypeByPercentIdentity
	HitSortingTypeByQueryCoverage
)

func (h HitSortingType) String() string {
	return strconv.Itoa(int(h))
}

func (h HitSortingType) IsValid() bool {
	return h <= HitSortingTypeByQueryCoverage
}
