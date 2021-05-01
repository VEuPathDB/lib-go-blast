package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/bval"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

const MinMaxHSPs uint32 = 1

func NewMaxHSPs(val uint32) MaxHSPs {
	return MaxHSPs{true, val}
}

func NewEmptyMaxHSPs() MaxHSPs {
	return MaxHSPs{}
}

func DecodeJSONMaxHSPs(dec *gojay.Decoder, val *MaxHSPs) error {
	val.set = true
	return dec.DecodeUint32(&val.val)
}

type MaxHSPs struct {
	set bool
	val uint32
}

func (m *MaxHSPs) IsSet() bool {
	return m.set
}

func (m *MaxHSPs) Get() uint32 {
	if !m.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return m.val
}

func (m *MaxHSPs) Set(val uint32) {
	m.set = true
	m.val = val
}

func (m *MaxHSPs) Unset() {
	m.set = false
	m.val = 0
}

func (m *MaxHSPs) Flag() string {
	return consts.FlagMaxHSPs
}

func (m *MaxHSPs) IsDefault() bool {
	return !m.IsSet()
}

func (m *MaxHSPs) ArgString() string {
	return strconv.FormatUint(uint64(m.Get()), 10)
}

func (m *MaxHSPs) FlagString() string {
	return m.Flag() + "=" + m.ArgString()
}

func (m *MaxHSPs) Validate(em bval.ValidationBuilder) {
	_ = em.U32GTEQ(m, MinMaxHSPs)
}
