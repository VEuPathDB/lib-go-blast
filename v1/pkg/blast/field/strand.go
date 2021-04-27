package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

const DefaultStrandType = StrandTypeBoth

func NewStrand(val StrandType) Strand {
	return Strand{true, val}
}

func NewEmptyStrand() Strand {
	return Strand{}
}

func DecodeJSONStrand(dec *gojay.Decoder, val *Strand) error {
	val.set = true
	return dec.DecodeString((*string)(&val.val))
}

type Strand struct {
	set bool
	val StrandType
}

func (s Strand) IsSet() bool {
	return s.set
}

func (s Strand) Get() StrandType {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *Strand) Set(val StrandType) {
	s.set = true
	s.val = val
}

func (s *Strand) Unset() {
	s.set = false
}

func (s Strand) Flag() string {
	return consts.FlagStrand
}

func (s Strand) IsDefault() bool {
	return !s.set || s.val == DefaultStrandType
}

func (s Strand) FlagString() string {
	return s.Flag() + "=" + s.Get().String()
}

type StrandType string

const (
	StrandTypeBoth  StrandType = "both"
	StrandTypeMinus StrandType = "minus"
	StrandTypePlus  StrandType = "plus"
)

func (s StrandType) String() string {
	return string(s)
}
