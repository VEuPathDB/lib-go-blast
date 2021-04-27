package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewUngappedAlignmentsOnly(val bool) UngappedAlignmentsOnly {
	return UngappedAlignmentsOnly{true, val}
}

func NewEmptyUngappedAlignmentsOnly() UngappedAlignmentsOnly {
	return UngappedAlignmentsOnly{}
}

func DecodeJSONUngappedAlignmentsOnly(dec *gojay.Decoder, val *UngappedAlignmentsOnly) error {
	val.set = true
	return dec.DecodeBool(&val.val)
}

type UngappedAlignmentsOnly struct {
	set bool
	val bool
}

func (s UngappedAlignmentsOnly) IsSet() bool {
	return s.set
}

func (s UngappedAlignmentsOnly) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *UngappedAlignmentsOnly) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *UngappedAlignmentsOnly) Unset() {
	s.set = false
}

func (s UngappedAlignmentsOnly) Flag() string {
	return consts.FlagUngappedAlignmentsOnly
}

func (s UngappedAlignmentsOnly) IsDefault() bool {
	return !s.set || !s.val
}

func (s UngappedAlignmentsOnly) ArgString() string {
	return ""
}

func (s UngappedAlignmentsOnly) FlagString() string {
	return s.Flag()
}

