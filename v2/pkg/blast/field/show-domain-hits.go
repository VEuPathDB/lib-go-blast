package field

import (
	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewShowDomainHits(val bool) ShowDomainHits {
	return ShowDomainHits{true, val}
}

func NewEmptyShowDomainHits() ShowDomainHits {
	return ShowDomainHits{}
}

func DecodeJSONShowDomainHits(dec *gojay.Decoder, val *ShowDomainHits) error {
	val.set = true
	return dec.Bool(&val.val)
}

type ShowDomainHits struct {
	set bool
	val bool
}

func (s *ShowDomainHits) IsSet() bool {
	return s.set
}

func (s *ShowDomainHits) Get() bool {
	if !s.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return s.val
}

func (s *ShowDomainHits) Set(val bool) {
	s.set = true
	s.val = val
}

func (s *ShowDomainHits) Unset() {
	s.set = false
}

func (s *ShowDomainHits) Flag() string {
	return consts.FlagShowDomainHits
}

func (s *ShowDomainHits) IsDefault() bool {
	return !s.set || !s.val
}

func (s *ShowDomainHits) FlagString() string {
	return s.Flag()
}

