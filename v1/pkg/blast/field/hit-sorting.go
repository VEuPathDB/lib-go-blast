package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewHitSorting(val HitSortingType) HitSorting {
	return HitSorting{true, val}
}

func NewEmptyHitSorting() HitSorting {
	return HitSorting{}
}

func DecodeJSONHitSorting(dec *gojay.Decoder, val *HitSorting) error {
	val.set = true
	return dec.DecodeUint8((*uint8)(&val.val))
}

type HitSorting struct {
	set bool
	val HitSortingType
}

func (h HitSorting) Flag() string {
	return consts.FlagSortHits
}

func (h HitSorting) IsDefault() bool {
	return !h.set
}

func (h HitSorting) FlagString() string {
	return h.Flag() + "=" + h.Get().String()
}

func (h HitSorting) IsSet() bool {
	return h.set
}

func (h HitSorting) Get() HitSortingType {
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
