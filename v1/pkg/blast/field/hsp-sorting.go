package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
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

func (h HSPSorting) Flag() string {
	return consts.FlagSortHSPs
}

func (h HSPSorting) IsDefault() bool {
	return !h.set
}

func (h HSPSorting) FlagString() string {
	return h.Flag() + "=" + h.Get().String()
}

func (h HSPSorting) IsSet() bool {
	return h.set
}

func (h HSPSorting) Get() HSPSortingType {
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
