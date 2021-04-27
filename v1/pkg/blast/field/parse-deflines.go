package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewParseDefLines(val bool) ParseDefLines {
	return ParseDefLines{true, val}
}

func NewEmptyParseDefLines() ParseDefLines {
	return ParseDefLines{}
}

func DecodeJSONParseDefLines(dec *gojay.Decoder, val *ParseDefLines) error {
	val.set = true
	return dec.DecodeBool(&val.val)
}

type ParseDefLines struct {
	set bool
	val bool
}

func (s ParseDefLines) IsSet() bool {
	return s.set
}

func (s ParseDefLines) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *ParseDefLines) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *ParseDefLines) Unset() {
	s.set = false
}

func (s ParseDefLines) Flag() string {
	return consts.FlagParseDefLines
}

func (s ParseDefLines) IsDefault() bool {
	return !s.set || !s.val
}

func (s ParseDefLines) ArgString() string {
	return ""
}

func (s ParseDefLines) FlagString() string {
	return s.Flag()
}

