package field

import (
	"strconv"

	"github.com/francoispqt/gojay"
	"github.com/veupathdb/lib-go-blast/v2/pkg/blast/consts"
)

func NewWindowSize(val uint32) WindowSize {
	return WindowSize{true, val}
}

func NewEmptyWindowSize() WindowSize {
	return WindowSize{}
}

func DecodeJSONWindowSize(dec *gojay.Decoder, val *WindowSize) error {
	val.set = true
	return dec.Uint32(&val.val)
}

type WindowSize struct {
	set bool
	val uint32
}

func (m *WindowSize) IsSet() bool {
	return m.set
}

func (m *WindowSize) Get() uint32 {
	if !m.IsSet() {
		panic("Cannot get the value of an empty option.")
	}

	return m.val
}

func (m *WindowSize) Set(val uint32) {
	m.set = true
	m.val = val
}

func (m *WindowSize) Unset() {
	m.set = false
	m.val = 0
}

func (WindowSize) Flag() string {
	return consts.FlagWindowSize
}

func (m *WindowSize) IsDefault() bool {
	return !m.set
}

func (m *WindowSize) ArgString() string {
	return strconv.FormatUint(uint64(m.Get()), 10)
}

func (m *WindowSize) FlagString() string {
	return m.Flag() + "=" + m.ArgString()
}

