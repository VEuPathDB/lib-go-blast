package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewRemote(val bool) Remote {
	return Remote{true, val}
}

func NewEmptyRemote() Remote {
	return Remote{}
}

func DecodeJSONRemote(dec *gojay.Decoder, val *Remote) error {
	val.set = true
	return dec.Bool(&val.val)
}

type Remote struct {
	set bool
	val bool
}

func (s Remote) IsSet() bool {
	return s.set
}

func (s Remote) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *Remote) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *Remote) Unset() {
	s.set = false
}

func (s Remote) Flag() string {
	return consts.FlagRemote
}

func (s Remote) IsDefault() bool {
	return !s.set || !s.val
}

func (s Remote) FlagString() string {
	return s.Flag()
}