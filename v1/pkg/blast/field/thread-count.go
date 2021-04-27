package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"lib-go-blast/v1/pkg/blast/consts"
)

const DefaultThreadCount uint8 = 1

func NewThreadCount(val uint8) ThreadCount {
	return ThreadCount{true, val}
}

func NewEmptyThreadCount() ThreadCount {
	return ThreadCount{}
}

func DecodeJSONThreadCount(dec *gojay.Decoder, val *ThreadCount) error {
	val.set = true
	return dec.DecodeUint8(&val.val)
}

type ThreadCount struct {
	set bool
	val uint8
}

func (s ThreadCount) IsSet() bool {
	return s.set
}

func (s ThreadCount) Get() uint8 {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *ThreadCount) Set(val uint8) {
	s.set = true
	s.val = val
}

func (s *ThreadCount) Unset() {
	s.set = false
}

func (s ThreadCount) Flag() string {
	return consts.FlagNumThreads
}

func (s ThreadCount) IsDefault() bool {
	return !s.set || s.val == DefaultThreadCount
}

func (s ThreadCount) ArgString() string {
	return strconv.Itoa(int(s.Get()))
}

func (s ThreadCount) FlagString() string {
	return s.Flag() + "=" + s.ArgString()
}
