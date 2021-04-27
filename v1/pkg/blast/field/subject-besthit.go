package field

import (
	"github.com/francoispqt/gojay"
	"lib-go-blast/v1/pkg/blast/consts"
)

func NewSubjectBestHit(val bool) SubjectBestHit {
	return SubjectBestHit{true, val}
}

func NewEmptySubjectBestHit() SubjectBestHit {
	return SubjectBestHit{}
}

func DecodeJSONSubjectBestHit(dec *gojay.Decoder, val *SubjectBestHit) error {
	val.set = true
	return dec.DecodeBool(&val.val)
}

type SubjectBestHit struct {
	set bool
	val bool
}

func (s SubjectBestHit) IsSet() bool {
	return s.set
}

func (s SubjectBestHit) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *SubjectBestHit) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *SubjectBestHit) Unset() {
	s.set = false
}

func (s SubjectBestHit) Flag() string {
	return consts.FlagSubjectBestHit
}

func (s SubjectBestHit) IsDefault() bool {
	return !s.set || !s.val
}

func (s SubjectBestHit) ArgString() string {
	return ""
}

func (s SubjectBestHit) FlagString() string {
	return s.Flag()
}

