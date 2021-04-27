package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewVersion(val bool) Version {
	return Version{true, val}
}

func NewEmptyVersion() Version {
	return Version{}
}

func DecodeJSONVersion(dec *gojay.Decoder, val *Version) error {
	val.set = true
	return dec.DecodeBool(&val.val)
}

type Version struct {
	set bool
	val bool
}

func (s Version) IsSet() bool {
	return s.set
}

func (s Version) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *Version) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *Version) Unset() {
	s.set = false
}

func (s Version) Flag() string {
	return consts.FlagVersion
}

func (s Version) IsDefault() bool {
	return !s.set || !s.val
}

func (s Version) FlagString() string {
	return s.Flag()
}

