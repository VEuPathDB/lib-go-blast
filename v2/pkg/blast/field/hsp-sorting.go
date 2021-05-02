package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewHSPSorting(val HSPSortingType) HSPSorting {
	return HSPSorting{true, val}
}

func NewEmptyHSPSorting() HSPSorting {
	return HSPSorting{}
}

func DecodeJSONHSPSorting(dec *gojay.Decoder, val *HSPSorting) error {
	val.set = true
	return dec.Uint8((*uint8)(&val.val))
}

type HSPSorting struct {
	set bool
	val HSPSortingType
}

func (h *HSPSorting) Flag() string {
	return consts.FlagSortHSPs
}

func (h *HSPSorting) IsDefault() bool {
	return !h.set
}

func (h *HSPSorting) FlagString() string {
	return h.Flag() + "=" + h.Get().String()
}

func (h *HSPSorting) IsSet() bool {
	return h.set
}

func (h *HSPSorting) Get() HSPSortingType {
	if !h.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return h.val
}

func (h *HSPSorting) Set(kind HSPSortingType) {
	h.set = true
	h.val = kind
}

func (h *HSPSorting) Unset() {
	h.set = false
}

func (h *HSPSorting) Validate(em bval.ValidationBuilder) {
	if !h.IsDefault() && !h.Get().IsValid() {
		_ = em.InvalidUint8Enum(
			h.Flag(),
			uint8(h.Get()),
			uint8(HSPSortingTypeByHSPExpectValue),
			uint8(HSPSortingTypeByHSPScore),
			uint8(HSPSortingTypeByHSPQueryStart),
			uint8(HSPSortingTypeByHSPPercentIdentity),
			uint8(HSPSortingTypeByHSPSubjectStart),
		)
	}
}

type HSPSortingType uint8

const (
	HSPSortingTypeByHSPExpectValue HSPSortingType = iota
	HSPSortingTypeByHSPScore
	HSPSortingTypeByHSPQueryStart
	HSPSortingTypeByHSPPercentIdentity
	HSPSortingTypeByHSPSubjectStart
)

func (h HSPSortingType) String() string {
	return strconv.Itoa(int(h))
}

func (h HSPSortingType) IsValid() bool {
	return h <= HSPSortingTypeByHSPSubjectStart
}
