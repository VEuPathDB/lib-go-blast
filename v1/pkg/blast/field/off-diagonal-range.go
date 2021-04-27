package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"lib-go-blast/v1/pkg/blast/consts"
)

func NewOffDiagonalRange(val uint32) OffDiagonalRange {
	return OffDiagonalRange{true, val}
}

func NewEmptyOffDiagonalRange() OffDiagonalRange {
	return OffDiagonalRange{}
}

func DecodeJSONOffDiagonalRange(dec *gojay.Decoder, val *OffDiagonalRange) error {
	val.set = true
	return dec.DecodeUint32(&val.val)
}

type OffDiagonalRange struct {
	set bool
	val uint32
}

func (m OffDiagonalRange) IsSet() bool {
	return m.set
}

func (m OffDiagonalRange) Get() uint32 {
	if !m.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return m.val
}

func (m *OffDiagonalRange) Set(val uint32) {
	m.set = true
	m.val = val
}

func (m *OffDiagonalRange) Unset() {
	m.set = false
	m.val = 0
}

func (OffDiagonalRange) Flag() string {
	return consts.FlagOffDiagonalRange
}

func (m OffDiagonalRange) IsDefault() bool {
	return !m.set
}

func (m OffDiagonalRange) ArgString() string {
	return strconv.FormatUint(uint64(m.Get()), 10)
}

func (m OffDiagonalRange) FlagString() string {
	return m.Flag() + "=" + m.ArgString()
}

