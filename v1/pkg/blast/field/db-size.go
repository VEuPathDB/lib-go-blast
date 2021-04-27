package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewDBSize(val int8) DBSize {
	return DBSize{true, val}
}

func NewEmptyDBSize() DBSize {
	return DBSize{}
}

func DecodeJSONDBSize(dec *gojay.Decoder, val *DBSize) error {
	val.set = true
	return dec.DecodeInt8(&val.val)
}

type DBSize struct {
	set bool
	val int8
}

func (s DBSize) IsSet() bool {
	return s.set
}

func (s DBSize) Get() int8 {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *DBSize) Set(val int8) {
	s.set = true
	s.val = val
}

func (s *DBSize) Unset() {
	s.set = false
}

func (s DBSize) Flag() string {
	return consts.FlagDBSize
}

func (s DBSize) IsDefault() bool {
	return !s.set
}

func (s DBSize) ArgString() string {
	return strconv.Itoa(int(s.Get()))
}

func (s DBSize) FlagString() string {
	return s.Flag() + "=" + s.ArgString()
}
