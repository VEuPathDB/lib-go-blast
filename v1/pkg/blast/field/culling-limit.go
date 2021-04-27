package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v1/pkg/blast/consts"
)

func NewCullingLimit(val uint32) CullingLimit {
	return CullingLimit{true, val}
}

func NewEmptyCullingLimit() CullingLimit {
	return CullingLimit{}
}

func DecodeJSONCullingLimit(dec *gojay.Decoder, val *CullingLimit) error {
	val.set = true
	return dec.DecodeUint32(&val.val)
}

type CullingLimit struct {
	set bool
	val uint32
}

func (m CullingLimit) IsSet() bool {
	return m.set
}

func (m CullingLimit) Get() uint32 {
	if !m.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return m.val
}

func (m *CullingLimit) Set(val uint32) {
	m.set = true
	m.val = val
}

func (m *CullingLimit) Unset() {
	m.set = false
	m.val = 0
}

func (m CullingLimit) Flag() string {
	return consts.FlagCullingLimit
}

func (m CullingLimit) IsDefault() bool {
	return !m.IsSet()
}

func (m CullingLimit) ArgString() string {
	return strconv.FormatUint(uint64(m.Get()), 10)
}

func (m CullingLimit) FlagString() string {
	return m.Flag() + "=" + m.ArgString()
}
