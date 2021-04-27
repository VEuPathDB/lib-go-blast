package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewMaxTargetSequences(val uint32) MaxTargetSequences {
	return MaxTargetSequences{true, val}
}

func NewEmptyMaxTargetSequences() MaxTargetSequences {
	return MaxTargetSequences{}
}

func DecodeJSONMaxTargetSequences(dec *gojay.Decoder, val *MaxTargetSequences) error {
	val.set = true
	return dec.DecodeUint32(&val.val)
}

type MaxTargetSequences struct {
	set bool
	val uint32
}

func (m MaxTargetSequences) IsSet() bool {
	return m.set
}

func (m MaxTargetSequences) Get() uint32 {
	if !m.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return m.val
}

func (m *MaxTargetSequences) Set(val uint32) {
	m.set = true
	m.val = val
}

func (m *MaxTargetSequences) Unset() {
	m.set = false
	m.val = 0
}

func (MaxTargetSequences) Flag() string {
	return consts.FlagMaxTargetSequences
}

func (m MaxTargetSequences) IsDefault() bool {
	return !m.set
}

func (m MaxTargetSequences) ArgString() string {
	return strconv.FormatUint(uint64(m.Get()), 10)
}

func (m MaxTargetSequences) FlagString() string {
	return m.Flag() + "=" + m.ArgString()
}

